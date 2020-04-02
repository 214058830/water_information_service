package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type Message struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	Mail             string   `json:"mail" orm:"size(64); column(mail)"`
	ArticleId        int      `json:"article_id" orm:"size(11);column(article_id)"`
	NoticeMail       string   `json:"notice_mail" orm:"size(64); column(notice_mail)"`
	Type             string   `json:"type" orm:"size(12); column(type)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const messageTableName = "message"

func (u *Message) TableName() string {
	return messageTableName
}

// 查询用户所有消息
func SelectMessage(mail string, messageType string) (r []Message, err error) {
	if _, err = db.QueryTable(messageTableName).Filter("mail", mail).Filter("type", messageType).Filter("flag", false).All(&r); err != nil {
		logs.Error(err)
	}
	return
}

// 增加消息
func InsertMessage(mail string, articleId int, noticeMail string, messageType string) (err error) {
	data := Message{Mail: mail, ArticleId: articleId, NoticeMail: noticeMail, Type: messageType}
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}
