package database

import "fmt"

func (d Database) TableInfo(name string) []Table {
	q := fmt.Sprintf("PRAGMA table_info(%s)", name)

	rows := d.ExecQueryRows(q)

	table_map := []Table{}
	for rows.Next() {
		t := Table{}
		var cid int
		rows.Scan(&cid, &t.name, &t.dtype, &t.notnull, &t.dval, &t.pk)
		table_map = append(table_map, t)
	}

	return table_map
}

func (d Database) GetTables() []string {
	rows := d.ExecQueryRows("SELECT name FROM sqlite_master WHERE type='table'")

	var tables []string
	for rows.Next() {
		var table string
		rows.Scan(&table)
		tables = append(tables, table)
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
