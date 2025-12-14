CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP NOT NULL   
);

CREATE TABLE IF NOT EXISTS purchases (
    id INT PRIMARY KEY,
    material_id INT NOT NULL,
    count DECIMAL(10, 2) NOT NULL,
    unit_price DECIMAL(10, 2) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    notes TEXT,
    purchase_date DATETIME NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
);

