CREATE DATABASE IF NOT EXISTS Todo;
USE Todo;

DROP TABLE IF EXISTS info;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE info(
    id int auto_increment primary key,
    title varchar(50) not null,
    description varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE = INNODB;

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