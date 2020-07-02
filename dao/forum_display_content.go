package dao

import (
	"water_information_service/lib"

	"github.com/astaxie/beego/logs"
)

type ForumArticle struct {
	Id               int      `json:"id" orm:"pk;size(11);column(id)"`
	Mail             string   `json:"mail" orm:"size(64);column(mail)"`
	UserName         string   `json:"user_name" orm:"size(64);column(user_name)"`
	Title            string   `json:"title" orm:"size(24);column(title)"`
	ShareNum         int      `json:"share_num" orm:"size(11);column(share_num)"`
	LikeNum          int      `json:"like_num" orm:"size(11);column(like_num)"`
	CommentNum       int      `json:"comment_num" orm:"size(11);column(comment_num)"`
	Logo             bool     `json:"logo" orm:"default(0); column(logo)"`
	Announcement     bool     `json:"announcement" orm:"default(0); column(announcement)"`
	ContentPath      string   `json:"content_path" orm:"size(64);column(content_path)"`
	CreateTime       lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
	LastModifiedTime lib.Time `json:"last_modified_time" orm:"auto_now; type(datetime); column(last_modified_time)"`
	Flag             bool     `json:"flag" orm:"default(0); column(flag)"`
}

// 数据库表名
const forumTableName = "forum_display_content"

func (u *ForumArticle) TableName() string {
	return forumTableName
}

// 查询所有论坛帖子列表
func SelectAllForumArticle() (r []ForumArticle, err error) {
	_, err = db.QueryTable(forumTableName).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 查询某用户的帖子列表 参数: [mail]
func SelectAllForumArticleByMail(mail string) (r []ForumArticle, err error) {
	_, err = db.QueryTable(forumTableName).Filter("Mail", mail).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 浏览论坛文章信息 参数: [id]
func SelecForumArticleById(id int) (r ForumArticle, err error) {
	_, err = db.QueryTable(forumTableName).Filter("Id", id).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 修改赞数信息 参数: [id, num] num: -1/1 点赞/取消赞
func UpdateLikeNum(id int, num int) (err error) {
	var articleInfo ForumArticle
	if articleInfo, err = SelecForumArticleById(id); err == nil {
		articleInfo.LikeNum += num
		if _, err = db.Update(&articleInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 修改分享数信息 参数: [id, num] num: -1/1 点赞/取消赞
func UpdateShareNum(id int, num int) (err error) {
	var articleInfo ForumArticle
	if articleInfo, err = SelecForumArticleById(id); err == nil {
		articleInfo.ShareNum += num
		if _, err = db.Update(&articleInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 修改评论数信息 参数: [id, num]
func UpdateCommentNum(id int, num int) (err error) {
	var articleInfo ForumArticle
	if articleInfo, err = SelecForumArticleById(id); err == nil {
		articleInfo.CommentNum += num
		if _, err = db.Update(&articleInfo); err != nil {
			logs.Error(err)
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 更新标题和文件路径 参数: [id, title, filepath]
func UpdateDisplayTitleAndFilePath(id int, title string, filepath string) (err error) {
	var articleInfo ForumArticle
	if articleInfo, err = SelecForumArticleById(id); err == nil {
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

// 增加论坛帖子 参数: [mail, username, title, filepath]
func InsertForumArticle(mail string, username string, title string, filePath string) (err error) {
	var data ForumArticle
	data.Mail = mail
	data.UserName = username
	data.Title = title
	data.ContentPath = filePath
	_, err = db.Insert(&data)
	if err != nil {
		logs.Error(err)
	}
	return err
}

// 更新置顶标志
func UpdateDisplayStickyLogo(id int, logo bool) (err error) {
	var articleInfo ForumArticle
	if articleInfo, err = SelecForumArticleById(id); err == nil {
		if articleInfo.Logo != logo {
			articleInfo.Logo = logo
			if _, err = db.Update(&articleInfo); err != nil {
				logs.Error(err)
			}
		}
	} else {
		logs.Error(err)
	}
	return err
}

// 查询公告帖子信息
func SelectAllForumArticleByAnnouncement() (r []ForumArticle, err error) {
	_, err = db.QueryTable(forumTableName).Filter("Announcement", true).All(&r)
	if err != nil {
		logs.Error(err)
		return
	}
	return
}

// 设置帖子为公告
func SetForumArticleAnnouncement(id int) (err error) {
	var articleInfo ForumArticle
	articleInfo.Id = id
	if articleInfo.Announcement != true {
		articleInfo.Announcement = true
		if _, err = db.Update(&articleInfo, "announcement"); err != nil {
			logs.Error(err)
		}
	}
	return err
}

// 取消所有的公告帖子
func CancelAllForumArticleAnnouncement() (err error) {
	var data []ForumArticle
	if data, err = SelectAllForumArticleByAnnouncement(); err != nil {
		logs.Error(err)
	} else {
		for _, v := range data {
			v.Announcement = false
			if _, err = db.Update(&v, "announcement"); err != nil {
				logs.Error(err)
			}
		}
	}
	logs.Info(data)
	return
}

// 删除展示内容
func DeleteDisplayContent(id int) (err error) {
	if _, err = db.Delete(&ForumArticle{Id: id}); err != nil {
		logs.Error(err)
	}
	return
}
