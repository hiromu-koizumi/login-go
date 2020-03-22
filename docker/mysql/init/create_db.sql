SET CHARSET UTF8;
DROP DATABASE IF EXISTS sample_docker_compose;
CREATE DATABASE IF NOT EXISTS sample_docker_compose DEFAULT CHARACTER SET utf8;
use sample_docker_compose;
CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL
)DEFAULT CHARACTER SET=utf8;
 
INSERT INTO users (name) VALUES ("太郎"),("花子"),("令和");