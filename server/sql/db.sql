CREATE DATABASE IF NOT EXISTS todo;
USE todo;

ALTER TABLE data
DROP FOREIGN KEY data_ibfk_1;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS data;



CREATE TABLE data (
    id int AUTO_INCREMENT primary key,
    user_id int not null,
    title varchar(50) not null,
    description varchar(50) not null unique,
    completed boolean not null,
    createdAt timestamp default current_timestamp(),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = INNODB;

CREATE TABLE users (
	id int AUTO_INCREMENT primary key,
	name varchar(50) not null,
	email varchar(50) not null unique,
	password varchar(100) not null,
	createdAt timestamp default current_timestamp()
) ENGINE = INNODB;