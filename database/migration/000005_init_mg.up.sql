CREATE TABLE IF NOT EXISTS Payments (
    id integer PRIMARY KEY AUTO_INCREMENT,  
    user_id integer,        
    order_id integer,
    Total_amount decimal(10, 2) NOT NULL,
    payment_status enum('pending', 'completed', 'failed') DEFAULT 'pending',
    FOREIGN KEY (user_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES Orders(id) ON DELETE CASCADE
);