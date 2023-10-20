CREATE TABLE IF NOT EXISTS `users` (
  `id`        varchar(255) COLLATE utf8mb4_bin NOT NULL ,
  `name`      varchar(255) COLLATE utf8mb4_bin,
  `img`       varchar(255) COLLATE utf8mb4_bin,
  PRIMARY KEY (`id`),
  INDEX `id_index` (`id`),
  INDEX `name_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;