package admin

import (
	"goblog/models"
	"math/rand"
	"strconv"
)

// ArticleController 文章控制器
type ArticleController struct {
	baseController
}

// Index 文章列表
func (c *ArticleController) Index() {
	p, _ := c.GetInt64("p", 1)

	// 类别
	categoryQuery := make(map[string]string)
	categoryQuery["Enabled"] = "1"
	categoryList, _ := models.GetAllCategory(categoryQuery, []string{"Id", "Category"}, []string{"Id"}, []string{"desc"}, 0, 100)

	// 文章
	var offset, limit int64
	limit = 20
	offset = (p - 1) * limit
	query := make(map[string]string)
	sortby := []string{"Weight", "Id"}
	order := []string{"desc", "desc"}
	articleList, total, _ := models.GetAllArticle(query, nil, sortby, order, offset, limit)

	c.Data["categoryList"] = categoryList
	c.Data["articleList"] = articleList
	c.Data["total"] = total
	c.Data["p"] = p
	c.Data["limit"] = limit
	c.TplName = "admin/" + c.controller + "/index.html"
}

// Save 文章设置
func (c *ArticleController) Save() {
	id := c.Ctx.Input.Param(":id")

	// 获取类别
	query := make(map[string]string)
	query["Enabled"] = "1"
	var fields = []string{"Id", "Category", "Color", "Alias"}
	var sortby = []string{"Weight", "Id"}
	var order = []string{"desc", "desc"}
	categoryList, _ := models.GetAllCategory(query, fields, sortby, order, 0, 20)
	c.Data["categoryList"] = categoryList

	if id == "" {
		m := models.Article{}
		m.Pv = rand.Intn(1000)

		c.Data["m"] = m
	} else {
		id64, _ := strconv.ParseInt(id, 10, 64)
		m, _ := models.GetArticleById(id64)

		c.Data["m"] = m
	}

	// c.Data["m"] = m
	c.TplName = "admin/" + c.controller + "/save.html"
}

// AjaxSave 文章设置 create or update
func (c *ArticleController) AjaxSave() {
	article := models.Article{}
	if err := c.ParseForm(&article); err != nil {
		return
	}

	if article.Id == 0 {
		// create
		article.CreatedAt = getCurrentDate()
		_, err := models.AddArticle(&article)
		if err != nil {
			c.Data["json"] = getResponse(-1, "新增失败", nil)
			c.ServeJSON()
		}
	} else {
		// update
		err := models.UpdateArticleById(&article)
		if err != nil {
			c.Data["json"] = getResponse(-1, "修改失败", nil)
			c.ServeJSON()
		}
	}

	c.Data["json"] = getResponse(0, "success", nil)
	c.ServeJSON()
	return
}

// AjaxRemove 删除文章
func (c *ArticleController) AjaxRemove() {
	id, _ := c.GetInt64("id", 10)
	err := models.DeleteArticle(id)
	if err != nil {
		c.Data["json"] = getResponse(-1, "删除失败", nil)
	}

	c.Data["json"] = getResponse(0, "success", "")
	c.ServeJSON()
	return
}
