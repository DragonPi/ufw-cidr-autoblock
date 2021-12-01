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
	"log"

	"github.com/spf13/cobra"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

var exclgithub string

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
		if err := u.PrepSQLite(verbose); err != nil {
			log.Fatal(err)
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

	applyCmd.PersistentFlags().StringVar(&exclgithub, "exclude-github", "", "exclude zones provided by GitHub API")
	applyCmd.PersistentFlags().BoolP("update-zones", "u", false, "update the zone files (will download/refresh zone files from internet)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printApply() {
	// Download/refresh zones from internet if requested
	// readout json file with exclusions
	// readout json fil with inclusions
	// cache json info into sqlite
	// backup previous settings
	//
	// apply new settings with data from sqlite and downloaded zone files + info from ini file
	// reload ufw
}

// appendZones adds CIDR zones to the "automatic_entries" sections of the specified json
func appendZones(jsonFile string, section string, subsection string, data []string) (err error) {

	return
}

// removeZones removes CIDR zones from the "automatic_entries" sections of the specified json
func removeZones(jsonFile string, section string, subsection string) (err error) {

	return
}
