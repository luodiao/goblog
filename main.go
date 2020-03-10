package main

import (
    "goblog/models"
	_ "goblog/routers"
	"github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _"github.com/go-sql-driver/mysql"
)
func init() {
    // 注册数据库
    models.RegisterDB()
}

func main() {
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)
    //运行
	beego.Run()
}

