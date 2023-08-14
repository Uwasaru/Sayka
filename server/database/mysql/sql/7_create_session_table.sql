CREATE TABLE `login_sessions` (
  `id` VARCHAR(255) NOT NULL,
  `user_id` int(8) NOT NULL,
  `expiry` datetime NOT NULL COMMENT 'sessionの有効期限',
  PRIMARY KEY (`id`)
)ENGINE = InnoDB;