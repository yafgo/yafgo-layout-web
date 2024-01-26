CREATE TABLE `t_dms_data` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `db_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据库id',
  `table_name` varchar(50) NOT NULL DEFAULT '' COMMENT '真实表名',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '显示表名',
  `comment` varchar(255) NOT NULL DEFAULT '' COMMENT '表注释',
  `filter` varchar(500) NOT NULL DEFAULT '' COMMENT '自定义过滤条件',
  `sql` varchar(2000) NOT NULL DEFAULT '' COMMENT '自定义SQL',
  `sort` varchar(200) NOT NULL DEFAULT '' COMMENT '自定义排序',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_data_id (db_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'dms数据表管理';