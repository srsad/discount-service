CREATE TABLE `users` (
	`id` INT(20) AUTO_INCREMENT,
	`first_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
	`last_name` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '',
	`active` BOOLEAN(20),
	`created_at` TIMESTAMP(20),
	`updated_at` TIMESTAMP(20) ON UPDATE CURRENT_TIMESTAMP,

	PRIMARY KEY (`id`)
);
