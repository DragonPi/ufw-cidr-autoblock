/*
Copyright © 2021 Koen Kumps

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	sql "github.com/DragonPi/ufw-cidr-autoblock/sqlite"
	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

var (
	exclgithub   bool
	updateRemote bool
	updateLocal  bool
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply rules based on CIDR zone files",
	Long: `This function will apply a set of firewall rules based on zone files with CIDR address blocks.
When accompanied by the dry-run flag it will generate a test.rules file which can be validated afterwards, but will not apply anything.
By default it will use zones files already present.  Add the update-zones flag to have these updated/downlaod from the internet.`,
	Example: "ufw-cidr-autoblock apply",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := sql.PrepSQLite(verbose); err != nil {
			u.Error.Fatal(err)
		}
		printApply()
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	applyCmd.PersistentFlags().BoolVarP(&exclgithub, "exclude-github", "", false, "exclude zones provided by GitHub API from firewall")
	applyCmd.PersistentFlags().BoolVarP(&updateRemote, "update-remote", "", false, "update the zone files (will download/refresh zone files from internet)")
	applyCmd.PersistentFlags().BoolVarP(&updateLocal, "update-local", "", false, "update the zone files (will download/refresh zone files from internet)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printApply() {
	var (
		unblockedZones []string
		err            error
	)

	continents := u.Continents{}
	metaData := u.GitHub{}
	allowedZones := u.Allowedzones{}
	blockedZones := u.Blockedzones{}

	if updateLocal {
		if err = updateJSON(&continents, &allowedZones, &blockedZones); err != nil {
			u.Error.Fatalln(err)
		}
	}

	if updateRemote {
		if err = updateGitHub(&metaData); err != nil {
			u.Error.Fatalln(err)
		}

		if err = updateZones(&continents); err != nil {
			u.Error.Fatalln(err)
		}
	}

	//////
	unblockedZones = u.MakeCountryUnblockArray(&continents)
	fmt.Printf("%+v\n", unblockedZones)
	// readout json file with exclusions
	// readout json file with inclusions
	// cache json info into sqlite
	// backup previous settings
	//
	// apply new settings with data from sqlite and downloaded zone files + info from ini file
	// reload ufw
}

// updateJSON refreshes sqlite cache with the rules/countries stored in .json
func updateJSON(continents *u.Continents, allowedZones *u.Allowedzones, blockedZones *u.Blockedzones) (err error) {
	// Write needed info in struct
	if err = u.UnmarshallCountries(&continents); err != nil {
		return err
	} else {
		if err = sql.CacheBlockedCountries(continents); err != nil {
			return err
		}
		if err = sql.CacheAllowedCountries(continents); err != nil {
			return err
		}
	}
	// Write needed info in struct
	if err = u.UnmarshallAllowedZones(&allowedZones); err != nil {
		return err
	} else {
		if err = sql.CacheAllowedZones(allowedZones); err != nil {
			return err
		}
	}
	// Write needed info in struct
	if err = u.UnmarshallBlockedZones(&blockedZones); err != nil {
		return err
	} else {
		if err = sql.CacheBlockedZones(blockedZones); err != nil {
			return err
		}
	}

	return nil
}

// updateGitHub refreshes sqlite cache with the data from internet (https://api.github.com/meta)
func updateGitHub(metaData *u.GitHub) (err error) {
	if err = u.DownloadGitHubIP(&metaData); err != nil {
		// info is cached in sqlite, so we can still
		// apply the rules but with "old" data, therefore only warn
		u.Warning.Println(err)
	} else {
		// Download successful so cache it in sqlite db
		if err = sql.CacheAllowedGitHub(metaData); err != nil {
			return err
		}
	}

	return nil
}

// updateZones downloads zone files and refreshes sqlite cache with the zone
// files from internet (http://ipverse.net/ipblocks/data/countries/)
func updateZones(continents *u.Continents) (err error) {
	countryZones := u.MakeCountryZoneArray(continents)

	if err = u.DownloadZoneFiles(countryZones); err != nil {
		u.Warning.Println(err)
	} else {
		// Download successful so cache it in sqlite db
		if err = sql.CacheZoneFiles(); err != nil {
			return err
		}
	}

	return nil
}
