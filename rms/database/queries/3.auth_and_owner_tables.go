package queries

var Query3_Auth_And_Owner_Tables = `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		phone TEXT NOT NULL UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE auth (
		id SERIAL PRIMARY KEY,
		password TEXT,
		is_active BOOLEAN DEFAULT false,
		temp_password TEXT,
		temp_password_expiry TEXT,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		user_id INT REFERENCES users(id)
	);

	ALTER TABLE restaurant ADD COLUMN owner_id INT REFERENCES users(id) DEFAULT NULL;
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
