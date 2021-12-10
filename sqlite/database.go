package sqlite

import (
	"os"
	"path/filepath"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// Prepares the SQLite database for caching exceptions
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

// ensureSQLite makes sure the database file exists upon application start
func ensureSQLite(verbose bool) (created bool, err error) {
	var file *os.File

	dbName := dbName()
	dbLocation := dbLocation()

	if u.DestinationExists(dbLocation) {
		if u.IsTerminal() && verbose {
			u.Info.Printf("%v already exists, skip creation\n", dbName)
		}
		return false, nil
	}

	if u.IsTerminal() || verbose {
		u.Info.Printf("db backend not found at configured location, creating %v...\n", dbName)
	}

	u.MakeDestination(dbLocation)
	// Create sqlite file
	if file, err = os.Create(dbLocation); err != nil {
		return false, err
	}

	file.Close()

	if u.IsTerminal() || verbose {
		u.Info.Printf("%v created.\n", dbName)
	}

	return true, nil
}

// openSQLite opens the sqlite database at the
// location configured in the config file
func openSQLite() (dbc *sql.DB, err error) {
	dbLocation := dbLocation()

	if dbc, err = sql.Open("sqlite3", dbLocation); err != nil {
		return nil, err
	}

	return dbc, nil
}

// dbDir is a shorthand to get the sqlite database folder
func dbDir() (dbDir string) {
	if viper.GetString("database.dbLocationHidden") == "yes" {
		dbDir = "." + viper.GetString("database.dblocation")
	} else {
		dbDir = viper.GetString("database.dblocation")
	}

	return
}

// dbName is a shorthand to get the sqlite database name
func dbName() (dbName string) {
	dbName = viper.GetString("defaults.filePrefix") + "-" + viper.GetString("database.dbname")

	if suffix := viper.GetString("defaults.fileSuffix"); suffix != "" {
		dbName += "-" + suffix
	}

	dbName += ".db"

	return
}

// dbLocation is a shorthand to get the full sqlite database path, filename included
func dbLocation() (dbLocation string) {
	rootdir := u.RootDir()
	dbDir := dbDir()
	dbName := dbName()

	dbLocation = filepath.Join(rootdir, dbDir, dbName)

	return
}
