CREATE DATABASE IF NOT EXISTS seller;

USE seller;

DROP TABLE IF EXISTS sales;
DROP TABLE IF EXISTS interactions;
DROP TABLE IF EXISTS client;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
	id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    email_backup VARCHAR(50) NOT NULL,
    password VARCHAR(100)NOT NULL,
    active BOOL NOT NULL DEFAULT true
) ENGINE=INNODB;

CREATE TABLE clients(
    id INT AUTO_INCREMENT PRIMARY KEY,

    seller_id INT NOT NULL,
    FOREIGN KEY (seller_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    name VARCHAR(50) NOT NULL,
    contacts TEXT,
    address VARCHAR(50),
    active BOOL NOT NULL DEFAULT true
) ENGINE=INNODB;

CREATE TABLE interactions(
    id INT AUTO_INCREMENT PRIMARY KEY,

    seller_id INT NOT NULL,
    FOREIGN KEY (seller_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    client_id INT NOT NULL,
    FOREIGN KEY (client_id)
    REFERENCES clients(id)
    ON DELETE CASCADE,

    status VARCHAR(50),
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    interaction VARCHAR(50),
    content TEXT,
    active BOOL NOT NULL DEFAULT true
) ENGINE=INNODB;

CREATE TABLE sales(
    id INT AUTO_INCREMENT PRIMARY KEY,

    seller_id INT NOT NULL,
    FOREIGN KEY (seller_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    client_id INT NOT NULL,
    FOREIGN KEY (client_id)
    REFERENCES client(id)
    ON DELETE CASCADE,

    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sale TEXT
) ENGINE=INNODB;