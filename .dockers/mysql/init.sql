SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN = 0;

CREATE DATABASE IF NOT EXISTS `go-course-users`;

CREATE TABLE `go-course-users`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `first_name` VARCHAR(50) NULL,
    `last_name` VARCHAR(50) NULL,
    `email` VARCHAR(100) NULL,
    PRIMARY KEY (`id`)
);
