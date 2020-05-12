package log

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var CodeMap map[string]string

func init() {
	runmode := beego.AppConfig.String("runmode")
	CodeMap = make(map[string]string)
	CodeMap = map[string]string{
		"2000": "ok",
		"2001": "用户名或密码错误",
		//"2002": "更新用户信息出错，请稍后重试",
		"2003": "此邮箱已注册用户",
		//"2004": "变动管理员权限失败，请重试",
		"2005": "该邮箱尚未注册",
		"2006": "该单位信息已被录入",
		"2007": "单位详情不存在",
		"2008": "token过期，请重新登录",
		"2009": "没有新的公告",
		"2010": "没有对应的session信息",
		"4000": "没有对应的请求信息",
		"5000": "服务器出错",
		"5001": "解析参数出错",
	}
	logs.EnableFuncCallDepth(true) // 允许打印文件名 + 行号
	logs.SetLogFuncCallDepth(3)    // 设置层级
	logs.Async()                   // 异步输出日志
	logs.Async(1e3)
	if (runmode == "prod") {
		err := logs.SetLogger(logs.AdapterMultiFile, `{"filename":"log/test.log","separate":["error", "info"]}`)
		if err != nil {
			logs.Error(err)
		}
	}

	logs.Info("Init log success...")
}
