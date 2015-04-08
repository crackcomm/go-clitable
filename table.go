// Package clitable implements methods for pretty command line table output.
package clitable

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"strings"
)

// Table - Table structure.
type Table struct {
	Fields     []string
	Rows       []map[string]string
	HideHead   bool // when true doesn't print header
	Markdown   bool
	Footer     map[string]string
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

func (t *Table) calculateSizes(row map[string]string) {
	for _, k := range t.Fields {
		v, ok := row[k]
		if !ok {
			continue
		}

		vlen := runewidth.StringWidth(v)
		// align to field name length
		if klen := runewidth.StringWidth(k); vlen < klen {
			vlen = klen
		}
		vlen += 2 // + 2 spaces
		if t.fieldSizes[k] < vlen {
			t.fieldSizes[k] = vlen
		}
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

		newRow[k] = val
	}

	t.calculateSizes(newRow)

	if len(newRow) > 0 {
		t.Rows = append(t.Rows, newRow)
	}
}

// Print - Prints table.
func (t *Table) Print() {
	if len(t.Rows) == 0 && t.Footer == nil {
		return
	}

	t.calculateSizes(t.Footer)

	if !t.Markdown {
		t.printDash()
	}

	if !t.HideHead {
		fmt.Println(t.getHead())
		if t.Markdown {
			t.printMarkdownDash()
		} else {
			t.printDash()
		}
	}

	for _, r := range t.Rows {
		fmt.Println(t.rowString(r))
		if !t.Markdown {
			t.printDash()
		}
	}

	if t.Footer != nil {
		if t.Markdown {
			t.printMarkdownDash()
		} else {
			t.printDash()
		}
		fmt.Println(t.rowString(t.Footer))
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
	spacesLeft := t.fieldSizes[name] - runewidth.StringWidth(value)
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

// printMarkdownDash - Prints dash in middle of table.
func (t *Table) printMarkdownDash() {
	r := make(map[string]string)
	for _, name := range t.Fields {
		r[name] = strings.Repeat("-", t.fieldSizes[name]-2)
	}
	fmt.Println(t.rowString(r))
}

// lineLength - Counts size of table line length (with spaces etc.).
func (t *Table) lineLength() (sum int) {
	sum = 1
	for _, l := range t.fieldSizes {
		sum += l + 1
	}
	return sum
}
