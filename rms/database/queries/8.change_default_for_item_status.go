package queries

var Query8_Change_Default_For_Item_Status = `
	ALTER TABLE menu_item ALTER COLUMN available SET DEFAULT TRUE;
`

// THIS MIGRATION IS APPLIED, DO NOT MODIFY THIS FILE
