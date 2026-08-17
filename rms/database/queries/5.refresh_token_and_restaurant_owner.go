package queries

var Query5_Refresh_Token_And_Restaurant_Owner = `
	ALTER TABLE auth ADD COLUMN refresh_token TEXT;

	ALTER TABLE users ADD COLUMN restaurant_id INT REFERENCES restaurant(id) DEFAULT NULL;
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
