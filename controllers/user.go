package controllers

import (
    "github.com/astaxie/beego"
    "goblog/models"
)

type UserController struct {
    beego.Controller
}

func (c *UserController) Get() {
    var user models.Users
    count, _ := user.Query().Count()
    c.Data["Website"] = count
    c.Data["Email"] = count
    c.TplName = "index.tpl"
}
