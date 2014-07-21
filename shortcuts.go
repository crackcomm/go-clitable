package clitable

// PrintRow - Prints table with only one row.
func PrintRow(fields []string, row map[string]interface{}) {
	table := New(fields)
	// add row
	table.AddRow(row)
	// And display table
	table.Print()
}

// PrintTable - Prints table.
func PrintTable(fields []string, rows []map[string]interface{}) {
	// Construct a table
	table := New(fields)
	for _, r := range rows {
		table.AddRow(r)
	}

	// And display table
	table.Print()
}
