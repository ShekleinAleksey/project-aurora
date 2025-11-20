CREATE TABLE IF NOT EXIST categories (
    id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	createdAt TIMESTAMP NOT NULL   
);

CREATE TABLE IF NOT EXIST expenses (
    id SERIAL PRIMARY KEY,
    amount INT,     
	category VARCHAR(255) NOT NULL,   
	description VARCHAR(255),
	date TIMESTAMP NOT NULL,        
	createdAt TIMESTAMP NOT NULL   
);

