package models

import (
	"water_information_service/dao"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 查询所有用户信息
func SelectAllUser() ([]dao.UserInformation, string) {
	data, err := dao.SelectAllUser()
	if err != nil {
		logs.Error(err.Error())
		return nil, "5000"
	}
	return data, "2000"
}

// 验证用户的账号密码是否正确
func SelectOneOfAllUser(mail string, password string) (r dao.UserInformation, code string) {
	var err error
	r, err = dao.SelectWithUserMail(mail)
	if err != nil || r.PassWord != password || r.Flag == true {
		if err != nil && err != orm.ErrNoRows {
			logs.Error(err)
			code = "5000"
		} else {
			code = "2001"
		}
	} else {
		code = "2000"
	}
	return
}

// 登录成功用户设置临时token用做通过后端鉴权
func SetToken(token string) {
	dao.SetToken(token)
}

// 注册新用户
func RegisterUserInfo(name string, mail string, password string) (r string, code string) {
	var err error
	_, err = dao.SelectWithUserMail(mail)
	if err != orm.ErrNoRows {
		if err != nil {
			logs.Error(err)
			code = "5000"
		} else {
			code = "2003"
		}
		return
	}
	err = dao.AddUser(name, mail, password)
	if err != nil {
		logs.Error(err)
		code = "5000"
	} else {
		code = "2000"
	}
	if err := dao.UpdateForumPropertyUserNum(1); err != nil {
		logs.Error(err)
	}
	return
}

// 更新用户信息
func UpdataUserInfo(id int, userName string, passWord string) (code string) {
	err := dao.UpdataUserInfo(id, userName, passWord)
	if err != nil {
		logs.Error(err)
		code = "5000"
	} else {
		code = "2000"
	}
	return
}

// 更新用户管理员权限
func UpdataAdmin(id int, flag bool) (code string) {
	if err := dao.UpdataAdmin(id, flag); err != nil {
		logs.Error(err)
		code = "5000"
	} else {
		code = "2000"
	}
	return
}
