package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

type emailInfo struct {
	FromEmail string `form:"fromEmail"`
	AliasName string `form:"aliasName"`
	SmtpServer string `form:"smtpServer"`
	Port string `form:"port"`
	Password string `form:"password"`
	VarField string `form:"varField"`
	DelayTime []int `form:"delayTime[]"`
}

type fileInfo struct {
	VarFileName string
	SubjectFileName string
	MessageFileName string
}

func (c *MainController) Get() {
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
	c.Render()
}
