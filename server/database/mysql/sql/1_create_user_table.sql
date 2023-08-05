CREATE TABLE `users` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `name`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `email`       varchar(255) UNIQUE COLLATE utf8mb4_bin NOT NULL,
  `password`    varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `github_url`  varchar(255) COLLATE utf8mb4_bin,
  `created_at`  timestamp    COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;