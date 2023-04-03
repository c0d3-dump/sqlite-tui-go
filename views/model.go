package views

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	cursor       []int
	textInput    textinput.Model
	tables       []Table
	tableData    TableData
	createTable  CreateTable
	createColumn CreateColumn
	queryText    string
}

type Table struct {
	cursor  int
	columns []string
	ids     []string
	data    TableData
}

type TableData struct {
	cursor int
	data   map[string]string
}

func NewModel() Model {
	ti := textinput.New()
	ti.Focus()
	ti.Cursor.SetMode(cursor.CursorStatic)

	return Model{
		cursor:    []int{0, 0, 0},
		textInput: ti,
		tables:    []Table{},
		tableData: TableData{},
		createTable: CreateTable{
			cursor: 0,
			name:   "user",
			columns: []CreateColumn{
				{name: "id", dtype: "INT", pk: true},
				{name: "name", dtype: "TEXT"},
			},
		},
		createColumn: CreateColumn{},
	}
}

type CreateTable struct {
	cursor  int
	name    string
	columns []CreateColumn
}

type CreateColumn struct {
	cursor  int
	name    string
	dtype   string
	notnull bool
	dval    string
	pk      bool
}
