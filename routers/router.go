package routers

import (
	"goblog/controllers"
	"goblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	/* Web */
	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/detail/:id", &controllers.MainController{}, "get:Detail")
	beego.Router("/archives", &controllers.MainController{}, "get:Archives")
	beego.Router("/tags", &controllers.MainController{}, "get:Tags")

	/* Admin */
	// login
	beego.Router("/admin/login", &admin.AdminController{}, "*:Login")
	beego.Router("/admin/ajaxLogin", &admin.AdminController{}, "post:AjaxLogin")
	beego.Router("/admin/logout", &admin.AdminController{}, "get:Logout")

	// register
	beego.Router("/admin/register", &admin.AdminController{}, "*:Register")
	beego.Router("/admin/ajaxRegister", &admin.AdminController{}, "post:AjaxRegister")

	// upload
	beego.Router("/admin/upload", &admin.AdminController{}, "post:Upload")

	// index
	beego.Router("/admin", &admin.IndexController{}, "get:Index")

	// category
	beego.Router("/admin/category", &admin.CategoryController{}, "get:Index")
	beego.Router("/admin/category/save/?:id", &admin.CategoryController{}, "get:Save")
	beego.Router("/admin/category/save", &admin.CategoryController{}, "post:AjaxSave")
	beego.Router("/admin/category/remove", &admin.CategoryController{}, "post:AjaxRemove")

	// article
	beego.Router("/admin/article", &admin.ArticleController{}, "get:Index")
	beego.Router("/admin/article/save/?:id", &admin.ArticleController{}, "get:Save")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:AjaxSave")
	beego.Router("/admin/article/remove", &admin.ArticleController{}, "post:AjaxRemove")
}
