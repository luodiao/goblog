package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableCategory_20200317_162457 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableCategory_20200317_162457{}
	m.Created = "20200317_162457"

	migration.Register("CreateTableCategory_20200317_162457", m)
}

// Run the migrations
func (m *CreateTableCategory_20200317_162457) Up() {
	m.SQL("CREATE TABLE `category` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父类ID',`category` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '类别名称',`weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重 从大到小',`alias` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '别名',`color` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '颜色',`pic` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '类别图片',`remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',`enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT '0:不启动 1:启用',`created_at` timestamp NULL DEFAULT NULL,`updated_at` timestamp NULL DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
}

// Reverse the migrations
func (m *CreateTableCategory_20200317_162457) Down() {
	m.SQL("DROP TABLE `category`")
}
