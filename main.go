package main

import (
	_ "Massmail/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.AppName = "Massmail"
	beego.BConfig.Listen.HTTPPort = 8081
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RunMode = "prod"
	//beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "2ck7PNoQue4uyXERWwji9hrCCveyUc"
	beego.BConfig.WebConfig.XSRFExpire = 3600

	beego.Run()
}
