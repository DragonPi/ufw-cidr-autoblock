package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Continents contains the ISO-3166-1 style country codes
// extracted from .uca-exclcountries.json divided by continent
type Continents struct {
	Afrika       Zones `json:"AFRICA"`
	Asia         Zones `json:"ASIA"`
	Europe       Zones `json:"EUROPE"`
	NorthAmerica Zones `json:"NORTH_AMERICA"`
	SouthAmerica Zones `json:"SOUTH_AMERICA"`
	Oceania      Zones `json:"OCEANIA"`
	Antartica    Zones `json:"ANTARTICA"`
}

// Zones contains for each continent (referenced in Contintents)
// the available zone files (Zones) and the allowed zones (Unblocked)
type Zones struct {
	Zones     []string `json:"zones"`
	Unblocked []string `json:"unblocked"`
}

// unmarshallCountries reads the info from provided json file
// and fills struct with the needed info.
func UnmarshallCountries(continents *Continents) (err error) {
	var (
		exclFile *os.File
		body     []byte
	)

	if exclFile, err = os.Open("exceptions/.uca-exclcountries.json"); err != nil {
		return fmt.Errorf("opening .uca-exclcountries.json: %w", err)
	}
	defer exclFile.Close()

	if body, err = ioutil.ReadAll(exclFile); err != nil {
		return fmt.Errorf("reading .uca-exclcountries.json: %w", err)
	}

	json.Unmarshal(body, &continents)

	return
}
