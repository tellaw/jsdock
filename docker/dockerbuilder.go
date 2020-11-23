package docker

import (
	"strings"

	"tellaw.org/jsdock/model"
)

/*
This file should contain all func related to the building of command line parameters
for the docker command line
*/
func buildCommand(profileData model.Profile, action string) [10]string {

	var commandLine [10]string

	commandLine[0] = "docker "

	switch strings.ToLower(action) {

	case "start":
		// Start action
		commandLine[1] = " start "
	default:
		// Default action to define
	}

	//return commandLine
	return commandLine

}

// buildParameters is used to build the command line parameters
func buildParameters(profileData model.Profile, commandLine []string) []string {

	commandLine[2] = getImage(profileData)

	return commandLine

}

func getImage(profileData model.Profile) string {
	return profileData.Image
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
