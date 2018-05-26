package controllers

import (
	"Massmail/models"
	"github.com/adam-hanna/randomstrings"
	"github.com/astaxie/beego"
	"path"
)

type ApiController struct {
	beego.Controller
}

// @router /api/uploadVarfile [post]
func (this *ApiController) Upload() {
	beego.Debug(this.Ctx.Request)
	beego.Debug(this.Ctx.Input.Param(":filetype"))
	fileType := this.Ctx.Input.Param(":filetype")
	switch fileType {
	case "varfile":
		varFileName, _ := randomstrings.GenerateRandomString(9)
		this.SaveToFile("varFile", path.Join("/tmp", varFileName))
		this.Data["json"] = map[string]string{"fileType": "varfile", "fileName": varFileName}
	case "subjectile":
		subjectFileName, _ := randomstrings.GenerateRandomString(9)
		this.SaveToFile("subjectFile", path.Join("/tmp", subjectFileName))
		this.Data["json"] = map[string]string{"fileType": "subjectfile", "fileName": subjectFileName}
	case "msgfile":
		msgFileName, _ := randomstrings.GenerateRandomString(9)
		this.SaveToFile("msgFile", path.Join("/tmp", msgFileName))
		this.Data["json"] = map[string]string{"fileType": "msgfile", "fileName": msgFileName}
	}

	//beego.Debug(varFileName)
	this.ServeJSON()
}

func (this *ApiController) SendEmail() {
	//beego.Info("%v", this.Input())
	mail := models.Mail{}
	if err := this.ParseForm(&mail); err != nil {
		beego.Debug(err)
	}
	beego.Info("%v", mail)
	/*
	   mh := models.Mail{
	     FileInfo:models.FileInfo{
	       VarFile: path.Join("/tmp", f.VarFileName),
	       SubjectFile: path.Join("/tmp", f.SubjectFileName),
	       MessageFile: path.Join("/tmp", f.MessageFileName),
	     },
	     Config:models.Config{
	       VarField:  e.VarField,
	       DelayTime: e.DelayTime,
	     },
	     SmtpServer:models.SmtpServer{
	       Host: e.SmtpServer,
	       Port: e.Port,
	     },
	     SenderId: e.FromEmail,
	     SenderAlias: e.AliasName,
	     Password: e.Password,
	   }

	   go mh.SendEmail()
	*/
	go mail.SendEmail()
}
