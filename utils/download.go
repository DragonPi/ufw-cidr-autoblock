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
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// GitHub is the base struct containing the data fetched from https://api.github.com/meta
type GitHub struct {
	VPA        bool      `json:"verifiable_password_authentication"`
	SSH_fp     GH_ssh_fp `json:"ssh_key_fingerprints"`
	Hooks      []string  `json:"hooks"`
	Web        []string  `json:"web"`
	API        []string  `json:"api"`
	Git        []string  `json:"git"`
	Packages   []string  `json:"packages"`
	Pages      []string  `json:"pages"`
	Importer   []string  `json:"importer"`
	Actions    []string  `json:"actions"`
	Dependabot []string  `json:"dependabot"`
}

// GH_ssh_fp contains the ssh fingerprint from GitHub meta endpoint
type GH_ssh_fp struct {
	RSA     string `json:"SHA256_RSA"`
	ECDSA   string `json:"SHA256_ECDSA"`
	ED25519 string `json:"SHA256_ED25519"`
}

// DownloadGithubIP fetches the CIDR zones from the GitHub API, meta endpoint
func DownloadGitHubIP(metaData **GitHub) (err error) {
	var (
		req  *http.Request
		res  *http.Response
		body []byte
	)

	url := "https://api.github.com/meta"

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return fmt.Errorf("generating new HTTP Request: %w", err)
	}

	req.Header.Set("User-Agent", "Download Github meta-data")

	if res, err = client.Do(req); err != nil {
		return fmt.Errorf("doing HTTP Request: %w", err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return fmt.Errorf("reading HTTP body: %w", err)
	}

	if err := json.Unmarshal(body, &metaData); err != nil {
		return fmt.Errorf("unmarchalling json string: %w", err)
	}

	return
}

// DownloadZoneFiles fetches the CIDR zones files from ipverse.net
//
// It uses concurrent connections to speed up the download
func DownloadZoneFiles(countryZones []string) (err error) {
	// Configure the number of concurrent connections to download the zone files
	concurrency := defineConcurrency()

	// Make downloads concurrently
	sem := make(chan bool, concurrency)

	for _, country := range countryZones {
		sem <- true

		go func(country string) {
			defer func() { <-sem }()

			if err = doZoneDownload(strings.ToLower(country)); err != nil {
				// If a file fails to download, I want to continue
				// with the others but send a warning nevertheless
				Warning.Println(err)
			}
		}(country)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	return
}

// doZoneDownload does the actual download of the zone file and saves
// it in the location set in the .uca-config.ini file
func doZoneDownload(country string) (err error) {
	var (
		req *http.Request
		res *http.Response
		out *os.File
	)

	baseURL := "http://ipverse.net/ipblocks/data/countries/"
	zone := country + ".zone"
	URL := baseURL + zone

	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	if req, err = http.NewRequest(http.MethodGet, URL, nil); err != nil {
		return fmt.Errorf("generating new HTTP Request: %w", err)
	}

	req.Header.Set("User-Agent", "Download zone files")

	if res, err = client.Do(req); err != nil {
		return fmt.Errorf("doing HTTP Request: %w", err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	zoneFile := filepath.Join(viper.GetString("zones.zonesLocation"), zone)
	MakeDestination(zoneFile)

	if out, err = os.Create(zoneFile); err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, res.Body)

	return
}
