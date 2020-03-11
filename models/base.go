package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql"
)



//连接mysql
func init() {
    //注册驱动
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //注册默认数据库
    dbhost := beego.AppConfig.String("dbhost")
    dbport := beego.AppConfig.String("dbport")
    dbuser := beego.AppConfig.String("dbuser")
    dbpassword := beego.AppConfig.String("dbpassword")
    dbname :=beego.AppConfig.String("dbname")
    dsn := dbuser + ":" +dbpassword +"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&loc=Asia%2FShanghai"
    orm.RegisterDataBase("default", "mysql", dsn)
    orm.RegisterModel(new(Users))
}

//注册 model
func CreateTable() {
    //注册驱动
    orm.Debug = true
    // 自动建表
    orm.RunSyncdb("default", false, true)
}