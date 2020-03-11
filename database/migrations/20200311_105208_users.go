package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20200311_105208 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20200311_105208{}
	m.Created = "20200311_105208"

	migration.Register("Users_20200311_105208", m)
}

// Run the migrations
func (m *Users_20200311_105208) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `users` (`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT, `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL, `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL, `email_verified_at` timestamp NULL DEFAULT NULL, `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL, `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL, `created_at` timestamp NULL DEFAULT NULL, `updated_at` timestamp NULL DEFAULT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
}

// Reverse the migrations
func (m *Users_20200311_105208) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	// m.SQL("DROP TABLE `users`")
}
