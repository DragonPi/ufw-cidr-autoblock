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

func (c *Cmd) addVersionCmd() {
	c.cmds["version"] = &cobra.Command{
		Use:     "version",
		Short:   "Display application version",
		Long:    "This command returns the version.",
		Example: "lnx-database-tool version",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			c.printVersion()
		},
	}
}

// printVersion returns the version of the application
func (c *Cmd) printVersion() {
	fmt.Printf("Last build: %s\n", c.signature["Buildstamp"])
	fmt.Printf("App Version: %s\n", c.signature["Version"])
	fmt.Printf("Githash: %s\n", c.signature["Githash"])
}
