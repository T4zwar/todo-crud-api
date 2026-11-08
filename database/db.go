package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() error {
	var err error

	DB, err = sql.Open("sqlite", "./todos.db")
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	query := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT FALSE
    );`

	_, err = DB.Exec(query)
	return err
}
