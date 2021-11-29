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
	"log"

	//"path/filepath"
	//"database/sql"
	//"time"

	"github.com/spf13/cobra"
	//_ "github.com/mattn/go-sqlite3"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "To test different functions during dev",
	Example: "ufw-cidr-autoblock test",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		printTest()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// printTest returns output from testing functionalities
func printTest() {
	metaData := u.GitHub{}

	if err := u.DownloadGitHubIP(&metaData); err != nil {
		log.Fatal(err)
	}

	/*
		jsonFile := filepath.Join("home", "koen", "Code", "GitHub", "ufw-cidr-autoblock", "exceptions", ".uca-exclzones.json")
		section := "GitHub"
		subsection := "hooks"

		if err := appendZones(jsonFile, section, subsection, metaData.Hooks); err != nil {
			log.Fatal(err)
		}
	*/

	/*fmt.Println(metaData.VPA)
	fmt.Printf("RSA %v\n", metaData.SSH_fp.RSA)
	fmt.Printf("ECDSA %v\n", metaData.SSH_fp.ECDSA)
	fmt.Printf("ED25519 %v\n", metaData.SSH_fp.ED25519)*/
	fmt.Println(metaData.Hooks)
	/*fmt.Println(metaData.Web)
	fmt.Println(metaData.API)
	fmt.Println(metaData.Git)
	fmt.Println(metaData.Packages)
	fmt.Println(metaData.Pages)
	fmt.Println(metaData.Importer)
	fmt.Println(metaData.Actions)
	fmt.Println(metaData.Dependabot)*/
}
