package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
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

// Allowedzones contains zones that should be explicitely allowed
// see exclzones.json file
type Allowedzones struct {
	CIDR []string `json:"allowed"`
}

// Blockedzones contains zones that should be explicitely blocked
// see inclzones.json file
type Blockedzones struct {
	CIDR []string `json:"blocked"`
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

	return nil
}

// unmarshallAllowedZones reads the info from provided json file
// and fills struct with the needed info.
func UnmarshallAllowedZones(allowedZones *Allowedzones) (err error) {
	var (
		exclFile *os.File
		body     []byte
	)

	if exclFile, err = os.Open(exclLocation()); err != nil {
		return fmt.Errorf("opening %s: %w", ExclName(), err)
	}
	defer exclFile.Close()

	if body, err = ioutil.ReadAll(exclFile); err != nil {
		return fmt.Errorf("reading %s: %w", ExclName(), err)
	}

	json.Unmarshal(body, &allowedZones)

	return nil
}

// unmarshallBlockedZones reads the info from provided json file
// and fills struct with the needed info.
func UnmarshallBlockedZones(blockedZones *Blockedzones) (err error) {
	var (
		exclFile *os.File
		body     []byte
	)

	if exclFile, err = os.Open(inclLocation()); err != nil {
		return fmt.Errorf("opening %s: %w", InclName(), err)
	}
	defer exclFile.Close()

	if body, err = ioutil.ReadAll(exclFile); err != nil {
		return fmt.Errorf("reading %s: %w", InclName(), err)
	}

	json.Unmarshal(body, &blockedZones)

	return nil
}

// exclDir is a shorthand to get the exclusions folder
func exclDir() string {
	return viper.GetString("exclusions.exclusionsLocation")
}

// exclName is a shorthand to get the exclusions file name
func ExclName() (exclName string) {
	if viper.GetString("exclusions.exclusionsHidden") == "yes" {
		exclName = "."
	}

	exclName += viper.GetString("defaults.filePrefix") + "-" + viper.GetString("exclusions.exclusionsName")

	if suffix := viper.GetString("defaults.fileSuffix"); suffix != "" {
		exclName += "-" + suffix
	}

	exclName += ".json"

	return
}

// exclLocation is a shorthand to get the full exclusions file path, filename included
func exclLocation() (exclLocation string) {
	rootdir := RootDir()
	exclDir := exclDir()
	exclName := ExclName()

	exclLocation = filepath.Join(rootdir, exclDir, exclName)

	return
}

// inclDir is a shorthand to get the inclusions folder
func inclDir() string {
	return viper.GetString("inclusions.inclusionsLocation")
}

// inclName is a shorthand to get the inclusions file name
func InclName() (inclName string) {
	if viper.GetString("inclusions.inclusionsHidden") == "yes" {
		inclName = "."
	}

	inclName += viper.GetString("defaults.filePrefix") + "-" + viper.GetString("inclusions.inclusionsName")

	if suffix := viper.GetString("defaults.fileSuffix"); suffix != "" {
		inclName += "-" + suffix
	}

	inclName += ".json"

	return
}

// inclLocation is a shorthand to get the full inclusions file path, filename included
func inclLocation() (inclLocation string) {
	rootdir := RootDir()
	inclDir := inclDir()
	inclName := InclName()

	inclLocation = filepath.Join(rootdir, inclDir, inclName)

	return
}
