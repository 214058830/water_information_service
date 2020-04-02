package dao

import (
	"github.com/astaxie/beego/orm"
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type ForumLike struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	ArticleId        int      `json:"article_id" orm:"size(11);column(article_id)"`
	LikeUserMail     string   `json:"like_user_mail" orm:"size(64);column(like_user_mail)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const forumLikeTableName = "forum_like"

func (u *ForumLike) TableName() string {
	return forumLikeTableName
}

// 删除点赞信息
func DeleteForumLikeByArticleIdAndMail(article_id int, mail string) (num int64, err error) {
	if num, err = db.Delete(&ForumLike{ArticleId: article_id, LikeUserMail: mail}, "article_id", "like_user_mail"); err != nil {
		logs.Error(err)
	}
	return
}

// 查询用户有没有给帖子点赞过
func SelectLikeByArticleAndMail(article_id int, mail string) (err error) {
	data := ForumLike{ArticleId: article_id, LikeUserMail: mail}
	err = db.Read(&data, "article_id", "like_user_mail")
	if err != nil && err != orm.ErrNoRows {
		logs.Error(err)
	}
	return
}

// 查询用户有没有给帖子点赞过
func ReadOrCreateLikeByArticleAndMail(article_id int, mail string) (created bool, err error) {
	data := ForumLike{ArticleId: article_id, LikeUserMail: mail}
	if created, _, err = db.ReadOrCreate(&data, "article_id", "like_user_mail"); err != nil {
		logs.Error(err)
	}
	return
}
