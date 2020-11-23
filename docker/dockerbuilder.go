package docker

import (
	"strings"

	"tellaw.org/jsdock/model"
)

/*
This file should contain all func related to the building of command line parameters
for the docker command line
*/
func buildCommand(profileData model.Profile, action string) []string {

	commandLine := []string{"", "", "", ""}

	switch strings.ToLower(action) {

	case "start":
		// Start action
		commandLine[0] = "run"
	default:
		// Default action to define
	}

	commandLine = buildParameters(profileData, commandLine)

	//return commandLine
	return commandLine

}

// buildParameters is used to build the command line parameters
func buildParameters(profileData model.Profile, commandLine []string) []string {

	commandLine[3] = getImage(profileData)
	commandLine[2] = "-d"
	commandLine[1] = getName(profileData)

	return commandLine

}

func getImage(profileData model.Profile) string {
	return profileData.Image
}

func getName(profileData model.Profile) string {
	return "--name=" + "jsdock-" + profileData.Alias
}

func getSourcesVolume(profileData model.Profile) string {
	return ""
}
func getVolumes(profileData model.Profile) string {
	return ""
}

func getEnv(profileData model.Profile) string {
	return ""
}
