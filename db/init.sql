CREATE DATABASE IF NOT EXISTS http_server
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

CREATE USER 'lssop'@'localhost' IDENTIFIED BY 'redacted';
GRANT ALL PRIVILEGES ON database_name.* TO 'lssop'@'localhost';
FLUSH PRIVILEGES;

SHOW DATABASES LIKE 'http_server';

USE http_server;

CREATE TABLE users (
    id INT AUTO_INCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME,
    PRIMARY KEY (id)
);

INSERT INTO users (username, password, created_at) VALUES ('lssop', 'redacted', NOW());