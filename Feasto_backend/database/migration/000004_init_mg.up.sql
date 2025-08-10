CREATE TABLE IF NOT EXISTS OrderItems(
    id integer PRIMARY KEY AUTO_INCREMENT,
    order_id integer,
    customer_id integer,
    product_id integer,
    quantity integer DEFAULT 1,
    FOREIGN KEY (order_id) REFERENCES Orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES Products(id) ON DELETE CASCADE   
); 