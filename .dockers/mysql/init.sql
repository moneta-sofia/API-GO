SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN = 0;

SET @@GLOBAL_GTID_PURGED=/*!80000 '+' */ '';

CREATE DATABASE IF NOT EXISTS `go_course_users`;

CREATE TABLE `go_course_users`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(50) NULL,
    `last_name` VARCHAR(50) NULL,
    `email` VARCHAR(100) NULL,
    PRIMARY KEY (`id`)
);
