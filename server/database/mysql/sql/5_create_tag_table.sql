CREATE TABLE `tags` (
  `id`       int(11) NOT NULL AUTO_INCREMENT,
  `post_id`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;