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

	controller, action := p.GetControllerAndAction()
	p.controller = strings.ToLower(controller[0 : len(controller)-10])
	p.action = strings.ToLower(action)

	p.Data["_xsrf"] = p.XSRFToken()
	p.Data["currentUser"] = p.GetSession("user")

	// 验证是否登录
	allowURI := []string{"admin/login", "admin/ajaxLogin", "admin/register", "admin/ajaxRegister"}
	if !inArray(allowURI, p.controller+"/"+p.action) && p.GetSession("user") == nil {
		p.Ctx.Redirect(302, "/admin/login")
	}

	// 登录访问allow路由
	if inArray(allowURI, p.controller+"/"+p.action) && p.GetSession("user") != nil {
		p.Ctx.Redirect(302, "/admin")
	}
}

func inArray(array []string, currentURI string) bool {
	for _, value := range array {
		if strings.ToLower(value) == currentURI {
			return true
		}
	}

	return false
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
}
