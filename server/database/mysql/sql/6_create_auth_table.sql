CREATE TABLE IF NOT EXISTS `github_auths` (
  `user_id` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザID',
  `access_token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'OAuth2のアクセストークン',
  `refresh_token` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'OAuth2のリフレッシュトークン',
  `expiry` datetime NOT NULL COMMENT 'アクセストークンの有効期限',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;