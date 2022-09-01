CREATE TABLE `go_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(128) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `tel` char(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '账号状态',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '用户组名',
  `remark` text COMMENT '备注',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

CREATE TABLE `go_rule` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL COMMENT '路径',
  `code` varchar(255) DEFAULT NULL COMMENT '权限标示码',
  `is_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否显示 0否 1是',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级id',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1菜单  0按钮',
  `sort` smallint(6) NOT NULL COMMENT '排序',
  `icon` varchar(255) DEFAULT NULL,
  `cache` tinyint(1) DEFAULT '1',
  `title` varchar(255) DEFAULT NULL,
  `permission` varchar(255) DEFAULT NULL COMMENT '权限别名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=225 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

CREATE TABLE `go_role_rule` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `rule_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `role_rule_role_id_rule_id_index` (`role_id`,`rule_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=513 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

CREATE TABLE `go_admin_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `admin_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `admin_role_admin_id_role_id_index` (`admin_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=132 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

