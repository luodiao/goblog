package main

import (
	_ "goblog/routers"
	"github.com/astaxie/beego"
    "goblog/models"
)
// 
func init() {
   models.CreateTable()
}
func main() {
	beego.Run()
}
