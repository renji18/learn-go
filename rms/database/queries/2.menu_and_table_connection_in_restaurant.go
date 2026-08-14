package queries

var Query2_Menu_And_Table_Connection_In_Restaurant = `
	CREATE TABLE menu_item(
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		category TEXT NOT NULL,
		price NUMERIC(10, 2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		restaurant_id INT REFERENCES restaurant(id)
	);

	CREATE TABLE r_table(
		id SERIAL PRIMARY KEY,
		number INT,
		chair_count INT,
		occupied BOOLEAN DEFAULT false,
		reserved BOOLEAN DEFAULT false,
		restaurant_id INT REFERENCES restaurant(id)
	);
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
