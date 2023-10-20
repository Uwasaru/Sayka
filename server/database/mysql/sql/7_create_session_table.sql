CREATE TABLE IF NOT EXISTS `login_sessions` (
  `id` VARCHAR(255) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `expiry` datetime NOT NULL COMMENT 'sessionの有効期限',
  PRIMARY KEY (`id`)
)ENGINE = InnoDB;