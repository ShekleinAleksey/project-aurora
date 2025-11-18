CREATE TABLE IF NOT EXIST expenses (
    id SERIAL PRIMARY KEY,
    amount INT,     
	category VARCHAR(255) NOT NULL,   
	description VARCHAR(255),
	date TIMESTAMP NOT NULL,        
	createdAt TIMESTAMP NOT NULL   
);