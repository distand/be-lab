CREATE TABLE `lab_booking` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL,
  `device_id` int(10) unsigned NOT NULL,
  `memo` varchar(1024) NOT NULL DEFAULT '',
  `status` tinyint(3) NOT NULL DEFAULT '1',
  `is_del` tinyint(3) NOT NULL DEFAULT '0',
  `stime` datetime NOT NULL,
  `etime` datetime NOT NULL,
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `device_id` (`device_id`),
  KEY `stime` (`stime`),
  KEY `etime` (`etime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='预约表';