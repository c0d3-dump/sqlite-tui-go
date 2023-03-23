package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func (d *Database) InitDatabase(database string) {
	db, err := sql.Open("sqlite3", database)

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

func (d Database) ExecQuery(query string) *sql.Rows {
	rows, err := d.db.Query(query)

	if err != nil {
		log.Println(query)
		log.Fatal("error querying : ", err)
	}

	return rows
}
