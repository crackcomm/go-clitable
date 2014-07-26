package clitable

import "strings"

var mapFields = []string{"Key", "Value"}

// PrintHorizontal - Prints horizontal table from a map.
func PrintHorizontal(m map[string]interface{}) {
	// Create new table with
	table := New(mapFields)
	// Convert map to rows list
	rows := mapToRows(m)
	// Add rows to table
	for _, row := range rows {
		table.AddRow(row)
	}
	// Do not print head
	table.HideHead = true
	// Print table
	table.Print()
}

func mapToRows(m map[string]interface{}) (rows []map[string]interface{}) {
	rows = []map[string]interface{}{}
	for key, value := range m {
		row := map[string]interface{}{}
		row["Key"] = strings.Title(key)
		row["Value"] = value
		rows = append(rows, row)
	}
	return
}
