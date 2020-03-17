package admin

import (
	"goblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type CategoryController struct {
	baseController
}

func (c *CategoryController) Index() {

	c.TplName = "admin/" + c.controller + "/index.html"
}

func (c *CategoryController) Save() {
	id := c.Ctx.Input.Param(":id")

	if id == "" {
		m := models.Category{}
		m.Color = "#000000"

		c.Data["m"] = m
	} else {
		beego.Info("no")
	}

	c.TplName = "admin/" + c.controller + "/save.html"
}

func (c *CategoryController) AjaxSave() {
	category := models.Category{}
	if err := c.ParseForm(&category); err != nil {
		return
	}

	// valid
	valid := validation.Validation{}
	valid.Required(category.Category, "Category")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.Data["json"] = getResponse(-1, err.Key+": "+err.Message, "")
			c.ServeJSON()
			return
		}
	}

	category.CreatedAt = getCurrentDate()
	beego.Info(getCurrentDate())
	_, err := models.AddCategory(&category)
	if err != nil {
		c.Data["json"] = getResponse(0, "注册失败", "")
		c.ServeJSON()
		return
	}

	c.Data["json"] = getResponse(0, "success", "")
	c.ServeJSON()
	return
}
