CREATE TABLE IF NOT EXISTS `posts` (
	`id` varchar(26) NOT NULL,
	`user_id` varchar(26) NOT NULL,
	`title` varchar(50) NOT NULL,
	`content` text NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	PRIMARY KEY (`id`),
	FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
