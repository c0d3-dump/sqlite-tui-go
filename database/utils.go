package database

import (
	"fmt"
)

func (d Database) GetTables() []string {
	rows := d.ExecQueryRows("SELECT name FROM sqlite_master WHERE type='table'")

	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		if table != "sqlite_sequence" {
			tables = append(tables, table)
		}
	}

	return tables
}

func (d Database) LenTable(name string) int {
	q := fmt.Sprintf("SELECT count(*) FROM %s", name)
	row := d.ExecQueryRow(q)

	var len int
	row.Scan(&len)
	return len
}
