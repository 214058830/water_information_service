package controllers

import (
	"encoding/json"
	"strconv"

	"water_information_service/dao"
	"water_information_service/log"
	"water_information_service/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

// 通用响应格式
type userRes struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}

// @Title Get
// @查新所有的用户信息
// @Success 200 success
// @Failure 403 is empty
// @router / [get]
func (this *UserController) Get() {
	var r userRes
	r.Data, r.Code = models.SelectAllUser()
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type loginUserReq struct {
	PassWord string `form:"password"`
	Mail     string `form:"mail"`
	Single   bool   `form:"single"`
}

// @Title Post
// @登录
// @Success 200 success
// @Failure 403 is empty
// @router /login [post]
func (this *UserController) Login() {
	var r userRes
	var user loginUserReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &user); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Data, r.Code = models.SelectOneOfAllUser(user.Mail, user.PassWord)
	if r.Code == "2000" {
		this.setUserSession(user)
		// redis中保存后端鉴权token
		models.SetToken(this.Ctx.GetCookie("token"))
	} else {
		r.Data = ""
	}
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Post
// @注销登录
// @Success 200 success
// @Failure 403 is empty
// @router /logout [post]
func (this *UserController) Logout() {
	var r userRes
	// redis中保存后端鉴权token
	dao.DelToken(this.Ctx.GetCookie("token"))
	r.Code = "2000"
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// 设置session
func (this *UserController) setUserSession(user loginUserReq) {
	if user.Single {
		this.SetSession("mail", user.Mail)
		this.SetSession("password", user.PassWord)
	} else if !user.Single {
		this.DelSession("mail")
		this.DelSession("password")
	}
}

type UserSessionRes struct {
	Data loginUserReq `json:"data"`
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
}

// @Title Get
// @获取用户的session信息
// @Success 200 success
// @Failure 403 is empty
// @router /session [get]
func (this *UserController) GetUserSession() {
	var r UserSessionRes
	if (this.GetSession("mail") == nil) {
		r.Code = "2010"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Data.Mail = this.GetSession("mail").(string)
	r.Data.PassWord = this.GetSession("password").(string)
	r.Data.Single = true
	r.Code = "2000"
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// 传入的数据格式
type registerUserInfoReq struct {
	Name     string `form:"name"`
	Mail     string `form:"mail"`
	PassWord string `form:"password"`
}

// @Title Post
// @注册
// @Success 200 success
// @Failure 403 is empty
// @router /register [post]
func (this *UserController) Register() {
	var r userRes
	var userInfo registerUserInfoReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &userInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Data, r.Code = models.RegisterUserInfo(userInfo.Name, userInfo.Mail, userInfo.PassWord)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type updataUserReq struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
	Id       string `form:"id"`
}

// @Title Post
// @个人中心修改用户信息
// @Success 200 success
// @Failure 403 is empty
// @router /updataUserInfo [post]
func (this *UserController) UpdataUserInfo() {
	var r userRes
	var userInfo updataUserReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &userInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	id, _ := strconv.Atoi(userInfo.Id)
	r.Code = models.UpdataUserInfo(id, userInfo.UserName, userInfo.PassWord)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

type updataAdminReq struct {
	Id   int  `form:"id"`
	Bool bool `form:"bool"`
}

// @Title Post
// @管理员权限设置
// @Success 200 success
// @Failure 403 is empty
// @router /updataAdmin [post]
func (this *UserController) UpdataAdmin() {
	var r userRes
	var userInfo updataAdminReq
	body := this.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &userInfo); err != nil {
		logs.Error(err)
		r.Code = "5000"
		r.Msg = log.CodeMap[r.Code]
		this.Data["json"] = r
		this.ServeJSON()
		return
	}
	r.Code = models.UpdataAdmin(userInfo.Id, userInfo.Bool);
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}
