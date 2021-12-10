package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// CacheCountries caches countrycode records per continent from the JSON files in sqlite database
func CacheCountries(countries *u.Continents) (err error) {
	var dbc *sql.DB

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	zones := countries.Afrika.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Afrika", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.Asia.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Asia", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.Europe.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Europe", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.NorthAmerica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "North-America", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.SouthAmerica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "South-America", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.Oceania.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Oceania", 0, dbc); err != nil {
			return err
		}
	}

	zones = countries.Antartica.Zones
	for _, code := range zones {
		if err = blockCountry(code, "Antartica", 0, dbc); err != nil {
			return err
		}
	}

	return nil
}

// allowCountry adds a record to the allowed country sqlite cache
func allowCountry(countryCode string, countryLong string, manual int, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO allowedcountries(country, country_long, manual)`
	stmt := `UPSERT INTO allowedcountries(country, country_long, manual)`

	data := []interface{}{countryCode, countryLong, manual}
	if err = insertZoneRecord(stmt, data, dbc); err != nil {
		return err
	}

	return nil
}

// blockCountry adds a record to the blocked country sqlite cache
func blockCountry(countryCode string, countryLong string, manual int, dbc *sql.DB) (err error) {
	// SQL statement for inserting the record
	//stmt := `INSERT INTO blockedcountries(country, country_long, manual)`
	stmt := `UPSERT INTO blockedcountries(country, country_long, manual)`

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
