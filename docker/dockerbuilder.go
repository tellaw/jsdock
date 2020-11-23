package docker

import (
	"strings"

	"tellaw.org/jsdock/model"
)

/*
This file should contain all func related to the building of command line parameters
for the docker command line
*/
func buildCommand(profileData model.Profile, action string) string {

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
	return " "

}

// buildParameters is used to build the command line parameters
func buildParameters(profileData model.Profile, commandLine []string) []string {

	return commandLine

}
