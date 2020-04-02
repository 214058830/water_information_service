package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type UserInformation struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	Mail             string   `json:"mail" orm:"size(64);column(mail)"`
	UserName         string   `json:"user_name" orm:"size(64);column(user_name)"`
	PassWord         string   `json:"pass_word" orm:"size(16);column(pass_word)"`
	Logo             bool     `json:"logo" orm:"default(0); size(12); column(logo)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const userTableName = "user"

//表名被改为UserTable
func (u *UserInformation) TableName() string {
	return userTableName
}

// 条件查询 按照UserMail查询用户信息 参数: [mail]
func SelectWithUserMail(mail string) (UserInformation, error) {
	var data UserInformation
	data.Mail = mail
	err := db.Read(&data, "Mail")
	if err != nil {
		logs.Error(err)
	}
	return data, err
}

// 查询所有用户信息
func SelectAllUser() (r []UserInformation, err error) {
	//_, err = db.QueryTable("user_information").Filter("Id", "201606010211").All(&r)
	_, err = db.QueryTable(userTableName).All(&r)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 更新用户信息 参数: [id, username, password]
func UpdataUserInfo(id int, userName string, passWord string) (error) {
	var data UserInformation
	data.Id = id
	data.UserName = userName
	data.PassWord = passWord
	_, err := db.Update(&data, "user_name", "pass_word")
	if err != nil {
		logs.Error(err)
	}
	return err
}

// 添加新用户 参数: [name, mail, password]
func AddUser(name string, mail string, password string) (error) {
	userInfo := UserInformation{}
	userInfo.UserName = name
	userInfo.Mail = mail
	userInfo.PassWord = password
	_, err := db.Insert(&userInfo)
	if err != nil {
		logs.Error(err)
	}
	return err
}

// 更新管理员人员 参数: [id, bool]
func UpdataAdmin(id int, flag bool) error {
	var user UserInformation
	user.Id = id
	user.Logo = flag
	_, err := db.Update(&user, "logo")
	if err != nil {
		logs.Error(err)
	}
	return err
}
