package main

import (
	"goblog/controllers"
	_ "goblog/routers"

	"github.com/astaxie/beego"
)

//
func init() {
	//    models.CreateTable()

	// 设置session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600

	// 模板方法
	beego.AddFuncMap("getCategory", controllers.GetCategory)
}
func main() {
	beego.Run()
}
