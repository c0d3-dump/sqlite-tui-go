package views

import (
	"sqlite-tui-go/database"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	cursor       []int
	d            database.Database
	textInput    textinput.Model
	currentTable int
	tables       []Table
	createTable  CreateTable
	createColumn CreateColumn
	queryText    string
}

type Table struct {
	cursor  int
	name    string
	columns []string
	types   []string
	data    [][]any
	addRow  AddRow
}

type AddRow struct {
	cursor int
	data   []string
}

func NewModel(db database.Database) Model {
	ti := textinput.New()
	ti.Focus()
	ti.Cursor.SetMode(cursor.CursorStatic)

	tables := GetTables(db)

	currentTable := -1
	if len(tables) > 0 {
		currentTable = 0
	}

	return Model{
		cursor:       []int{0, 0, 0},
		d:            db,
		textInput:    ti,
		currentTable: currentTable,
		tables:       tables,
		createTable:  CreateTable{},
		createColumn: CreateColumn{},
	}
}

func (m *Model) UpdateTable() {
	m.tables = GetTables(m.d)

	currentTable := -1
	if len(m.tables) > 0 {
		currentTable = 0
	}
	m.currentTable = currentTable
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
