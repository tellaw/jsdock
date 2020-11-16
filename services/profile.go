package services

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"tellaw.org/jsdock/config"
	"tellaw.org/jsdock/model"
)

// getProfilesPath Method used to return the profiles files paths
func getProfilesPath() string {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir

}

// GetProfileLocation Config path and config dir name
func GetProfileLocation() string {
	return getProfilesPath() + "/" + config.ProfilesDirName + "/"
}

// HasProfileFile Method used to know if a config file exists
func HasProfileFile(fileName string) bool {

	fullFileName := ""
	fullFileName = GetProfileLocation() + fileName

	_, err := os.Stat(fullFileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
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
			path = strings.Replace(path, ".json", "", -1)
			profiles = append(profiles, strings.Replace(path, profileDir, "", -1))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, profile := range profiles {
		fmt.Println(profile)
	}

	return profiles

}

// LoadProfile Loading profile
func LoadProfile(profileName string) model.Profile {

	var profile model.Profile
	profile.Alias = "test"

	return profile

}
