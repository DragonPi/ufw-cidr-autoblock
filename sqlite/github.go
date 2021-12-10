package sqlite

import (
	"strings"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	u "github.com/DragonPi/ufw-cidr-autoblock/utils"
)

// CacheGitHub caches the info from the GitHub meta endpoint into sqlite database
func CacheAllowedGitHub(metaData *u.GitHub) (err error) {
	var (
		dbc *sql.DB
	)

	if dbc, err = openSQLite(); err != nil {
		return err
	}
	// Close database after function completed
	defer dbc.Close()

	// Cleanup GitHub cache in sqlite db
	if err = cacheCleanupGitHub(dbc); err != nil {
		return err
	}

	allowedGitHub := strings.Split(viper.GetString("exclusions.GitHub"), ",")
	for _, reference := range allowedGitHub {
		switch reference {
		case "hooks":
			for _, zone := range metaData.Hooks {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "web":
			for _, zone := range metaData.Web {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "api":
			for _, zone := range metaData.API {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "git":
			for _, zone := range metaData.Git {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "packages":
			for _, zone := range metaData.Packages {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "pages":
			for _, zone := range metaData.Pages {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "importer":
			for _, zone := range metaData.Importer {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "actions":
			for _, zone := range metaData.Actions {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		case "dependabot":
			for _, zone := range metaData.Dependabot {
				if err = allowZone(zone, "GitHub-"+reference, "No", dbc); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// cacheCleanupGithub cleans the GitHub cache in sqlite database
// otherwise old GitHub zones could remain present in cache.
func cacheCleanupGitHub(dbc *sql.DB) (err error) {
	query := "DELETE FROM allowedzones WHERE reference like 'GitHub-%';"
	if err = doQuery(query, dbc); err != nil {
		return err
	}

	return nil
}
