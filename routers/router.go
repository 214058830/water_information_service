// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"encoding/json"

	"water_information_service/controllers"
	"water_information_service/dao"
	"water_information_service/log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type res struct {
	Data interface{} `json:"data"`
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
}

// 前置路由过滤器
var UrlManager = func(ctx *context.Context) {
	token := ctx.Input.Cookie("token")
	if resToken := dao.GetToken(token); resToken != "1" {
		var r res
		r.Code = "2008"
		r.Msg = log.CodeMap[r.Code]
		data, _ := json.Marshal(r)
		ctx.ResponseWriter.Write(data)
	}
}

func init() {
	beego.InsertFilter("/v1/filter/*", beego.BeforeRouter, UrlManager)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/water",
			beego.NSInclude(
				&controllers.WaterInformationController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/forum",
			beego.NSInclude(
				&controllers.ForumController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSInclude(
				&controllers.MessageController{},
			),
		),
		// 需要过滤的路由
		beego.NSNamespace("/filter",
			beego.NSNamespace("/water",
				beego.NSInclude(
					&controllers.WaterInformationController{},
				),
			),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&controllers.UserController{},
				),
			),
			beego.NSNamespace("/forum",
				beego.NSInclude(
					&controllers.ForumController{},
				),
			),
			beego.NSNamespace("/message",
				beego.NSInclude(
					&controllers.MessageController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
