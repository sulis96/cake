CREATE TABLE `cake`(
    `id` INT NOT NULL AUTO_INCREMENT , 
    `title` CHAR(50) NOT NULL , 
    `description` VARCHAR(255) NOT NULL, 
    `rating` FLOAT, 
    `image` VARCHAR(255), 
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL , 
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP , 
    PRIMARY KEY (`id`));