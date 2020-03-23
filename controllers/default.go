package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
	template string
}

// Prepare 渲染前调用方法
func (c *MainController) Prepare() {
	// 默认default模板
	c.template = "default"

	// 布局模板
	c.Layout = "template/" + c.template + "/layout.tpl"
}

// Index 首页
func (c *MainController) Index() {
	c.TplName = "template/" + c.template + "/index.tpl"
}

// Archives 归档
func (c *MainController) Archives() {
	c.TplName = "template/" + c.template + "/archives.tpl"
}
