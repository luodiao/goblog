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

func (p *baseController) Prepare() {
	controller, action := p.GetControllerAndAction()
	p.controller = strings.ToLower(controller[0 : len(controller)-10])
	p.action = strings.ToLower(action)
}
