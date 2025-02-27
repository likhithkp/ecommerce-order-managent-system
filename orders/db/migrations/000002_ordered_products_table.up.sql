CREATE TABLE order_products (
    order_id VARCHAR(50),
    prod_id VARCHAR(50),
    quantity INT,
    PRIMARY KEY (order_id, prod_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);