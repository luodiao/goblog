package routers

import (
	"goblog/controllers"
	"goblog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	// beego.Router("/admin", &admin.AdminController{})
	// beego.AutoRouter(&admin.AdminController{})

	// Admin
	beego.Router("/admin/login", &admin.AdminController{}, "*:Login")

	beego.Router("/admin/register", &admin.AdminController{}, "*:Register")
	beego.Router("/admin/ajaxRegister", &admin.AdminController{}, "post:AjaxRegister")
}
