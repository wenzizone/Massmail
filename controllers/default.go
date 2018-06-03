package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"Massmail/data"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) GetHome() {
	//c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	datas, err := data.Asset("views/index.tpl")
	if err != nil {
		//c.String(400, err.Error())
		return
	}
	c.Ctx.Output.Body(datas)
}

func (c *MainController) GetAsset() {
	splat := c.Ctx.Input.Param(":splat")
	datas, err := data.Asset("static/" + splat)
	if err != nil {
		//c.String(400, err.Error())
		return
	}
	c.Ctx.Output.Header("Content-Type",assetContentType(splat))
	c.Ctx.Output.Body(datas)
}