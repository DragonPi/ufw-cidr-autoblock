package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// CacheBlockedZones caches allowed CIDR records from the JSON files in sqlite database
func CacheAllowedZones(allowedZones *u.Allowedzones) (err error) {
	var dbc *sql.DB

	CIDR := allowedZones.CIDR

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	for _, zone := range CIDR {
		if err = allowZone(zone, u.exclName(), 0, dbc); err != nil {
			return err
		}
	}

	return nil
}

// CacheBlockedZones caches blocked CIDR records from the JSON files in sqlite database
func CacheBlockedZones(blockedZones *u.Blockedzones) (err error) {
	var dbc *sql.DB

	CIDR := blockedZones.CIDR

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	for _, zone := range CIDR {
		if err = blockZone(zone, inclName(), 0, dbc); err != nil {
			return err
		}
	}

	return nil
}

func CacheZoneFiles() (err error) {

	return nil
}

// ListAllowedZones prompts the allowed zones as cached in sqlite
func ListAllowedZones() (err error) {
	var (
		row *sql.Rows
		dbc *sql.DB
	)

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	if row, err = dbc.Query("SELECT * FROM allowedzones ORDER BY reference"); err != nil {
		return err
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var zone string
		var reference string
		var manual int
		row.Scan(&id, &zone, &reference, &manual)
		u.Info.Println("Allowed zone: ", zone, " ", reference, " ", manual)
	}

	return nil
}

// ListBlockedZones prompts the blocked zones as cached in sqlite
func ListBlockedZones() (err error) {
	var (
		row *sql.Rows
		dbc *sql.DB
	)

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	if row, err = dbc.Query("SELECT * FROM blockedzones ORDER BY reference"); err != nil {
		return err
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var zone string
		var reference string
		var manual int
		row.Scan(&id, &zone, &reference, &manual)
		u.Info.Println("Blockedzone: ", zone, " ", reference, " ", manual)
	}

	return nil
}

// allowZone adds a record to the allowed zone sqlite cache
func allowZone(zone string, reference string, manual int, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO allowedzones(zone, reference, manual)`
	stmt := `UPSERT INTO allowedzones(zone, reference, manual)`

	data := []interface{}{zone, reference, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// blockZone adds a record to the blocked zone sqlite cache
func blockZone(zone string, reference string, manual int, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO blockedzones(zone, reference, manual)`
	stmt := `UPSERT INTO blockedzones(zone, reference, manual)`

	data := []interface{}{zone, reference, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// insertZoneRecord inserts a zone record in the database
func insertZoneRecord(query string, data []interface{}, dbc *sql.DB) (err error) {
	var statement *sql.Stmt

	// Prepare SQL statement
	// This is to avoid SQL injections
	querySuffix := " VALUES (?, ?, ?)"
	query += querySuffix
	if statement, err = dbc.Prepare(query); err != nil {
		return err
	}

	if _, err = statement.Exec(data[0], data[1], data[2]); err != nil {
		return err
	}

	return nil
}
