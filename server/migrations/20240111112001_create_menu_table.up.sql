CREATE TABLE `t_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '路由地址',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '路由名称(唯一)',
  `label` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `redirect` varchar(255) NOT NULL DEFAULT '' COMMENT '重定向地址',
  `components` varchar(255) NOT NULL DEFAULT '' COMMENT '页面组件',
  `order` tinyint(4) NOT NULL DEFAULT '0' COMMENT '排序',
  `meta` json DEFAULT null COMMENT 'meta信息',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id),
  KEY idx_menu_created_at (created_at),
  KEY idx_menu_updated_at (updated_at)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单';