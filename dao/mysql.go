package dao

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 数据库驱动包
)

var db orm.Ormer

func init() {
	driverName := beego.AppConfig.String("driverName")
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	// dbConn := "root:13572281710@tcp(127.0.0.1:3306)/testDemo?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		logs.Error(err)
	}
	orm.SetMaxIdleConns("default", 1000) // 设置闲置的连接数
	orm.SetMaxOpenConns("default", 2000) // 最大打开的连接数

	orm.RegisterModel(new(UserInformation))
	orm.RegisterModel(new(WaterInformation))
	orm.RegisterModel(new(WaterInformationDetail))
	orm.RegisterModel(new(ForumArticle))
	orm.RegisterModel(new(ForumProperty))
	orm.RegisterModel(new(ForumLike))
	orm.RegisterModel(new(ForumContent))
	orm.RegisterModel(new(ForumComment))
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(ReadMessage))

	db = orm.NewOrm()
	logs.Info("Init mysql success...")
}
