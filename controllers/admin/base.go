package admin

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	controller string
	action     string
}

func (p *baseController) Prepare() {
	p.Layout = "admin/layout.tpl"
	p.Data["_xsrf"] = p.XSRFToken()

	controller, action := p.GetControllerAndAction()
	p.controller = strings.ToLower(controller[0 : len(controller)-10])
	p.action = strings.ToLower(action)
}

func getResponse(code int, msg string, data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["msg"] = msg
	response["data"] = data

	return response
}

func getCurrentDate() string {

	return time.Now().Format("2006-01-02 15:04:05")
	// current := time.Now()
	// year := current.Year()
	// month := current.Month()
	// day := current.Day()
	// hour := current.Hour()
	// minute := current.Minute()
	// second := current.Second()

	// return string(year) + "-" + string(month) + "-" + string(day) + " " + string(hour) + ":" + string(minute) + ":" + string(second)
}
