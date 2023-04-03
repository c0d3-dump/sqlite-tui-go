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
	name    string
	columns []string
	ids     []int
	data    []TableData
}

type TableData struct {
	data map[string]string
}

func NewModel() Model {
	ti := textinput.New()
	ti.Focus()
	ti.Cursor.SetMode(cursor.CursorStatic)

	return Model{
		cursor:    []int{0, 0, 0},
		textInput: ti,
		tables: []Table{
			{
				cursor:  0,
				name:    "user",
				columns: []string{"name", "email", "password"},
				ids:     []int{0, 1},
				data: []TableData{
					{
						data: map[string]string{
							"name":     "b2b",
							"email":    "b2b@gmail.com",
							"password": "1234",
						},
					},
					{
						data: map[string]string{
							"name":     "test",
							"email":    "test@gmail.com",
							"password": "1234",
						},
					},
				},
			},
		},
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
