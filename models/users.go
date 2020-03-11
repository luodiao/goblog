package models

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type Users struct {
    Id              int64
    Username        string `orm:"size(30)"`
    Password        string `orm:"size(32)"`
    Pickname        string `orm:"size(20)"`
    Created         time.Time `orm:"auto_now_add;type(datetime)"`
    Updated         time.Time `orm:"auto_now;type(datetime)"`
    Headimgurl      string `orm:"size(200)"`
    Client_ip       string `orm:"size(20)"`
}



//新增
func (m *Users) Insert() error {
    if _, err := orm.NewOrm().Insert(m); err != nil {
        return err
    }
    return nil
}

//执行sql
func (m *Users) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(m)
}