USE feasto;
CREATE TABLE IF NOT EXISTS Users (
    id integer PRIMARY KEY AUTO_INCREMENT,
    username varchar(100) NOT NULL,
    mobile_number bigint NOT NULL UNIQUE,
    email varchar(100) NOT NULL UNIQUE,
    user_role enum('admin' , 'customer' , 'chef') NOT NULL,
    password_hash varchar(255) NOT NULL
);  