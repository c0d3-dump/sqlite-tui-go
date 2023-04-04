package views

import (
	"fmt"
	"sqlite-tui-go/database"
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
