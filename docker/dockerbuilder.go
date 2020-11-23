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
	commandLine []string
	
	 "docker "

	switch strings.ToLower(action) {

	case "start":
		// Start action
		commandLine += " start "
	default:
		// Default action to define
	}

	return commandLine

}
