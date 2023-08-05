CREATE TABLE `comments` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `post_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `content`     text COLLATE utf8mb4_bin NOT NULL,
  `type`        int(3) NOT NULL DEFAULT '0',
  `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;