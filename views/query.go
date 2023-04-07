package views

import (
	"fmt"
	"sqlite-tui-go/database"
	"strings"
)

type TableInfo struct {
	name    string
	dtype   string
	notnull bool
	dval    any
	pk      bool
}

func GetTableInfo(d database.Database, name string) []TableInfo {
	q := fmt.Sprintf("PRAGMA table_info(%s)", name)

	rows := d.ExecQueryRows(q)

	table_map := []TableInfo{}
	for rows.Next() {
		t := TableInfo{}
		var cid int
		rows.Scan(&cid, &t.name, &t.dtype, &t.notnull, &t.dval, &t.pk)
		table_map = append(table_map, t)
	}

	return table_map
}

func (m Model) CreateTable(t CreateTable) {
	var query string
	query = "CREATE TABLE " + t.name + "("

	colLen := len(t.columns)

	query += "id INTEGER PRIMARY KEY AUTOINCREMENT,"

	for i, col := range t.columns {
		query += col.name + " " + col.dtype
		if i < colLen-1 {
			query += ","
		}
	}

	query += ");"

	m.d.ExecStatement(query)
}

func GetTables(db database.Database) []Table {
	var tables []Table

	for _, tableName := range db.GetTables() {
		var cols []string
		var types []string

		for _, tableInfo := range GetTableInfo(db, tableName) {
			cols = append(cols, tableInfo.name)
			types = append(types, tableInfo.dtype)
		}

		rows := db.ExecQueryRows("SELECT * FROM " + tableName + ";")

		var data [][]any

		colLen := len(cols)
		for rows.Next() {
			t := make([]any, colLen)
			tptr := make([]any, colLen)

			for i := 0; i < colLen; i++ {
				tptr[i] = &t[i]
			}

			rows.Scan(tptr...)
			data = append(data, t)
		}

		tables = append(tables, Table{
			name:    tableName,
			columns: cols,
			types:   types,
			data:    data,
			addRow: AddRow{
				data: make([]string, colLen-1),
			},
		})
	}

	return tables
}

func (m Model) AddRow() {
	table := m.tables[m.currentTable]

	var modifiedData []string
	for i, d := range table.addRow.data {
		switch table.types[i+1] {
		case "TEXT":
			modifiedData = append(modifiedData, "'"+d+"'")
		case "INTEGER":
			modifiedData = append(modifiedData, d)
		}
	}

	m.d.ExecStatement("INSERT INTO " + table.name + " (" + strings.Join(table.columns[1:], ",") + ") VALUES (" + strings.Join(modifiedData, ",") + ");")
}
