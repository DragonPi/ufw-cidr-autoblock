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
package main

import "github.com/DragonPi/ufw-cidr-autoblock/cmd"

// Build the executable with the following command (adapt main.Version as needed)
// go build -ldflags "-X github.com/DragonPi/ufw-cidr-autoblock/cmd.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` \
//                    -X github.com/DragonPi/ufw-cidr-autoblock/cmd.Githash=`git rev-parse HEAD` \
//                    -X github.com/DragonPi/ufw-cidr-autoblock/cmd.Version=0.5"

func main() {
	cmd.Execute()
}
