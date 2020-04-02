package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type ForumProperty struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	UserNum          int      `json:"user_num" orm:"size(11);column(user_num)"`
	ArticleNum       int      `json:"article_num" orm:"size(11);column(article_num)"`
	CommentNum       int      `json:"comment_num" orm:"size(11);column(comment_num)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const forumPropertyTableName = "forum_property"

func (u *ForumProperty) TableName() string {
	return forumPropertyTableName
}

// 查询论坛属性信息
func SelectForumProperty() (r ForumProperty, err error) {
	_, err = db.QueryTable(forumPropertyTableName).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 修改会员数信息
func UpdateForumPropertyUserNum(num int) (err error) {
	var ForumPropertyInfo ForumProperty
	if ForumPropertyInfo, err = SelectForumProperty(); err == nil {
		ForumPropertyInfo.UserNum += num
		if _, err = db.Update(&ForumPropertyInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 修改帖子数信息
func UpdateForumPropertyArticleNum(num int) (err error) {
	var ForumPropertyInfo ForumProperty
	if ForumPropertyInfo, err = SelectForumProperty(); err == nil {
		ForumPropertyInfo.ArticleNum += num
		if _, err = db.Update(&ForumPropertyInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 修改评论数信息
func UpdateForumPropertyCommentNum(num int) (err error) {
	var ForumPropertyInfo ForumProperty
	if ForumPropertyInfo, err = SelectForumProperty(); err == nil {
		ForumPropertyInfo.CommentNum += num
		if _, err = db.Update(&ForumPropertyInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}