package queries

var Query7_Toggle_MenuItem_Status = `
	ALTER TABLE menu_item ADD COLUMN available BOOLEAN DEFAULT FALSE;
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
