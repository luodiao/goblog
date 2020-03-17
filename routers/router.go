package routers

import (
	"goblog/controllers"
	"goblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	/* Admin */
	// login
	beego.Router("/admin/login", &admin.AdminController{}, "*:Login")
	beego.Router("/admin/ajaxLogin", &admin.AdminController{}, "post:AjaxLogin")

	// register
	beego.Router("/admin/register", &admin.AdminController{}, "*:Register")
	beego.Router("/admin/ajaxRegister", &admin.AdminController{}, "post:AjaxRegister")

	// index
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
}
