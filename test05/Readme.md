# Go-Mysql Driver

https://github.com/go-sql-driver/mysql

Create Database:
CREATE DATABASE test_go_db;
```
CREATE TABLE `test_go_db`.`users` ( 
    `id` INT(5) UNSIGNED NOT NULL AUTO_INCREMENT , 
    `name` VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL , `age` INT(5) UNSIGNED NOT NULL ,
     PRIMARY KEY (`id`)) ENGINE = MyISAM;
```