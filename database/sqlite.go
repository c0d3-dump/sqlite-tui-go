package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Database struct {
	db *sql.DB
}

func (d *Database) InitDatabase(database string) {
	db, err := sql.Open("sqlite", database)

	if err != nil {
		log.Println(database)
		log.Fatal("error connecting : ", err)
	}

	d.db = db
}

func (d Database) States() sql.DBStats {
	return d.db.Stats()
}

func (d Database) ExecStatement(statement string, args ...any) sql.Result {
	st, err := d.db.Prepare(statement)

	if err != nil {
		log.Println(statement)
		log.Fatal("error preparing : ", err)
	}

	out, err := st.Exec(args...)

	if err != nil {
		log.Println(statement)
		log.Fatal("error executing : ", err)
	}

	return out
}

func (d Database) ExecQueryRows(query string, args ...any) *sql.Rows {
	rows, err := d.db.Query(query, args...)

	if err != nil {
		log.Println(query)
		log.Fatal("error querying : ", err)
	}

	return rows
}

func (d Database) ExecQueryRow(query string, args ...any) *sql.Row {
	row := d.db.QueryRow(query, args...)

	err := row.Err()
	if err != nil {
		log.Panicln(query)
		log.Fatal("error quering : ", err)
	}

	return row
}
