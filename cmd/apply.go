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
)

func (c *Cmd) addApplyCmd() {
	c.cmds["apply"] = &cobra.Command{
		Use:   "apply",
		Short: "Apply rules based on CIDR zone files",
		Long: `This function will apply a set of firewall rules based on zone files with CIDR address blocks.
When accompanied by the dry-run flag it will generate a test.rules file which can be validated afterwards, but will not apply anything.
By default it will use zones files already present.  Add the update-zones flag to have these updated/downlaod from the internet.`,
		Example: "ufw-cidr-autoblock apply",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			c.printApply()
		},
	}
	c.cmds["apply"].PersistentFlags().StringVar(&c.exclcountries, "countries-allow", "", "config file (default is $HOME/.uca-exclcountries.json)")
	c.cmds["apply"].PersistentFlags().StringVar(&c.exclzones, "zones-allow", "", "config file (default is $HOME/.uca-exclzones.json)")
	c.cmds["apply"].PersistentFlags().StringVar(&c.inclzones, "zones-block", "", "config file (default is $HOME/.uca-inclzones.json)")
	c.cmds["apply"].PersistentFlags().BoolP("dry-run", "d", false, "create test.rules but do not apply.")
	c.cmds["apply"].PersistentFlags().BoolP("update-zones", "u", false, "update the zone files (will download/refresh zone files from internet)")
	c.cmds["apply"].PersistentFlags().StringVar(&c.exclgithub, "exclude-github", "", "exclude zones provided by GitHub API")
}

// printRevert prints output from the revert function
func (c *Cmd) printApply() {
	fmt.Println("apply called")
}
