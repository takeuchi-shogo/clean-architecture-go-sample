CREATE TABLEIF NOT EXISTS `users` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`display_name` varchar(50) NOT NULL,
	`screen_name` varchar(50) NOT NULL,
	`email` varchar(30),

	PRIMARY KEY (`id`),
)
