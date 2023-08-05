CREATE TABLE `posts` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `title`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `github_url`  varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `app_url`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `slide_url`   varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `description` text          COLLATE utf8mb4_bin NOT NULL,
  `created_at`  TIMESTAMP    COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;