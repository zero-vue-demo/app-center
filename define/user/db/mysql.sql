DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (

  `id`             int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name`           varchar(50)      NOT NULL                COMMENT '名称',
  `account`        varchar(50)      NOT NULL                COMMENT '账号',
  `password`       varchar(255)     NOT NULL                COMMENT '密码',
  `created_at`     datetime         NOT NULL                COMMENT '创建时间',
  `last_signin_at` datetime         DEFAULT NULL            COMMENT '上次登录时间',
  `updated_at`     datetime         DEFAULT NULL            COMMENT '更新时间',
  `deleted_at`     datetime         DEFAULT NULL            COMMENT '删除时间',

  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账户信息表';DROP TABLE IF EXISTS `users_logs`;
CREATE TABLE `users_logs` (

  `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户日志ID',
  `user_id`    int(11) unsigned    NOT NULL                COMMENT '用户ID',
  `request_ip` varchar(50)         NOT NULL                COMMENT 'IP 地址',
  `content`    varchar(255)        NOT NULL                COMMENT '日志内容',
  `created_at` datetime            NOT NULL                COMMENT '创建时间',

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户操作日志表';DROP TABLE IF EXISTS `users_tree_structure`;
CREATE TABLE `users_tree_structure` (

  `account_id` int(11) unsigned NOT NULL     COMMENT '账号ID',
  `tree_left`  int(11)          DEFAULT NULL COMMENT '左值',
  `tree_right` int(11)          DEFAULT NULL COMMENT '右值',
  `tree_level` int(11)          DEFAULT NULL COMMENT '层级',
  `tree_group` int(11)          DEFAULT NULL COMMENT '分组',

  `name`           varchar(50)      NOT NULL                COMMENT '名称',
  `account`        varchar(50)      NOT NULL                COMMENT '账号',

  PRIMARY KEY (`account_id`),
  KEY `tree_left` (`tree_left`),
  KEY `tree_right` (`tree_right`),
  KEY `tree_group` (`tree_group`),
  KEY `tree_level` (`tree_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户树形结构表';