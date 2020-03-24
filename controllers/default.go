package controllers

import (
	"goblog/models"
	"math/rand"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

var globalCategoryList []interface{}

type MainController struct {
	beego.Controller
	template string
	limit    int64
}

// SiteConfig 站点配置
type SiteConfig struct {
	Title    string
	Name     string
	Avatar   string
	Location string
	Navs     map[int]map[string]string
}

// Prepare 渲染前调用方法
func (c *MainController) Prepare() {
	// 默认default模板
	c.template = "default"

	// 默认分页
	c.limit = 20

	// 布局模板
	c.Layout = "template/" + c.template + "/layout.tpl"

	// 站点配置
	c.Data["site"] = getSiteConfig()

	// 获取类别
	categoryList, _ := models.GetAllCategory(
		map[string]string{"Enabled": "1"},
		[]string{"Id", "Category", "Alias"},
		[]string{"Id"},
		[]string{"desc"},
		0,
		100,
	)
	globalCategoryList = categoryList
	c.Data["categoryList"] = categoryList

	// 最新文章
	articleTop5, total, _ := models.GetAllArticle(
		map[string]string{"Status": "1"},
		[]string{"Id", "FkCategoryId", "Title", "Remark", "UpdatedAt"},
		[]string{"Weight", "Id"},
		[]string{"desc", "desc"},
		0,
		5,
	)
	c.Data["articleTop5"] = articleTop5

	// 文章总数
	c.Data["total"] = total
}

// Index 首页
func (c *MainController) Index() {
	p, _ := c.GetInt64("p", 1)
	alias := c.GetString("alias", "")
	categoryID := c.GetString("category", "")

	// 获取文章
	var offset int64
	offset = (p - 1) * c.limit

	query := map[string]string{}
	query["Status"] = "1"
	if alias != "" {
		query["FkCategoryId"] = strconv.FormatInt(getCategoryIdByAlias(alias), 10)
	}
	if categoryID != "" {
		query["FkCategoryId"] = categoryID
	}

	articleList, _, _ := models.GetAllArticle(
		query,
		[]string{"Id", "FkCategoryId", "Title", "Remark", "UpdatedAt"},
		[]string{"Weight", "Id"},
		[]string{"desc", "desc"},
		offset,
		c.limit,
	)
	c.Data["p"] = p
	c.Data["articleList"] = articleList

	c.TplName = "template/" + c.template + "/index.tpl"
}

// Detail 文章详情
func (c *MainController) Detail() {
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id, 10, 64)
	article, _ := models.GetArticleById(id64)
	c.Data["article"] = article

	c.TplName = "template/" + c.template + "/detail.tpl"
}

// Archives 归档
func (c *MainController) Archives() {
	articleList, _, _ := models.GetAllArticle(
		map[string]string{"Status": "1"},
		[]string{"Id", "FkCategoryId", "Title", "Remark", "UpdatedAt"},
		[]string{"UpdatedAt"},
		[]string{"desc"},
		0,
		100,
	)
	c.Data["articleList"] = articleList

	c.TplName = "template/" + c.template + "/archives.tpl"
}

// Tags 标签
func (c *MainController) Tags() {
	tags := make(map[string]map[string]string)
	colors := [7]string{
		"",
		"layui-bg-orange",
		"layui-bg-green",
		"layui-bg-cyan",
		"layui-bg-blue",
		"layui-bg-black",
		"layui-bg-gray",
	}

	// 文章tags
	articleList, _, _ := models.GetAllArticle(
		map[string]string{"Status": "1"},
		[]string{"Id", "Tags"},
		[]string{"UpdatedAt"},
		[]string{"desc"},
		0,
		50,
	)
	for _, article := range articleList {
		tmp := article.(map[string]interface{})
		if tmp["Tags"].(string) != "" {
			s := strings.Split(tmp["Tags"].(string), ",")
			for i := 0; i < len(s); i++ {
				tags[s[i]] = map[string]string{
					"url":   "/?tag=" + s[i],
					"color": colors[rand.Intn(7)],
				}
			}
		}
	}

	// 分类tags
	for _, category := range globalCategoryList {
		tmp := category.(map[string]interface{})
		tags[tmp["Category"].(string)] = map[string]string{
			"url":   "/?alias=" + tmp["Alias"].(string),
			"color": colors[rand.Intn(7)],
		}
	}

	c.Data["tags"] = tags

	c.TplName = "template/" + c.template + "/tags.tpl"
}

func getSiteConfig() SiteConfig {
	navs := make(map[int]map[string]string)
	navs[0] = map[string]string{"title": "主页", "url": "/"}
	navs[1] = map[string]string{"title": "归档", "url": "/archives"}
	navs[2] = map[string]string{"title": "标签", "url": "/tags"}
	navs[3] = map[string]string{"title": "关于", "url": "/about"}
	navs[4] = map[string]string{"title": "GitHub", "url": "https://github.com/Gurudin"}

	return SiteConfig{
		Title:    "古鲁丁 | 博客",
		Name:     "Gurudin",
		Avatar:   "https://avatars1.githubusercontent.com/u/29497772?s=460&u=06d328ff03b80422754b1fd62eb5051d17aac2a6&v=4",
		Location: "Bei jing",
		Navs:     navs,
	}
}

func getCategoryIdByAlias(alias string) int64 {
	for _, category := range globalCategoryList {
		tmp := category.(map[string]interface{})
		if tmp["Alias"].(string) == alias {
			return tmp["Id"].(int64)
		}
	}

	return 0
}

func GetCategory(id int64) string {
	for _, category := range globalCategoryList {
		tmp := category.(map[string]interface{})
		if tmp["Id"] == id {
			return tmp["Category"].(string)
		}
	}

	return ""
}
