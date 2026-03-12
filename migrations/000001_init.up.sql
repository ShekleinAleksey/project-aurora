CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP NOT NULL  
);

CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    --material_id INT NOT NULL,
    material VARCHAR(255),
    count DECIMAL(10, 2) NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    notes TEXT,
    purchase_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

