package controllers

import (
	"github.com/astaxie/beego"
	"water_information_service/log"
	"water_information_service/models"
)

type MessageController struct {
	beego.Controller
}

// 通用响应格式
type messageRes struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}

// @Title Get
// @查询用户的某种类型所有消息
// @Success 200 success
// @Failure 403 is empty
// @router / [get]
func (this *MessageController) GetAllMessage() {
	var r messageRes
	mail := this.GetString("mail")
	messageType := this.GetString("type")
	r.Data, r.Code = models.GetAllMessage(mail, messageType)
	models.UpdataReadMessageNum(mail, messageType)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}

// @Title Get
// @查询用户的未读消息数量
// @Success 200 success
// @Failure 403 is empty
// @router /number [get]
func (this *MessageController) GetMessageNum() {
	var r messageRes
	mail := this.GetString("mail")
	r.Data, r.Code = models.GetMessageNum(mail)
	r.Msg = log.CodeMap[r.Code]
	this.Data["json"] = r
	this.ServeJSON()
}
