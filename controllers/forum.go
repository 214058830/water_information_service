package controllers

import (
	"encoding/json"

	"water_information_service/log"
	"water_information_service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ForumController struct {
	beego.Controller
}

// 通用响应格式
type forumRes struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}

// @Title Get
// @获取所有论坛帖子列表
// @Success 200 success
// @Failure 403 is empty
// @router / [get]
func (this *ForumController) Get() {
	var r forumRes
	r.Data, r.Code = models.SelectAllForumArticle()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type getForumArticleByMailReq struct {
	Mail string `json:"mail"`
}

// @Title Post
// @获取某个用户的论坛帖子列表
// @Success 200 success
// @Failure 403 is empty
// @router / [post]
func (this *ForumController) Post() {
	var r forumRes
	var req getForumArticleByMailReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Data, r.Code = models.SelectAllForumArticleByMail(req.Mail)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @查看某一篇帖子内容
// @Success 200 success
// @Failure 403 is empty
// @router /browse [get]
func (this *ForumController) Browse() {
	var r forumRes
	id, _ := this.GetInt("id")
	r.Data, r.Code = models.BrowseArticle(id)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type releaseForumArticleReq struct {
	Title          string `json:"title"`
	Content        string `json:"content"`
	DisplayContent string `json:"display_content"`
	Username       string `json:"username"`
	Mail           string `json:"mail"`
}

// @Title Post
// @发布帖子
// @Success 200 success
// @Failure 403 is empty
// @router /release [post]
func (this *ForumController) Release() {
	var r forumRes
	var req releaseForumArticleReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.ReleaseForumArticle(req.Mail, req.Username, req.Title, req.Content, req.DisplayContent)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @查询论坛属性
// @Success 200 success
// @Failure 403 is empty
// @router /property [get]
func (this *ForumController) GetForumProperty() {
	var r forumRes
	r.Data, r.Code = models.SelectForumProperty()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type LikeReq struct {
	ArticleId    int    `json:"article_id"`
	LikeUserMail string `json:"like_user_mail"`
	Flag         bool   `json:"flag"`
}

// @Title POST
// @论坛帖子赞操作
// @Success 200 success
// @Failure 403 is empty
// @router /like [post]
func (this *ForumController) Like() {
	var r forumRes
	var req LikeReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.Like(req.ArticleId, req.LikeUserMail, req.Flag)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @获取当前用户对帖子是否点赞
// @Success 200 success
// @Failure 403 is empty
// @router /like_flag [get]
func (this *ForumController) SelectLikeFlag() {
	var r forumRes
	articleId, _ := this.GetInt("article_id")
	userMail := this.GetString("like_user_mail")
	r.Data, r.Code = models.SelectLikeFlag(articleId, userMail)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type GetContentByIdReq struct {
	ArticleId int `json:"article_id"`
}

// @Title Post
// @获取某篇论坛帖子编辑内容
// @Success 200 success
// @Failure 403 is empty
// @router /contentById [post]
func (this *ForumController) GetContentById() {
	var r forumRes
	var req GetContentByIdReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Data, r.Code = models.SelectForumContentByID(req.ArticleId)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @更新编辑后的帖子内容
// @Success 200 success
// @Failure 403 is empty
// @router /edit [post]
func (this *ForumController) UpdataContent() {
	var r forumRes
	var req models.UpdataContentReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UpdataContent(req)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @添加评论
// @Success 200 success
// @Failure 403 is empty
// @router /comment [post]
func (this *ForumController) Comment() {
	var r forumRes
	var req models.CommentReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.InsertComment(req)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @帖子置顶/取消置顶
// @Success 200 success
// @Failure 403 is empty
// @router /sticky [post]
func (this *ForumController) Sticky() {
	var r forumRes
	var req models.StickyReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.Sticky(req)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @获取公告
// @Success 200 success
// @Failure 403 is empty
// @router /announcement [get]
func (this *ForumController) GetAnnouncement() {
	var r forumRes
	r.Data, r.Code = models.GetAnnouncement()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @设置公告
// @Success 200 success
// @Failure 403 is empty
// @router /set_announcement [post]
func (this *ForumController) SetAnnouncement() {
	var r forumRes
	var req models.AnnouncementReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UpdataAnnouncement(req)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}
