package admin

type IndexController struct {
	baseController
}

func (c *IndexController) Index() {
	c.Layout = "admin/layout.tpl"

	c.TplName = "admin/" + c.controller + "/index.html"
}
