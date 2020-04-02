package models

import (
	"water_information_service/dao"
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type messageRes struct {
	NoticeUser       string   `json:"notice_user"`
	ArticleId        int      `json:"article_id"`
	ArticleTitle     string   `json:"article_title"`
	CreateTime       lib.Time `json:"create_time"`
	LastModifiedTime lib.Time `json:"last_modified_time"`
}

// 获取用户某种类型所有消息
func GetAllMessage(mail string, messageType string) (r interface{}, code string) {
	var res []messageRes
	if temp, err := dao.SelectMessage(mail, messageType); err != nil {
		logs.Error(err)
		code = "5000"
		return
	} else {
		for _, v := range temp {
			userInfo, _ := dao.SelectWithUserMail(v.NoticeMail)
			articleInfo, _ := dao.SelecForumArticleById(v.ArticleId)
			res = append(res, messageRes{NoticeUser: userInfo.UserName, ArticleId: v.ArticleId, ArticleTitle: articleInfo.Title, CreateTime: v.CreateTime, LastModifiedTime: v.LastModifiedTime})
		}
	}
	r = res
	code = "2000"
	return
}

// 更新用户已读消息个数
func UpdataReadMessageNum(mail string, messageType string) {
	if temp, err := dao.SelectMessage(mail, messageType); err != nil {
		logs.Error(err)
	} else {
		if err = dao.AlterReadMessage(mail, messageType, len(temp)); err != nil {
			logs.Error(err)
		}
	}
}

type MessageNumber struct {
	Like    int `json:"like_num"`
	Comment int `json:"comment_num"`
}

// 获取用户的未读消息数量
func GetMessageNum(mail string) (r interface{}, code string) {
	var data MessageNumber
	messageType := []string{"like", "comment"}
	for _, v := range messageType {
		readTemp, readerr := dao.SelectReadMessage(mail, v)
		if readerr != nil {
			logs.Error(readerr)
			code = "5000"
			return
		}
		temp, err := dao.SelectMessage(mail, v)
		if err != nil {
			logs.Error(err)
			code = "5000"
			return
		} else {
			if (v == "like") {
				data.Like = len(temp) - readTemp.ReadNum
			} else {
				data.Comment = len(temp) - readTemp.ReadNum
			}
		}
	}
	code = "2000"
	r = data
	return
}

// 插入点赞评论等通知消息
func InsertMessage(article_id int, notice_mail string, messageType string) {
	articleInfo, err := dao.SelecForumArticleById(article_id)
	if err != nil {
		logs.Error(err)
	}
	err = dao.InsertMessage(articleInfo.Mail, article_id, notice_mail, messageType)
	if err != nil {
		logs.Error(err)
	}
}
