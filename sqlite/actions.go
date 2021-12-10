package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// doQuery does a query on the sqlite database
// but returns no data
func doQuery(query string, dbc *sql.DB) (err error) {
	var statement *sql.Stmt
	// Prepare SQL Statement
	if statement, err = dbc.Prepare(query); err != nil {
		return err
	}
	defer statement.Close()
	// Execute SQL Statement
	statement.Exec()

	return nil
}

// truncateTable flushes an entire table using the TRUNCATE optimizer
func truncateTable(table string, dbc *sql.DB) (err error) {
	query := "DELETE FROM" + table + ";"

	if err := doQuery(query, dbc); err != nil {
		return err
	}

	return nil
}
