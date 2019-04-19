SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `link`;
CREATE TABLE `link` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `long_url` text NOT NULL,
  `short_url` varchar(100) NULL DEFAULT '',
  `host` varchar(145) null DEFAULT 'localhost',
  `custom_hash` varchar(145) null DEFAULT '',
  `clicked` bigint(20) NULL DEFAULT 0,
  `last_clicked_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
