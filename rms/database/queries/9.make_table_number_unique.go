package queries

var Query9_Make_TableNumber_Unique = `
	ALTER TABLE r_table DROP COLUMN number;

	ALTER TABLE r_table ADD COLUMN table_number INT NOT NULL UNIQUE;
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
