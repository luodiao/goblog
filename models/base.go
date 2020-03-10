package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    "time"
)

type Users struct {
    Id              int64
    Username        string `orm:"size(30)"`
    password        string `orm:"size(32)"`
    nickname        string `orm:"size(20)"`
    Created         time.Time `orm:"auto_now_add;type(datetime)"`
    Updated         time.Time `orm:"auto_now;type(datetime)"`
    Headimgurl      string `orm:"size(200)"`
    Client_ip       string `orm:"size(20)"`
}



func RegisterDB() {
    //注册 model
    orm.RegisterModel(new(Users))
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

}