package utils

import (
	"strconv"

	"github.com/spf13/viper"
)

// MakeCountryZoneArray condenses all the "Zones" from the Continents struct 
// into a single array for easy iteration
func MakeCountryZoneArray(continents *Continents) (countryZones []string) {
	// Make a single array from all the zones with a known CIDR file
	countryZones = append(append(append(append(append(append(append(countryZones,
		continents.Afrika.Zones...),
		continents.Asia.Zones...),
		continents.Europe.Zones...),
		continents.NorthAmerica.Zones...),
		continents.SouthAmerica.Zones...),
		continents.Oceania.Zones...),
		continents.Antartica.Zones...)

	return
}

// MakeCountryUnblockArray condenses all the "Unblocked" from the Continents struct 
// into a single array for easy iteration
func MakeCountryUnblockArray(continents *Continents) (unblockedZones []string) {
	// Make a single array from all the zones with a known CIDR file
	unblockedZones = append(append(append(append(append(append(append(unblockedZones,
		continents.Afrika.Unblocked...),
		continents.Asia.Unblocked...),
		continents.Europe.Unblocked...),
		continents.NorthAmerica.Unblocked...),
		continents.SouthAmerica.Unblocked...),
		continents.Oceania.Unblocked...),
		continents.Antartica.Unblocked...)

	return
}

// defineConcurrency sets the concurrency level based on the config file and sets
// a default value if the value from the configfile is empty or non-existing
func defineConcurrency() (concurrency int) {
	var err error

	if viper.GetString("defaults.concurrency") == "" || viper.GetString("defaults.concurrency") == "0" {
		concurrency = 1
	} else {
		if concurrency, err = strconv.Atoi(viper.GetString("defaults.concurrency")); err != nil {
			return 1
		}
	}

	return
}
