SET CHARSET UTF8;
DROP DATABASE IF EXISTS sample_docker_compose;
CREATE DATABASE IF NOT EXISTS sample_docker_compose DEFAULT CHARACTER SET utf8;
use sample_docker_compose;
CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
name TEXT NOT NULL,
mail TEXT NOT NULL
)
DEFAULT CHARACTER SET=utf8;
 
INSERT INTO users (name,mail) VALUES ("太郎","aaa"),("花子","uuu"),("令和","sss");
