package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// populateSQLite creates the empty database tables
func populateSQLite(verbose bool) (err error) {
	var dbc *sql.DB

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	if err = createBlockedzonesTable(verbose, dbc); err != nil {
		return err
	}
	if err = createAllowedzonesTable(verbose, dbc); err != nil {
		return err
	}
	if err = createBlockedcountriesTable(verbose, dbc); err != nil {
		return err
	}
	if err = createAllowedcountriesTable(verbose, dbc); err != nil {
		return err
	}

	return nil
}

// createBlockzonesTable executes the SQL Statement for creating
// table containing explicitely blocked zones
func createBlockedzonesTable(verbose bool, dbc *sql.DB) (err error) {
	stmt := `CREATE TABLE blockedzones (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"zone" TEXT,
		"reference" TEXT,
		"manual" INTEGER,
		UNIQUE(zone)
	  );`

	if verbose {
		u.Info.Println("Creating table blockzones...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		return err
	}

	if verbose {
		u.Info.Println("Created table blockzones...")
	}

	return nil
}

// createAllowzonesTable executes the SQL Statement for creating
// table containing explicitely allowed zones
func createAllowedzonesTable(verbose bool, dbc *sql.DB) (err error) {
	stmt := `CREATE TABLE allowedzones (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"zone" TEXT,
		"reference" TEXT,
		"manual" INTEGER,
		UNIQUE(zone)
  	);`

	if verbose {
		u.Info.Println("Creating table allowzones...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		return err
	}

	if verbose {
		u.Info.Println("Created table allowzones...")
	}

	return nil
}

// createBlockedcountriesTable executes the SQL Statement for creating
// table containing explicitely allowed countries
func createBlockedcountriesTable(verbose bool, dbc *sql.DB) (err error) {
	stmt := `CREATE TABLE blockedcountries (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"country" TEXT,
		"country_long" TEXT,
		"manual" INTEGER,
		UNIQUE(country)
  	);`

	if verbose {
		u.Info.Println("Creating table blockedcountries...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		return err
	}

	if verbose {
		u.Info.Println("Created table blockedcountries...")
	}

	return nil
}

// createAllowedcountriesTable executes the SQL Statement for creating
// table containing explicitely allowed countries
func createAllowedcountriesTable(verbose bool, dbc *sql.DB) (err error) {
	stmt := `CREATE TABLE allowedcountries (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"country" TEXT,
		"country_long" TEXT,
		"manual" INTEGER,
		UNIQUE(country)
  	);`

	if verbose {
		u.Info.Println("Creating table allowedcountries...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		return err
	}

	if verbose {
		u.Info.Println("Created table allowedcountries...")
	}

	return nil
}
