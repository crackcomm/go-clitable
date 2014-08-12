// Package clitable implements methods for pretty command line table output.
package clitable

import (
	"fmt"
	"strings"
)

// Table - Table structure.
type Table struct {
	Fields     []string
	Rows       []map[string]string
	HideHead   bool // when true doesn't print header
	fieldSizes map[string]int
}

// New - Creates a new table.
func New(fields []string) *Table {
	return &Table{
		Fields:     fields,
		Rows:       make([]map[string]string, 0),
		fieldSizes: make(map[string]int),
	}
}

// AddRow - Adds table row.
func (t *Table) AddRow(row map[string]interface{}) {
	newRow := make(map[string]string)
	for _, k := range t.Fields {
		v := row[k]
		// If is not nil format
		// else value is empty string
		var val string
		if v == nil {
			val = ""
		} else {
			val = fmt.Sprintf("%v", v)
		}

		valLen := len(val)
		// align to field name length
		if valLen < len(k) {
			valLen = len(k)
		}
		valLen += 2 // + 2 spaces
		if t.fieldSizes[k] < valLen {
			t.fieldSizes[k] = valLen
		}
		newRow[k] = val
	}
	if len(newRow) > 0 {
		t.Rows = append(t.Rows, newRow)
	}
}

// Print - Prints table.
func (t *Table) Print() {
	if len(t.Rows) == 0 {
		return
	}
	t.printDash()
	if !t.HideHead {
		fmt.Println(t.getHead())
		t.printDash()
	}
	for _, r := range t.Rows {
		fmt.Println(t.rowString(r))
		t.printDash()
	}
}

// getHead - Returns table header containing fields names.
func (t *Table) getHead() string {
	s := "|"
	for _, name := range t.Fields {
		s += t.fieldString(name, strings.Title(name)) + "|"
	}
	return s
}

// rowString - Creates a string row.
func (t *Table) rowString(row map[string]string) string {
	s := "|"
	for _, name := range t.Fields {
		value := row[name]
		s += t.fieldString(name, value) + "|"
	}
	return s
}

// fieldString - Creates field value string.
func (t *Table) fieldString(name, value string) string {
	value = fmt.Sprintf(" %s ", value)
	spacesLeft := t.fieldSizes[name] - len(value)
	if spacesLeft > 0 {
		for i := 0; i < spacesLeft; i++ {
			value += " "
		}
	}
	return value
}

// printDash - Prints dash (on top and header).
func (t *Table) printDash() {
	s := "|"
	for i := 0; i < t.lineLength()-2; i++ {
		s += "-"
	}
	s += "|"
	fmt.Println(s)
}

// lineLength - Counts size of table line length (with spaces etc.).
func (t *Table) lineLength() (sum int) {
	sum = 1
	for _, l := range t.fieldSizes {
		sum += l + 1
	}
	return sum
}
