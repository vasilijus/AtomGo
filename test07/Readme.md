# Go-Mysql Driver

https://github.com/go-sql-driver/mysql


```
CREATE TABLE `test_go_db`.`articles` ( 
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
     `title` VARCHAR(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL , `description` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT, NULL , 
     `text` TEXT CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL , 
     PRIMARY KEY (`id`)) ENGINE = MyIsam;
```