DROP TABLE IF EXISTS `admins_logs`;
CREATE TABLE `admins_logs` (

  `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '管理员日志ID',
  `admin_id`   int(11) unsigned    NOT NULL                COMMENT '管理员ID',
  `request_ip` varchar(50)         NOT NULL                COMMENT 'IP地址',
  `content`    varchar(255)        NOT NULL                COMMENT '日志内容',
  `created_at` datetime            NOT NULL                COMMENT '创建时间',

  PRIMARY KEY (`id`),
  KEY `admin_id` (`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员操作日志表';