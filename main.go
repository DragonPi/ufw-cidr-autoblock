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

import (
	"github.com/DragonPi/ufw-cidr-autoblock/cmd"
)

//go build -ldflags "-X main.Signature.Buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.Signature.Githash=`git rev-parse HEAD` -X main.Signature.Version=version_string"

var (
	Version    string
	Buildstamp string
	Githash    string
)

func main() {
	signature := make(map[string]string)
	signature["Version"] = Version
	signature["Buildstamp"] = Buildstamp
	signature["Githash"] = Githash
	c := cmd.New(signature)
	c.Execute()
}
