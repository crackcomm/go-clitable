package clitable

// PrintRow - Prints table with only one row.
func PrintRow(fields []string, row map[string]interface{}) {
	table := New(fields)
	table.AddRow(row)
	table.Print()
}

// PrintTable - Prints table.
func PrintTable(fields []string, rows []map[string]interface{}) {
	table := New(fields)
	for _, r := range rows {
		table.AddRow(r)
	}
	table.Print()
}
