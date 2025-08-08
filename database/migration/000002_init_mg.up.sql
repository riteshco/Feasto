CREATE TABLE IF NOT EXISTS Products ( 
    id integer PRIMARY KEY AUTO_INCREMENT,
    product_name varchar(100) NOT NULL UNIQUE, 
    isavailable boolean DEFAULT true,
    price decimal(10, 2) NOT NULL,
    category varchar(100) DEFAULT NULL,
    image_url varchar(255) DEFAULT NULL
);


-- demo products
INSERT INTO Products (product_name , price , category , image_url) VALUES ('Pizza' , 399 , 'Fast Food' , '/demo_products/pizza.jpg');
INSERT INTO Products (product_name , price , category , image_url) VALUES ('Burger' , 129 , 'Fast Food' , '/demo_products/burger.jpg');