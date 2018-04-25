package routers

import (
	"Massmail/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/upload/:filetype", &controllers.ApiController{}, "post:Upload")
    beego.Router("/api/sendmail",&controllers.ApiController{},"post:SendEmail")
}
