package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ReadMessage struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	Mail             string   `json:"mail" orm:"size(64); column(mail)"`
	Type             string   `json:"type" orm:"size(12); column(type)"`
	ReadNum          int      `json:"read_num" orm:"size(11); column(read_num)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const readMessageTableName = "read_message"

func (u *ReadMessage) TableName() string {
	return readMessageTableName
}

// 查询用户已读消息个数
func SelectReadMessage(mail string, messageType string) (r ReadMessage, err error) {
	r.Mail = mail
	r.Type = messageType
	if _, _, err = db.ReadOrCreate(&r, "mail", "type"); err != nil {
		logs.Error(err)
	}
	return
}

// 修改已读消息个数
func AlterReadMessage(mail string, messageType string, number int) (err error) {
	if _, err = db.QueryTable(readMessageTableName).Filter("mail", mail).Filter("type", messageType).Update(orm.Params{
		"read_num": number,
	}); err != nil {
		logs.Error(err)
	}
	return
}
