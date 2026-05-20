-- +migrate Up
CREATE TABLE users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(64) NOT NULL,
    phone_number VARCHAR(32) NOT NULL UNIQUE,
    created_at TIMESTAMP 
);

-- +migrate Down
DROP TABLE users;
