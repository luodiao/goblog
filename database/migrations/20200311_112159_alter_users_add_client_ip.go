package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterUsersAddClientIp_20200311_112159 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterUsersAddClientIp_20200311_112159{}
	m.Created = "20200311_112159"

	migration.Register("AlterUsersAddClientIp_20200311_112159", m)
}

// Run the migrations
func (m *AlterUsersAddClientIp_20200311_112159) Up() {
	m.SQL("ALTER TABLE `users` add `client_ip` varchar(20) default '';")

}

// Reverse the migrations
func (m *AlterUsersAddClientIp_20200311_112159) Down() {
	m.SQL("ALTER TABLE `users` DROP `client_ip`;")
}
