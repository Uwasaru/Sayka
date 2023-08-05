CREATE TABLE `session` (
  `id` varchar(255) NOT NULL,
  `session_id` varchar(255) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `token_expire` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;