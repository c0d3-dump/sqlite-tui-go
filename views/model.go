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
	ids     []int
	data    [][]any
}

func NewModel(db database.Database) Model {
	ti := textinput.New()
	ti.Focus()
	ti.Cursor.SetMode(cursor.CursorStatic)

	var tables []Table

	for _, tableName := range db.GetTables() {
		var cols []string

		for _, tableInfo := range GetTableInfo(db, tableName) {
			if tableInfo.name != "id" {
				cols = append(cols, tableInfo.name)
			}
		}

		rows := db.ExecQueryRows("SELECT * FROM " + tableName + ";")

		var ids []int
		var data [][]any

		colLen := len(cols) + 1
		for rows.Next() {
			// TODO: convert it to any to make this code run
			t := make([]any, colLen)
			tptr := make([]any, colLen)

			for i := 0; i < colLen; i++ {
				tptr[i] = &t[i]
			}

			rows.Scan(tptr...)
			ids = append(ids, int(t[0].(int64)))
			data = append(data, t[1:colLen])
		}

		tables = append(tables, Table{
			name:    tableName,
			columns: cols,
			ids:     ids,
			data:    data,
		})
	}

	// fmt.Print(tables)

	return Model{
		cursor:       []int{0, 0, 0},
		d:            db,
		textInput:    ti,
		currentTable: 0,
		tables:       tables,
		createTable:  CreateTable{},
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
