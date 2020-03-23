package main

import (
	_ "goblog/routers"

	"github.com/astaxie/beego"
)

//
func init() {
	//    models.CreateTable()

	// 设置session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
}
func main() {
	beego.Run()
}
