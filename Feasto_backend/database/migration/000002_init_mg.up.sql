CREATE TABLE IF NOT EXISTS Products ( 
    id integer PRIMARY KEY AUTO_INCREMENT,
    product_name varchar(100) NOT NULL UNIQUE, 
    isavailable boolean DEFAULT true,
    price decimal(10, 2) NOT NULL,
    category varchar(100) DEFAULT NULL,
    image_url varchar(255) DEFAULT NULL
);


-- demo products
INSERT INTO Products (product_name , price , category , image_url) VALUES ('Pizza' , 39 , 'Fast Food' , 'https://images8.alphacoders.com/369/369063.jpg');
INSERT INTO Products (product_name , price , category , image_url) VALUES ('Burger' , 29 , 'Fast Food' , 'https://images2.alphacoders.com/135/1353005.png');
INSERT INTO Products (product_name , price , category , image_url) VALUES ('Pasta' , 69.69 , 'Italian' , 'https://cdn.pixabay.com/photo/2022/10/12/22/09/spaghetti-bolognese-7517639_960_720.jpg');