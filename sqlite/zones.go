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

	// Cleanup JSON cache in sqlite db
	if err = cacheCleanupAllowed(dbc); err != nil {
		return err
	}

	for _, zone := range CIDR {
		if err = allowZone(zone, u.ExclName(), "No", dbc); err != nil {
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

	// Cleanup JSON cache in sqlite db
	if err = cacheCleanupBlocked(dbc); err != nil {
		return err
	}

	for _, zone := range CIDR {
		if err = blockZone(zone, u.InclName(), "No", dbc); err != nil {
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

	// Iterate and fetch the records from result cursor
	for row.Next() {
		var id int
		var zone string
		var reference string
		var manual string
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

	// Iterate and fetch the records from result cursor
	for row.Next() {
		var id int
		var zone string
		var reference string
		var manual string
		row.Scan(&id, &zone, &reference, &manual)
		u.Info.Println("Blockedzone: ", zone, " ", reference, " ", manual)
	}

	return nil
}

// allowZone adds a record to the allowed zone sqlite cache
func allowZone(zone string, reference string, manual string, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	stmt := `INSERT INTO 
		allowedzones(zone, reference, manual)
		VALUES (?, ?, ?)
		ON CONFLICT(zone) DO 
		UPDATE SET
		reference='` + reference + `',
		manual='` + manual + `'`

	data := []interface{}{zone, reference, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// blockZone adds a record to the blocked zone sqlite cache
func blockZone(zone string, reference string, manual string, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	stmt := `INSERT INTO 
	blockedzones(zone, reference, manual)
	VALUES (?, ?, ?)
	ON CONFLICT(zone) DO 
	UPDATE SET
	reference='` + reference + `',
	manual='` + manual + `'`

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
	if statement, err = dbc.Prepare(query); err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(data[0], data[1], data[2]); err != nil {
		return err
	}

	return nil
}

// cacheCleanupAllowed cleans the JSON cache in sqlite database
// otherwise old JSON zones could remain present in cache.
func cacheCleanupAllowed(dbc *sql.DB) (err error) {
	query := "DELETE FROM allowedzones WHERE reference like '%" + u.ExclName() + "%';"
	if err = doQuery(query, dbc); err != nil {
		return err
	}

	return nil
}

// cacheCleanupBlocked cleans the JSON cache in sqlite database
// otherwise old JSON zones could remain present in cache.
func cacheCleanupBlocked(dbc *sql.DB) (err error) {
	query := "DELETE FROM blockedzones WHERE reference like '%" + u.InclName() + "%';"
	if err = doQuery(query, dbc); err != nil {
		return err
	}

	return nil
}
