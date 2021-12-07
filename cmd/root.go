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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

var (
	cfgFile string
	verbose bool
	dryrun  bool
)

// Info writes logs in the color blue with "INFO: " as prefix
var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Warning writes logs in the color yellow with "WARNING: " as prefix
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)

// Error writes logs in the color red with "ERROR: " as prefix
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

// Debug writes logs in the color cyan with "DEBUG: " as prefix
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ufw-cidr-autoblock",
	Short: "Generate firewall rules and apply with ufw",
	Long: `This tool generates firewall rules based on CIDR zone files downloaded from the internet.
By default it will block all known IP-zones. Two config files exist for exclusions.
- .uca-exclcountry.json with country zones to exclude (basically what you like to allow)
- .uca-exclzone.json with CIDR zones to exclude (basically a subset of CIDR you like to allow)

CIDR address block lists provided by http://ipverse.net

The .uca-exclzone.json can be appended automatically with for example:
- Github webhooks IP zones fetched from https://api.github.com/meta`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.uca-config.ini)")
	rootCmd.PersistentFlags().BoolVar(&dryrun, "dry-run", false, "create test.rules but do not apply.")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".uca-config" (without extension)
		// secondly in local directory (next to executable)
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("ini")
		viper.SetConfigName(".uca-config")
	}

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if u.IsTerminal() && verbose {
			log.Printf("Using config file: %v\n", viper.ConfigFileUsed())
		}
	}
}
