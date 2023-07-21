DROP TABLE IF EXISTS `users_logs`;
CREATE TABLE `users_logs` (

  `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户日志ID',
  `user_id`    int(11) unsigned    NOT NULL                COMMENT '用户ID',
  `request_ip` varchar(50)         NOT NULL                COMMENT 'IP 地址',
  `content`    varchar(255)        NOT NULL                COMMENT '日志内容',
  `created_at` datetime            NOT NULL                COMMENT '创建时间',

  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户操作日志表';