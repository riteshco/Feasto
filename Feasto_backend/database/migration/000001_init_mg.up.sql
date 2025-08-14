USE test_db;
CREATE TABLE IF NOT EXISTS Users (
    id integer PRIMARY KEY AUTO_INCREMENT,
    username varchar(100) NOT NULL,
    mobile_number bigint NOT NULL UNIQUE,
    email varchar(100) NOT NULL UNIQUE,
    user_role enum('admin' , 'customer' , 'chef') NOT NULL,
    password_hash varchar(255) NOT NULL
    change_role_to ENUM('customer' , 'admin' , 'chef') DEFAULT NULL
);  

-- admin account
-- password is "secret" for admin.
INSERT INTO Users (username , mobile_number , email , user_role , password_hash) VALUES ('admin' , '1234512345' , 'admin@gmail.com' , 'admin' , '$2a$10$vzca.bQW09GcMu5UFVErCe6qmmt0vsLonHnen6NUJ7oQfXEM2pv0K');

-- chef account
-- password is "secret" for chef as well.
INSERT INTO Users (username , mobile_number , email , user_role , password_hash) VALUES ('chef' , '1234512344' , 'chef@gmail.com' , 'chef' , '$2a$10$vzca.bQW09GcMu5UFVErCe6qmmt0vsLonHnen6NUJ7oQfXEM2pv0K');