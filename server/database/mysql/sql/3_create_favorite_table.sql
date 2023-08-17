CREATE TABLE `favorites` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `sayka_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (sayka_id) REFERENCES saykas(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;