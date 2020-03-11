package main

import (
	_ "goblog/routers"

	"github.com/astaxie/beego"
)

//
func init() {
	//    models.CreateTable()
}
func main() {
	beego.Run()
}
