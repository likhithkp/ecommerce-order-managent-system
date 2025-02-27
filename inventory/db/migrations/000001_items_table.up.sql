CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,                 
    description TEXT NOT NULL,                     
    price DECIMAL(10,2) NOT NULL,      
    count INT NOT NULL CHECK (count >= 0),
    category TEXT NOT NULL,                          
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
