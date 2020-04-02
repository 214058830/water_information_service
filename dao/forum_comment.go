package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type ForumComment struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	ArticleId        int      `json:"article_id" orm:"size(11);column(article_id)"`
	UserMail         string   `json:"user_mail" orm:"size(64);column(user_mail)"`
	Content          string   `json:"content" orm:"size(64);column(content)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const forumCommentTableName = "forum_comment"

func (u *ForumComment) TableName() string {
	return forumCommentTableName
}

// 查询帖子评论内容
func SelectForumCommentByArticleID(articleId int) (r []ForumComment, err error) {
	_, err = db.QueryTable(forumCommentTableName).Filter("ArticleId", articleId).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 增加帖子评论
func InsertForumComment(article_id int, mail string, comntent string)(err error){
	data := ForumComment{}
	data.ArticleId = article_id
	data.UserMail = mail
	data.Content = comntent
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return
}

// 删除帖子评论
func DeleteForumComment(article_id int)(err error){
	if _, err = db.Delete(&ForumComment{ArticleId: article_id}); err != nil {
		logs.Error(err)
	}
	return
}