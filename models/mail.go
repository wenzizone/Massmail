package models

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"os"
	"bufio"
	//"github.com/sadlil/gologger"
	"net/smtp"
	"crypto/tls"
	"bytes"
	"github.com/astaxie/beego"
	"html/template"
	"math/rand"
	//"net/mail"
)

type FileInfo struct {
	VarFile string
	SubjectFile string
	MessageFile string
}

type Config struct {
	VarField string
	DelayTime []int
}

type SmtpServer struct {
	Host string
	Port string
}

type Mail struct {
	FileInfo
	Config
	SmtpServer
	SenderId	string
	SenderAlias string
	ToIds		string
	Subject		string
	Body     	string
	Password 	string
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	if mail.SenderAlias != "" {
		message += fmt.Sprintf("From: %s<%s>\r\n",mail.SenderAlias, mail.SenderId)
	} else {
		message += fmt.Sprintf("From: %s\r\n", mail.SenderId)
	}

	message += fmt.Sprintf("To: %s\r\n", mail.ToIds)

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}

func (m *Mail) SendEmail() {
	beego.Debug(m)

	f,err := os.Open(m.VarFile)
	if err != nil {
		beego.Debug(err)
	}
	defer f.Close()

	varsMap := make(map[interface{}]interface{})

	varFieldSlice := strings.Split(m.VarField, ",")

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		time.Sleep(time.Duration(m.generateTandomSleeptime()) * time.Second)
		//time.Sleep(10 * time.Second)
		//fmt.Println(scanner.Text())

		varFieldContentSlice := strings.Split(scanner.Text(),",")

		for _, i := range varFieldSlice {
			ii, _ := strconv.Atoi(i)
			varsMap["Var"+i] = varFieldContentSlice[ii]
		}

		subject := m.generateTitle(varsMap)
		emailBody := m.generateEmailMessage(varsMap)
		m.sendingEmail(varFieldContentSlice[0], subject, emailBody)
	}
}

// sending email
func (m *Mail) sendingEmail(toEmail string, subject string, body string) {
	//loger := gologger.GetLogger(gologger.FILE, path.Join(*logDir,"mail.log"))

	//mail.senderId = *from
	//mail.senderAlias = *alias
	m.ToIds = toEmail
	m.Subject = subject
	m.Body = body

	messageBody := m.BuildMessage()

	beego.Info(m.Host)
	//build an auth
	auth := smtp.PlainAuth("", m.SenderId, m.Password, m.Host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         m.Host,
	}

	conn, err := tls.Dial("tcp", m.ServerName(), tlsconfig)
	if err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	client, err := smtp.NewClient(conn, m.Host)
	if err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	// step 2: add all from and to
	if err = client.Mail(m.SenderId); err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}
	if err = client.Rcpt(m.ToIds); err != nil {
		beego.Error(fmt.Sprintf("%s 地址有问题，%s", m.ToIds, err))
		return
	}

	// Data
	w, err := client.Data()
	if err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	err = w.Close()
	if err != nil {
		beego.Debug(fmt.Sprintf("%v",err))
		return
	}

	client.Quit()

	beego.Info(fmt.Sprintf("邮件发送成功,%s", m.ToIds))
}

// load title template file and return title content
func (f *FileInfo) generateTitle(varFieldMap map[interface{}]interface{}) string {
	var doc bytes.Buffer

	tmpl := template.Must(template.ParseFiles(f.SubjectFile))
	//tmpl, _ := template.ParseFiles(*titleFile)

	err := tmpl.Execute(&doc, varFieldMap)
	if err != nil {
		beego.Debug(fmt.Sprint("template execution: %s", err))
	}
	s := doc.String()
	return s
}

func (f *FileInfo) generateEmailMessage(varFieldMap map[interface{}]interface{}) string {
	var doc bytes.Buffer

	tmpl := template.Must(template.ParseFiles(f.MessageFile))

	err := tmpl.Execute(&doc, varFieldMap)
	if err != nil {
		beego.Debug(fmt.Sprint("template execution: %s", err))
	}
	s := doc.String()
	return s
}

func (c *Config) generateTandomSleeptime() int {
	//delayArray := strings.Split(*delayTime, ",")
	rand.Seed(time.Now().Unix())
	min, max := c.DelayTime[0], c.DelayTime[1]
	//max, _ := config.delayTime[1]
	//return 1
	randTimedelay := rand.Intn(max-min)+min
	//fmt.Println(randTimedelay)
	return randTimedelay
}