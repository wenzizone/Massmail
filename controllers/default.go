package controllers

import (
	"github.com/astaxie/beego"
	//"html/template"
	"github.com/adam-hanna/randomstrings"
	"path"
	"MassMail/models"
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
	//c.Data["xsrfdata"]=template.HTML(c.XSRFFormHTML())
	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) Post() {
	e := emailInfo{}
	f := fileInfo{}

	// generate random file name
	f.VarFileName, _ = randomstrings.GenerateRandomString(9)
	f.SubjectFileName, _ = randomstrings.GenerateRandomString(9)
	f.MessageFileName, _ = randomstrings.GenerateRandomString(9)

	// save var file to tmp folder
	c.SaveToFile("varFile", path.Join("/tmp",f.VarFileName))
	c.SaveToFile("subjectFile", path.Join("/tmp",f.SubjectFileName))
	c.SaveToFile("messageFile", path.Join("/tmp",f.MessageFileName))

	if err := c.ParseForm(&e); err!= nil {
		beego.Debug(err)
	}
	beego.Info("%v", e)

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
	/*
	conf.SetValue(e.VarField,e.DelayTime)

	mail := models.Mail{}
	mail.SetValue(e.FromEmail, e.AliasName)

	smtpServer := models.SmtpServer{}
	smtpServer.SetValue(e.SmtpServer, e.Port, e.Password)

	fileInfo := models.FileInfo{}
	fileInfo.SetValue(path.Join("/tmp",f.VarFileName), path.Join("/tmp",f.SubjectFileName), path.Join("/tmp",f.MessageFileName))
	*/

	go mh.SendEmail()
}
