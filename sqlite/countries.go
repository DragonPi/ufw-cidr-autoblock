package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// CacheCountries caches countrycode records per continent from the JSON files in sqlite database
func CacheBlockedCountries(countries *u.Continents) (err error) {
	var dbc *sql.DB

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	zones := countries.Afrika.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Afrika", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Asia.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Asia", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Europe.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Europe", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.NorthAmerica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "North-America", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.SouthAmerica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "South-America", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Oceania.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Oceania", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Antartica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Antartica", "No", dbc); err != nil {
			return err
		}
	}

	return nil
}

func CacheAllowedCountries(countries *u.Continents) (err error) {
	var dbc *sql.DB

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	zones := countries.Afrika.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "Afrika", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Asia.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "Asia", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Europe.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "Europe", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.NorthAmerica.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "North-America", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.SouthAmerica.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "South-America", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Oceania.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "Oceania", "No", dbc); err != nil {
			return err
		}
	}

	zones = countries.Antartica.Unblocked
	for _, code := range zones {
		if err = allowCountry(code, "Antartica", "No", dbc); err != nil {
			return err
		}
	}

	return nil
}

// allowCountry adds a record to the allowed country sqlite cache
func allowCountry(countryCode string, countryLong string, manual string, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO allowedcountries(country, country_long, manual)`
	stmt := `INSERT INTO 
	allowedcountries(country, country_long, manual)
	VALUES (?, ?, ?)
	ON CONFLICT(country) DO 
	UPDATE SET
	country_long='` + countryLong + `',
	manual='` + manual + `'`

	data := []interface{}{countryCode, countryLong, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// blockCountry adds a record to the blocked country sqlite cache
func blockCountry(countryCode string, countryLong string, manual string, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO blockedcountries(country, country_long, manual)`
	stmt := `INSERT INTO 
	blockedcountries(country, country_long, manual)
	VALUES (?, ?, ?)
	ON CONFLICT(country) DO 
	UPDATE SET
	country_long='` + countryLong + `',
	manual='` + manual + `'`

	data := []interface{}{countryCode, countryLong, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// ListAllowedZones prompts the allowed countries as cached in sqlite
func ListAllowedCountries() (err error) {

	return
}

// ListBlockedCountries prompts the blocked countries as cached in sqlite
func ListBlockedCountries() (err error) {

	return
}
