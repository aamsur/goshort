SET FOREIGN_KEY_CHECKS = 0;

ALTER TABLE `link`
ADD UNIQUE `link_unique_short_url_host` (`short_url`, `host`),
ADD INDEX `link_index_short_url` (`short_url`),
ADD INDEX `link_index_host` (`host`),
ADD INDEX `link_index_id` (`id`);

ALTER TABLE `link_log`
ADD INDEX `id` (`id`),
ADD INDEX `link_log_index_short_url` (`short_url`);
