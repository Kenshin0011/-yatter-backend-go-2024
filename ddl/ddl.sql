CREATE TABLE `account` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `status` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account_id` int NOT NULL,
  `url` varchar(255),
  `content` text NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `account`(`id`)
);
