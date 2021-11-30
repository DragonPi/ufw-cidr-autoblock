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

	"github.com/spf13/cobra"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// These variables are populated during build (-ldflags -> see main.go)
var (
	Version    string
	Buildstamp string
	Githash    string
)

// Struct to hold the application signature so it can be easily used elsewhere
type Signature struct {
	Version    string
	Buildstamp string
	Githash    string
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Display application version",
	Long:    "This command returns the version.",
	Example: "ufw-cidr-autoblock version",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := u.PrepSQLite(verbose); err != nil {
			log.Fatal(err)
		}
		printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// printVersion returns the version of the application
func printVersion() {
	s := &Signature{
		Version:    Version,
		Buildstamp: Buildstamp,
		Githash:    Githash,
	}

	fmt.Printf("Last build: %s\n", s.Buildstamp)
	fmt.Printf("App Version: %s\n", s.Version)
	fmt.Printf("Githash: %s\n", s.Githash)
}
