CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    Name VARCHAR(255) NOT NULL, 
    Quantity int  NOT NULL, 
    Unit VARCHAR(15)
);

 