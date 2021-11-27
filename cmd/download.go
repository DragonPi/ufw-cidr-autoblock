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

// DownloadGithubIP fetches the CIDR zones from the GitHub API, meta endpoint
func DownloadGitHubIP(zones []string, verbose bool) (cidr []string, err error) {

	//curl -H "Accept: application/vnd.github.v3+json" https://api.github.com/meta
	//curl -H "Accept: application/vnd.github.v3+json" https://api.github.com/meta | jq .hooks[] | sed 's/"//g' > $PATH_TO_LISTS/.ex.zone

	return
}

// saveZones saves the downloaded zones files locally
func saveZones() (err error) {

	return
}

// appendZones adds CIDR zones to the "automatic_entries" sections of the specified
func appendZones(jsonFile string, section string, subsection string) (err error) {

	return
}
