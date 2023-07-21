DROP TABLE IF EXISTS `users_tree_structure`;
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