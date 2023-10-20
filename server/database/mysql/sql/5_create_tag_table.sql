CREATE TABLE IF NOT EXISTS `tags` (
  `id`       int(11) NOT NULL AUTO_INCREMENT,
  `sayka_id`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  INDEX `sayka_id_index` (`sayka_id`),
  PRIMARY KEY (id),
  FOREIGN KEY (sayka_id) REFERENCES saykas(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;