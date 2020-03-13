package admin

import (
	"strings"

	"github.com/astaxie/beego"
)

type baseController struct {
	beego.Controller
	controller string
	action     string
}

type params struct {
	Action string      `form:"action"`
	Model  interface{} `form:"data"`
}

func (p *baseController) Prepare() {
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
