package routers

import (
	"Massmail/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/static/*", &controllers.MainController{},"get:GetAsset")
	beego.Router("/", &controllers.MainController{},"get:GetHome")
	beego.Router("/api/upload/:filetype", &controllers.ApiController{}, "post:Upload")
	beego.Router("/api/sendmail", &controllers.ApiController{}, "post:SendEmail")
}
