package admin

import (
	"goblog/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/gurudin/pinyin"
)

// CategoryController 类别控制器
type CategoryController struct {
	baseController
}

// Index 类别管理
func (c *CategoryController) Index() {
	query := make(map[string]string)

	var fields = []string{"Id", "Category", "Weight", "Alias", "Color", "Remark", "Enabled", "UpdatedAt"}

	var sortby = []string{"Id"}

	var order = []string{"desc"}

	list, err := models.GetAllCategory(query, fields, sortby, order, 0, 20)
	if err != nil {
		beego.Info("err")
	}

	c.Data["list"] = list
	c.Data["total"] = len(list)
	c.TplName = "admin/" + c.controller + "/index.html"
}

// Save 新增 or 修改
func (c *CategoryController) Save() {
	id := c.Ctx.Input.Param(":id")

	if id == "" {
		m := models.Category{}
		m.Color = "#000000"

		c.Data["m"] = m
	} else {
		id64, _ := strconv.ParseInt(id, 10, 64)
		m, _ := models.GetCategoryById(id64)
		c.Data["m"] = m
	}

	c.TplName = "admin/" + c.controller + "/save.html"
}

// AjaxSave 新增 or 修改
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

	category.Alias = pinyin.Permalink(category.Category, "")
	if category.Id == 0 {
		category.CreatedAt = getCurrentDate()
		_, err := models.AddCategory(&category)
		if err != nil {
			c.Data["json"] = getResponse(-1, "新增失败", "")
			c.ServeJSON()
			return
		}
	} else {
		err := models.UpdateCategoryById(&category)
		if err != nil {
			c.Data["json"] = getResponse(-1, "修改失败", "")
			c.ServeJSON()
			return
		}
	}

	c.Data["json"] = getResponse(0, "success", "")
	c.ServeJSON()
	return
}

// AjaxRemove 删除类别
func (c *CategoryController) AjaxRemove() {

	beego.Info(c.GetString("id"))

	id64, _ := c.GetInt64("id", 10)
	err := models.DeleteCategory(id64)
	if err != nil {
		c.Data["json"] = getResponse(-1, "删除失败", nil)
	}

	c.Data["json"] = getResponse(0, "success", "")
	c.ServeJSON()
	return
}
