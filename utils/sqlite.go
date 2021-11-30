package utils

import (
	"log"
	"os"
	"path/filepath"

	//"database/sql"

	"github.com/spf13/viper"
	//_ "github.com/mattn/go-sqlite3"
)

// Prepares the SQLite database for storing exceptions
func PrepSQLite(verbose bool) (err error) {
	var created bool

	if created, err = ensureSQLite(verbose); err != nil {
		return err
	}
	if created {
		if err = populateSQLite(); err != nil {
			return err
		}
	}

	return nil
}

func ensureSQLite(verbose bool) (created bool, err error) {
	var file *os.File

	rootdir := RootDir()

	dbLocation := filepath.Join(rootdir,
		viper.GetString("database.dblocation"),
		viper.GetString("database.dbname"),
	)

	if DestinationExists(dbLocation) {
		if IsTerminal() && verbose {
			log.Printf("%v already exists, skip creation\n", viper.GetString("database.dbname"))
		}
		return false, nil
	}

	if IsTerminal() || verbose {
		log.Printf("db not found, creating %v...\n", viper.GetString("database.dbname"))
	}

	MakeDestination(dbLocation)
	// SQLite is a file based database.
	// Create SQLite file
	if file, err = os.Create(dbLocation); err != nil {
		return false, err
	}

	file.Close()

	if IsTerminal() || verbose {
		log.Printf("%v created.\n", viper.GetString("database.dbname"))
	}

	return true, nil
}

func populateSQLite() (err error) {

	return
}
