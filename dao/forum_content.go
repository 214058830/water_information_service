package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type ForumContent struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	Title            string   `json:"title" orm:"size(24);column(title)"`
	ContentPath      string   `json:"content_path" orm:"size(64);column(content_path)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const forumContentTableName = "forum_content"

func (u *ForumContent) TableName() string {
	return forumContentTableName
}

// 查询论坛帖子编辑内容信息 参数:[id]
func SelecForumContentById(id int) (r ForumContent, err error) {
	_, err = db.QueryTable(forumContentTableName).Filter("Id", id).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 更新标题和文件路径 参数: [id, title, filepath]
func UpdateContentTitleAndFilePath(id int, title string, filepath string) (err error) {
	var articleInfo ForumContent
	if articleInfo, err = SelecForumContentById(id); err == nil {
		articleInfo.Title = title
		articleInfo.ContentPath = filepath
		if _, err = db.Update(&articleInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 增加论坛帖子编辑内容信息 参数: [title, filepath]
func InsertForumContent(title string, filePath string) (err error) {
	var data ForumContent
	data.Title = title
	data.ContentPath = filePath
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return err
}
