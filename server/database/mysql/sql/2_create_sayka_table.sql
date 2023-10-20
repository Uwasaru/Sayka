CREATE TABLE IF NOT EXISTS `saykas` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `title`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `user_id`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `github_url`  varchar(255) COLLATE utf8mb4_bin,
  `app_url`     varchar(255) COLLATE utf8mb4_bin,
  `slide_url`   varchar(255) COLLATE utf8mb4_bin,
  `article_url` varchar(255) COLLATE utf8mb4_bin,
  `figma_url`   varchar(255) COLLATE utf8mb4_bin,
  `description` text          COLLATE utf8mb4_bin NOT NULL,
  `created_at`  timestamp    COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  timestamp    COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX `user_id_index` (`user_id`),
  INDEX `id_index` (`id`),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;