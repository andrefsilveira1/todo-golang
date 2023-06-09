CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS info;



CREATE TABLE info(
    id int AUTO_INCREMENT primary key,
    title varchar(50) not null,
    description varchar(50) not null unique,
    completed boolean not null,
    createdAt timestamp default current_timestamp()
) ENGINE = INNODB;

INSERT INTO info (title, description, completed) values
("New post created!", "This is a new post created", false),
("Second post created!", "This is the second post", true),
("Third post created!", "This is the third post", false);

SELECT * FROM info;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id int AUTO_INCREMENT primary key,
	name varchar(50) not null,
	email varchar(50) not null unique,
	password varchar(100) not null,
	createdAt timestamp default current_timestamp()
) ENGINE = INNODB;

INSERT INTO users (name, email, password) values
("André Freitas", "freitasandre38@gmail.com", "1234456"),
("Emanuel Deodato", "emanoeldeodato@gmail.com", "abcbdd"),
("Lucas Vinicius", "lucasvini@gmail.com", "aakaksjdqi");

SELECT * FROM users;


-- CREATE TABLE seguidores(
--     usuario_id int not null, 
--     FOREIGN KEY (usuario_id) 
--     REFERENCES usuarios(id)
--     ON DELETE CASCADE,
--     seguidor_id int not null,
--     FOREIGN KEY (seguidor_id)
--     REFERENCES usuarios(id)
--     ON DELETE CASCADE,

--     primary key (usuario_id, seguidor_id)
-- )ENGINE = INNODB;