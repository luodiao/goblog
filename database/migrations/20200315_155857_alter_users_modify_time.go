package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterUsersModifyTime_20200315_155857 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterUsersModifyTime_20200315_155857{}
	m.Created = "20200315_155857"

	migration.Register("AlterUsersModifyTime_20200315_155857", m)
}

// Run the migrations
func (m *AlterUsersModifyTime_20200315_155857) Up() {
	m.SQL("ALTER TABLE `users` MODIFY `created_at` int(11);")
	m.SQL("ALTER TABLE `users` MODIFY `updated_at` int(11);")
}

// Reverse the migrations
func (m *AlterUsersModifyTime_20200315_155857) Down() {
	m.SQL("ALTER TABLE `users` MODIFY `created_at` timestamp;")
	m.SQL("ALTER TABLE `users` MODIFY `updated_at` timestamp;")
}
