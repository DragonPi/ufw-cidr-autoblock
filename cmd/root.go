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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// This struct holds the CLI attributes given
type Cmd struct {
	Root          *cobra.Command
	cmds          map[string]*cobra.Command
	verbose       bool
	dryRun        bool
	exclcountries string
	exclzones     string
	exclgithub    string
	inclzones     string
}

// Instantiate the root CLI
func New() *Cmd {
	c := &Cmd{
		Root: &cobra.Command{
			Use:   "ufw-cidr-autoblock",
			Short: "Generate firewall rules and apply with ufw",
			Long: `This tool generates firewall rules based on CIDR zone files downloaded from the internet.
By default it will block all known IP-zones. Two config files exist for exclusions.
- .uca-exclcountry.json with country zones to exclude (basically what you like to allow)
- .uca-exclzone.json with CIDR zones to exclude (basically a subset of CIDR you like to allow)

CIDR address block lists provided by http://ipverse.net

The .uca-exclzone.json can be appended automatically with for example:
- Github webhooks IP zones fetched from https://api.github.com/meta`,
		},
	}

	c.setup()

	return c
}

func (c *Cmd) Execute() {
	c.Root.Execute()
}

func (c *Cmd) setup() {
	c.cmds = make(map[string]*cobra.Command)

	c.initConfig()

	// Global flags
	c.Root.PersistentFlags().BoolP("verbose", "v", false, "verbose")
	// Local flags
	// c.Root.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	c.addVersionCmd()
	c.addApplyCmd()
	c.addRevertCmd()

	// Add root commands
	// ufw-cidr-autoblock <root-command>
	for _, ccmd := range c.cmds {
		c.Root.AddCommand(ccmd)
	}

	// Set to true for production application.
	// false enables the bash completion module to generate bash_completion script.
	// Output needs to be copied to /etc/bash_completion.d/lnx-database-tools
	c.Root.CompletionOptions.DisableDefaultCmd = true
	// When set to false additional info is printed for each command during bash_completion
	c.Root.CompletionOptions.DisableDescriptions = true
}

// initConfig reads in config files and ENV variables if set.
func (c *Cmd) initConfig() {
	if c.exclcountries != "" {
		// Use config file from the flag.
		viper.SetConfigFile(c.exclcountries)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".uca-exclcountry" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".uca-exclcountries")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
