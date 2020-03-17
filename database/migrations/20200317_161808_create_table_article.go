package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableArticle_20200317_161808 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableArticle_20200317_161808{}
	m.Created = "20200317_161808"

	migration.Register("CreateTableArticle_20200317_161808", m)
}

// Run the migrations
func (m *CreateTableArticle_20200317_161808) Up() {
	m.SQL("CREATE TABLE `article` (`id` int(10) unsigned NOT NULL AUTO_INCREMENT,`fk_category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别外键',`fk_users_id` int(11) NOT NULL COMMENT '用户外键',`title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',`weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重 从大到小',`cover` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '封面图片',`remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '简介',`content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',`tags` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签',`pv` int(11) NOT NULL COMMENT '访问量',`real_pv` int(11) NOT NULL COMMENT '真实访问量',`status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '0:下线 1:上线 2:待审核',`created_at` timestamp NULL DEFAULT NULL,`updated_at` timestamp NULL DEFAULT NULL,PRIMARY KEY (`id`),KEY `article_fk_category_id_status_index` (`fk_category_id`,`status`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
}

// Reverse the migrations
func (m *CreateTableArticle_20200317_161808) Down() {
	m.SQL("DROP TABLE `article`")
}
