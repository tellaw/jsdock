package services

import (
	"fmt"
	"io/ioutil"

	"tellaw.org/jsdock/docker"
	"tellaw.org/jsdock/jsonparser"
	"tellaw.org/jsdock/prompt"
)

// Attach is used to configure a directory to work with this app
func Attach() {

	// Grab available profiles
	profiles := GetProfileList()

	// Prompt for the profile name
	result := prompt.InList(profiles, "Select a Profile")

	// Create json file at root containing the name of the profile
	content := "{\"profile\":\"" + result + "\"}"

	d1 := []byte(content)
	err := ioutil.WriteFile(".jsdock", d1, 0644)

	if err != nil {
		panic(err)
	}

}

// Start is the method used to start a docker kit
func Start(profileName string) {

	/*
		Look if profile is running
			stop if running with down command
			start it then
	*/
	if !docker.IsProfileRunning(profileName) {
		fmt.Println("Profile [" + profileName + "]is not running")
	} else {
		fmt.Println("Profile [" + profileName + "] is running")
	}

	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		panic("The requested profile doesn't exists")
	}

	profileData := jsonparser.LoadProfileJSON(GetProfileLocation(), profileName)

	docker.StartProfile(profileData)

}
