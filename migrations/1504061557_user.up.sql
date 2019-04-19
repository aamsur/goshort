SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(145) NOT NULL,
  `last_login` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `last_logout_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `user` (`id`, `full_name`, `email`, `address`, `username`, `password`, `last_login`, `created_at`, `updated_at`, `last_logout_at`) VALUES
(1,	'admin',	'',	'',	'admin',	'$2y$12$9xF9.RZ.gHLaCDWSAnremOQXqOeAU14eq5iw8jrouOqQYmiFVfDCW',	'2019-03-15 16:17:24',	'2017-08-21 10:20:07',	NULL,	NULL);