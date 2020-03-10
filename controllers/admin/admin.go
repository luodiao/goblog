package admin

import "fmt"

type AdminController struct {
	baseController
}

// 登录
func (c *AdminController) Login() {
	fmt.Println(c.controller, c.action)
	c.TplName = "index.tpl"
}
