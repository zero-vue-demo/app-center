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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户账户信息表';