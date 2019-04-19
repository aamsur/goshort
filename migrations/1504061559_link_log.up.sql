SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `link_log`;
CREATE TABLE `link_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `short_url` varchar(100) NULL DEFAULT '',
  `clicked_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
