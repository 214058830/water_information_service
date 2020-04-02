package main

import (
	_ "water_information_service/api"
	_ "water_information_service/routers"
	_ "water_information_service/log"
	_ "water_information_service/dao"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
