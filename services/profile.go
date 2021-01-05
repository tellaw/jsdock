package services

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"tellaw.org/jsdock/config"
)

// getProfilesPath Method used to return the profiles files paths
func getProfilesPath() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/jsdock/"

}

// GetProfileLocation Config path and config dir name
func GetProfileLocation() string {
	return getProfilesPath() + "/" + config.ProfilesDirName + "/"
}

// HasProfileFile Method used to know if a config file exists
func HasProfileFile(fileName string) bool {

	jsonfullFileName := getProfilesPath() + fileName + ".json"
	yamlfullFileName := getProfilesPath() + fileName + ".yaml"

	status := false

	if _, err := os.Stat(jsonfullFileName); err == nil {
		status = true

	} else if _, err := os.Stat(yamlfullFileName); err == nil {
		status = true
	}

	return status
}

/*
GetProfileList func used to list profiles in config dir
Return a []string
*/
func GetProfileList() []string {

	var profiles []string

	profileDir := getProfilesPath()

	err := filepath.Walk(profileDir, func(path string, info os.FileInfo, err error) error {
		if path != profileDir {
			if strings.Contains(path, ".json") {
				path = strings.Replace(path, ".json", "", -1)
				profiles = append(profiles, strings.Replace(path, profileDir, "", -1))
			} else if strings.Contains(path, ".yaml") {
				path = strings.Replace(path, ".yaml", "", -1)
				profiles = append(profiles, strings.Replace(path, profileDir, "", -1))
			}
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return profiles

}
