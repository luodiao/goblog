package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "11"
	c.Data["Email"] = "count"
	c.TplName = "index.tpl"
}
