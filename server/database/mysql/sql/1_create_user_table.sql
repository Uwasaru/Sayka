CREATE TABLE `users` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `github_id`        varchar(255) UNIQUE COLLATE utf8mb4_bin NOT NULL,
  `github_icon`       varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `access_token`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `reflesh_token`  varchar(255) COLLATE utf8mb4_bin,
  `token_expire`  timestamp    COLLATE utf8mb4_bin NOT NULL,
  `created_at`  timestamp    COLLATE utf8mb4_bin NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;