package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"tellaw.org/jsdock/docker"
	"tellaw.org/jsdock/jsonparser"
	"tellaw.org/jsdock/prompt"
)

// Attach is used to configure a directory to work with this app
func Attach() string {

	fmt.Println("Select the profile to attach to this directory")

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

	return result

}

// Stop function is used to stop started continers
func Stop(profileName string) {

	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use for container")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := jsonparser.LoadProfileJSON(getProfilesPath(), profileName)
	// Check if a profile with same alias is already runnin
	if docker.IsProfileRunning(docker.GetAlias(profileData)) {
		log.Println("Profile [" + profileName + "] is running, stopping...")
		// Stop it
		docker.StopProfile(profileData)
	} else {
		log.Println("profile is not running")
	}
	//log.Println("profile not running anymore")

}

// Connect Method is used to open a bash shell in the container
func Connect(profileName string) {
	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use for container")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := jsonparser.LoadProfileJSON(getProfilesPath(), profileName)
	docker.Connect(profileData)
}

// Start is the method used to start a docker kit
func Start(profileName string) {

	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use for container")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := jsonparser.LoadProfileJSON(getProfilesPath(), profileName)

	// Check if a profile with same alias is already runnin
	if docker.IsProfileRunning(docker.GetAlias(profileData)) {
		log.Println("Profile [" + profileName + "] is running, stopping...")
		// Stop it
		docker.StopProfile(profileData)
	} else {
		log.Println("profile is not running")
	}
	log.Println("profile not running anymore")

	if docker.IsProfileStopped(docker.GetAlias(profileData)) {
		// Then remove the profile using a simple docker rm
		log.Println("Profile [" + profileName + "] is stopped, removing...")
		docker.RemoveProfile(profileData)
	} else {
		log.Println("profile is not stopped")
	}
	log.Println("profile stopped")
	docker.StartProfile(profileData)

}

// CheckAndInitNetwork method is used to check if jsdock network is available
func CheckAndInitNetwork() {
	if !docker.IsNetworkAvailable() {
		docker.InitNetwork()
	}
}
