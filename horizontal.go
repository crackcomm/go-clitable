package clitable

import "strings"

var mapFields = []string{"Key", "Value"}

// PrintHorizontal - Prints horizontal table from a map.
func PrintHorizontal(m map[string]interface{}) {
	table := New(mapFields)
	rows := mapToRows(m)
	for _, row := range rows {
		table.AddRow(row)
	}
	table.HideHead = true
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
