package services

import (
	"io/ioutil"
	"log"

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

	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	// Check if a profile with same alias is already runnin
	if docker.IsProfileRunning(profileName) {
		log.Println("Profile [" + profileName + "] is running")
		// Stop it
	}

	if docker.IsProfileStopped(profileName) {
		// Then remove the profile using a simple docker rm
	}

	profileData := jsonparser.LoadProfileJSON(GetProfileLocation(), profileName)

	docker.StartProfile(profileData)

}
