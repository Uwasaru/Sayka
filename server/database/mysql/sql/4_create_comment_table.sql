CREATE TABLE `comments` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `sayka_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `content`     text COLLATE utf8mb4_bin NOT NULL,
  `type`        int(3) NOT NULL DEFAULT '0',
  `created_at`  timestamp    COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (sayka_id) REFERENCES saykas(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;