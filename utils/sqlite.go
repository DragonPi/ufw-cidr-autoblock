package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

// CacheGitHub caches the info from the GitHub meta endpoint into sqlite database
func CacheGitHub(metaData *GitHub) (err error) {
	dbc := openSQLite()
	// Defer closing the database connection
	defer dbc.Close()

	allowedGitHub := strings.Split(viper.GetString("exclusions.GitHub"), ",")
	for _, reference := range allowedGitHub {
		switch reference {
		case "hooks":
			for _, zone := range metaData.Hooks {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "web":
			for _, zone := range metaData.Web {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "api":
			for _, zone := range metaData.API {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "git":
			for _, zone := range metaData.Git {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "packages":
			for _, zone := range metaData.Packages {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "pages":
			for _, zone := range metaData.Pages {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "importer":
			for _, zone := range metaData.Importer {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "actions":
			for _, zone := range metaData.Actions {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		case "dependabot":
			for _, zone := range metaData.Dependabot {
				allowZone(zone, "GitHub-"+reference, 0, dbc)
			}
		}
	}

	return
}

// ListAllowedZones prompts the allowed zones as cached in sqlite
func ListAllowedZones() (err error) {
	var row *sql.Rows

	dbc := openSQLite()
	// Defer closing the database connection
	defer dbc.Close()

	if row, err = dbc.Query("SELECT * FROM allowzones ORDER BY reference"); err != nil {
		Error.Fatal(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var zone string
		var reference string
		var manual int
		row.Scan(&id, &zone, &reference, &manual)
		Info.Println("Allowed zone: ", zone, " ", reference, " ", manual)
	}

	return
}

// ListAllowedZones prompts the allowed zones as cached in sqlite
func ListBlockedZones() (err error) {
	var row *sql.Rows

	dbc := openSQLite()
	// Defer closing the database connection
	defer dbc.Close()

	if row, err = dbc.Query("SELECT * FROM blockzones ORDER BY reference"); err != nil {
		Error.Fatal(err)
	}
	defer row.Close()

	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var zone string
		var reference string
		var manual int
		row.Scan(&id, &zone, &reference, &manual)
		Info.Println("Bockedzone: ", zone, " ", reference, " ", manual)
	}

	return
}

// Prepares the SQLite database for storing exceptions
func PrepSQLite(verbose bool) (err error) {
	var created bool

	if created, err = ensureSQLite(verbose); err != nil {
		return err
	}
	if created {
		if err = populateSQLite(verbose); err != nil {
			return err
		}
	}

	return nil
}

// allowZone adds a records to the allowed zone sqlite cache
func allowZone(zone string, reference string, manual int, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	stmt := `INSERT INTO allowzones(zone, reference, manual)`

	data := []interface{}{zone, reference, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return
}

// ensureSQLite makes sure the database file exists upon application start
func ensureSQLite(verbose bool) (created bool, err error) {
	var file *os.File

	rootdir := RootDir()

	dbLocation := filepath.Join(rootdir,
		viper.GetString("database.dblocation"),
		viper.GetString("database.dbname"),
	)

	if DestinationExists(dbLocation) {
		if IsTerminal() && verbose {
			Info.Printf("%v already exists, skip creation\n", viper.GetString("database.dbname"))
		}
		return false, nil
	}

	if IsTerminal() || verbose {
		Info.Printf("db not found, creating %v...\n", viper.GetString("database.dbname"))
	}

	MakeDestination(dbLocation)
	// Create sqlite file
	if file, err = os.Create(dbLocation); err != nil {
		return false, err
	}

	file.Close()

	if IsTerminal() || verbose {
		Info.Printf("%v created.\n", viper.GetString("database.dbname"))
	}

	return true, nil
}

// populateSQLite creates an empty database
func populateSQLite(verbose bool) (err error) {
	var stmt string

	dbc := openSQLite()
	// Defer closing the database connection
	defer dbc.Close()

	// SQL Statement for creating table containing
	// explicitely blocked zones
	stmt = `CREATE TABLE blockzones (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"zone" TEXT,
		"reference" TEXT,
		"manual" INTEGER
	  );`

	if verbose {
		Info.Println("Creating table blockzones...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		Error.Fatalln(err)
	}

	if verbose {
		Info.Println("Created table blockzones...")
	}

	// SQL Statement for creating table containing
	// explicitely allowed zones
	stmt = `CREATE TABLE allowzones (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"zone" TEXT,
		"reference" TEXT,
		"manual" INTEGER
	  );`

	if verbose {
		Info.Println("Creating table allowzones...")
	}

	if err = doQuery(stmt, dbc); err != nil {
		Error.Fatalln(err)
	}

	if verbose {
		Info.Println("Created table allowzones...")
	}

	return
}

// doQuery does a query on the sqlite database
// but returns no data
func doQuery(query string, db *sql.DB) (err error) {
	var statement *sql.Stmt
	// Prepare SQL Statement
	if statement, err = db.Prepare(query); err != nil {
		Error.Fatalln(err.Error())
	}
	// Execute SQL Statement
	statement.Exec()

	return
}

// insertZoneRecord inserts a zone record in the database
func insertZoneRecord(query string, data []interface{}, db *sql.DB) (err error) {
	var statement *sql.Stmt

	// Prepare SQL statement
	// This is to avoid SQL injections
	querySuffix := " VALUES (?, ?, ?)"
	query += querySuffix
	if statement, err = db.Prepare(query); err != nil {
		return err
	}

	fmt.Printf("%v %v %v\n", data[0], data[1], data[2])

	if _, err = statement.Exec(data[0], data[1], data[2]); err != nil {
		return err
	}

	return
}

// openSQLite opens the sqlite database at the
// location configured in the config file
func openSQLite() (dbc *sql.DB) {
	rootdir := RootDir()

	dbLocation := filepath.Join(rootdir,
		viper.GetString("database.dblocation"),
		viper.GetString("database.dbname"),
	)

	// Open the sqlite File
	dbc, _ = sql.Open("sqlite3", dbLocation)

	return
}
