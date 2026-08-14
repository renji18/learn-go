package queries

var Query6_Remove_Temp_Fields_From_Auth = `
	ALTER TABLE auth DROP COLUMN temp_password;
	
	ALTER TABLE auth DROP COLUMN temp_password_expiry;
`