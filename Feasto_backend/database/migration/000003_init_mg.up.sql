CREATE TABLE IF NOT EXISTS Orders (
    id integer PRIMARY KEY AUTO_INCREMENT,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    current_status enum('pending', 'accepted', 'rejected', 'delivered') DEFAULT 'pending',
    customer_id integer,
    chef_id integer DEFAULT NULL,
    table_number integer NOT NULL,
    instructions text DEFAULT NULL,
    FOREIGN KEY (customer_id) REFERENCES Users(id) ON DELETE CASCADE,
    FOREIGN KEY (chef_id) REFERENCES Users(id) ON DELETE SET NULL
); 