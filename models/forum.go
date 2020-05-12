package models

import (
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"water_information_service/lib"

	"water_information_service/dao"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 查询所有论坛帖子列表
func SelectAllForumArticle() (r []dao.ForumArticle, code string) {
	data, err := dao.SelectAllForumArticle()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	} else {
		temp := make([]dao.ForumArticle, 0, len(data))
		for _, v := range data {
			if v.Logo == true {
				r = append(r, v)
			} else {
				temp = append(temp, v)
			}
		}
		r = append(r, temp...)
		code = "2000"
	}
	return
}

// 查询某个用户的论坛帖子列表  参数: [mail]
func SelectAllForumArticleByMail(mail string) (r []dao.ForumArticle, code string) {
	var err error
	r, err = dao.SelectAllForumArticleByMail(mail)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	} else {
		code = "2000"
	}
	return
}

// 保存帖子内容到文件中 参数: [content] 返回值: [filepath]
func writeFileByArticleContent(content string) (string, error) {
	//获取当前时间戳
	timeUnixNano := time.Now().UnixNano()
	fileName := strconv.FormatInt(timeUnixNano, 10) + ".txt"
	filePath := beego.AppConfig.String("articlePath") + fileName
	file, err := os.Create(filePath);
	defer file.Close()
	if err != nil {
		logs.Error(err)
	} else {
		file.WriteString(content);
	}
	return filePath, err
}

// 发布论坛帖子
func ReleaseForumArticle(mail string, username string, title string, content string, display_content string) (code string) {
	filePath, err := writeFileByArticleContent(display_content)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	err = dao.InsertForumArticle(mail, username, title, filePath)
	if err == nil {
		filePath, err = writeFileByArticleContent(content)
		if err != nil {
			logs.Error(err)
			code = "5000"
			return
		}
		err = dao.InsertForumContent(title, filePath)
	}
	if err == nil {
		if err = dao.UpdateForumPropertyArticleNum(1); err != nil {
			logs.Error(err)
		}
	}
	code = "2000"
	return
}

type Comment struct {
	UserName   string   `json:"user_name" orm:"size(64);column(user_name)"`
	Content    string   `json:"content" orm:"size(64);column(content)"`
	CreateTime lib.Time `json:"create_time" orm:"auto_now_add;type(datetime);column(create_time)"`
}

type Article struct {
	ForumArticle dao.ForumArticle `json:"forum_article"`
	Content      string           `json:"content"`
	Comment      []Comment        `json:"comment"`
}

// 浏览文章 参数:[id]
func BrowseArticle(id int) (r Article, code string) {
	var comment []Comment
	// 获取帖子文章内容
	data, err := dao.SelecForumArticleById(id)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	bytes, err := ioutil.ReadFile(data.ContentPath)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	// 获取评论内容
	ArticleComments, err := dao.SelectForumCommentByArticleID(id)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	temp := Comment{}
	for _, value := range ArticleComments {
		user, err := dao.SelectWithUserMail(value.UserMail)
		if err != nil {
			logs.Error(err)
			temp.UserName = value.UserMail
		} else {
			temp.UserName = user.UserName
		}
		temp.Content = value.Content
		temp.CreateTime = value.CreateTime
		comment = append(comment, temp)
	}
	// 反转slice 逆序评论时间显示
	for i, j := 0, len(comment)-1; i < j; i, j = i+1, j-1 {
		comment[i], comment[j] = comment[j], comment[i]
	}
	r.Comment = comment
	r.ForumArticle = data
	r.Content = string(bytes)
	code = "2000"
	return
}

type ForumProperty struct {
	UserNum    int `json:"user_num"`
	ArticleNum int `json:"article_num"`
	CommentNum int `json:"conmment_num"`
}

// 获取论坛属性
func SelectForumProperty() (r ForumProperty, code string) {
	data, err := dao.SelectForumProperty()
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	r.ArticleNum = data.ArticleNum
	r.CommentNum = data.CommentNum
	r.UserNum = data.UserNum
	code = "2000"
	return
}

// 查询用户对当前id帖子是否点过赞
func SelectLikeFlag(article_id int, like_user_mail string) (r interface{}, code string) {
	if err := dao.SelectLikeByArticleAndMail(article_id, like_user_mail); err != nil {
		if err == orm.ErrNoRows {
			r = false
			code = "2000"
		} else {
			logs.Error(err)
			code = "5000"
		}
	} else {
		r = true
		code = "2000"
	}
	return
}

// 论坛帖子点赞操作
func Like(article_id int, like_user_mail string, flag bool) (code string) {
	if flag {
		if created, err := dao.ReadOrCreateLikeByArticleAndMail(article_id, like_user_mail); err != nil {
			logs.Error(err)
			code = "5000"
		} else {
			if created {
				if err = dao.UpdateLikeNum(article_id, 1); err != nil {
					logs.Error(err)
				}
				// 发送点赞的消息通知
				InsertMessage(article_id, like_user_mail, "like")
			}
			code = "2000"
		}
	} else {
		if num, err := dao.DeleteForumLikeByArticleIdAndMail(article_id, like_user_mail); err != nil {
			logs.Error(err)
			code = "5000"
		} else {
			if num > 0 {
				if err = dao.UpdateLikeNum(article_id, -1); err != nil {
					logs.Error(err)
				}
			}
			code = "2000"
		}
	}
	return
}

// 论坛帖子点赞操作
func Share(article_id int) (code string) {
	if err := dao.UpdateShareNum(article_id, 1); err != nil {
		logs.Error(err)
	}
	return "2000"
}

type Content struct {
	ForumArticle dao.ForumContent `json:"forum_article"`
	Content      string           `json:"content"`
}

// 查询论坛帖子编辑内容
func SelectForumContentByID(ArticleId int) (r Content, code string) {
	var err error
	r.ForumArticle, err = dao.SelecForumContentById(ArticleId)
	if err != nil {
		logs.Error(err)
		code = "5000"
		return
	}
	if bytes, err := ioutil.ReadFile(r.ForumArticle.ContentPath); err != nil {
		logs.Error(err)
		code = "5000"
		return
	} else {
		r.Content = string(bytes)
	}
	code = "2000"
	return
}

type UpdataContentReq struct {
	ArticleId      int    `json:"article_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	DisplayContent string `json:"display_content"`
}

// 更新帖子内容
func UpdataContent(req UpdataContentReq) (code string) {
	// 查询-删除文件-修改数据库信息
	if display, err := dao.SelecForumArticleById(req.ArticleId); err == nil {
		if err = removeFile(display.ContentPath); err != nil {
			logs.Error(err)
			code = "5000"
			return
		} else {
			if filePath, err := writeFileByArticleContent(req.DisplayContent); err != nil {
				code = "5000"
				return
			} else {
				if err = dao.UpdateDisplayTitleAndFilePath(req.ArticleId, req.Title, filePath); err != nil {
					logs.Error(err)
					code = "5000"
					return
				}
			}
		}
	} else {
		code = "5000"
		return
	}
	if content, err := dao.SelecForumContentById(req.ArticleId); err == nil {
		if err = removeFile(content.ContentPath); err != nil {
			logs.Error(err)
			code = "5000"
			return
		} else {
			if filePath, err := writeFileByArticleContent(req.Content); err != nil {
				logs.Error(err)
				code = "5000"
				return
			} else {
				if err = dao.UpdateContentTitleAndFilePath(req.ArticleId, req.Title, filePath); err != nil {
					logs.Error(err)
					code = "5000"
					return
				}
			}
		}
	} else {
		logs.Error(err)
		code = "5000"
		return
	}
	code = "2000"
	return
}

// 删除帖子内容所存储的文件 参数: [filepath]
func removeFile(filePath string) (error) {
	err := os.Remove(filePath);
	if err != nil {
		logs.Error(err);
	}
	return err
}

type CommentReq struct {
	Article_id int    `json:"article_id"`
	Mail       string `json:"mail"`
	Comment    string `json:"comment"`
}

// 为帖子增加一条评论
func InsertComment(r CommentReq) (code string) {
	var err error
	if err = dao.InsertForumComment(r.Article_id, r.Mail, r.Comment); err != nil {
		logs.Error(err)
		code = "5000"
		return
	} else {
		if err = dao.UpdateCommentNum(r.Article_id, 1); err != nil {
			logs.Error(err)
			if err = dao.DeleteForumComment(r.Article_id); err != nil {
				logs.Error(err)
			}
		} else {
			// 发送点赞的消息通知
			InsertMessage(r.Article_id, r.Mail, "comment")

			if err = dao.UpdateForumPropertyCommentNum(1); err != nil {
				logs.Error(err)
			}
		}
	}
	code = "2000"
	return
}

type StickyReq struct {
	Article_id int  `json:"article_id"`
	Flag       bool `json:"flag"`
}

// 给帖子置顶或取消置顶
func Sticky(r StickyReq) (code string) {
	if err := dao.UpdateDisplayStickyLogo(r.Article_id, r.Flag); err != nil {
		logs.Error(err)
		code = "5000"
	} else {
		code = "2000"
	}
	return
}

// 获取公告帖子的id主键
func GetAnnouncement() (interface{}, string) {
	data, err := dao.SelectAllForumArticleByAnnouncement()
	if err != nil {
		logs.Error(err)
		return "", "5000"
	} else {
		if len(data) != 0 {
			return data[0].Id, "2000"
		} else {
			return "", "2009"
		}
	}
}

type AnnouncementReq struct {
	Article_id int  `json:"article_id"`
	Flag       bool `json:"flag"`
}

// 更新公告
func UpdataAnnouncement(req AnnouncementReq) (code string) {
	err := dao.CancelAllForumArticleAnnouncement()
	if err != nil {
		logs.Error(err)
		return "5000"
	} else if req.Flag == true {
		if err = dao.SetForumArticleAnnouncement(req.Article_id); err != nil {
			logs.Error(err)
		}
	}
	return "2000"
}
