package dao

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

var (
	pool      *redis.Pool
	redisHost = beego.AppConfig.String("reids_host")
)

// newRedisPool:创建redis连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				logs.Error(err)
				return nil, err
			}
			return c, nil
		},
		//定时检查redis是否出状况
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

//初始化redis连接池
func init() {
	pool = newRedisPool()
	logs.Info("Init redis success...")
}

//对外暴露连接池
func RedisPool() *redis.Pool {
	return pool
}

// 设置token
func SetToken(key string) {
	rc := RedisPool().Get()
	defer rc.Close()
	//默认不选库操作的是0, 1为token库
	if _, err := rc.Do("select", 1); err != nil {
		logs.Error(err)
	}
	// 写入值后一小时后过期
	_, err := rc.Do("set", key, true, "EX", "3600")
	if err != nil {
		logs.Error(err)
	}
}

// 获取token
func GetToken(key string) interface{} {
	rc := RedisPool().Get()
	defer rc.Close()
	//默认不选库操作的是0, 1为token库
	if _, err := rc.Do("select", 1); err != nil {
		logs.Error(err)
	}
	r, _ := redis.String(rc.Do("get", key))
	return r
}

// 删除token
func DelToken(key string) interface{} {
	rc := RedisPool().Get()
	defer rc.Close()
	//默认不选库操作的是0, 1为token库
	if _, err := rc.Do("select", 1); err != nil {
		logs.Error(err)
	}
	r, _ := redis.String(rc.Do("del", key))
	return r
}