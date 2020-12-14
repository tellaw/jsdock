package docker

import (
	"log"
	"strings"

	"tellaw.org/jsdock/config"
	"tellaw.org/jsdock/model"
)

/*
This file should contain all func related to the building of command line parameters
for the docker command line
*/
func buildCommand(profileData model.Profile, action string) []string {

	commandLine := []string{""}

	switch strings.ToLower(action) {

	case "start":
		// Start action
		commandLine[0] = "run"
		commandLine = append(commandLine, getName(profileData))
		commandLine = append(commandLine, "-d")
		commandLine = append(commandLine, getSourcesVolume(profileData))
		commandLine = getPorts(commandLine, profileData)
		commandLine = getEnv(commandLine, profileData)
		commandLine = getVolumes(commandLine, profileData)
		commandLine = append(commandLine, getImage(profileData))

	case "stop":
		commandLine[0] = "stop"
		commandLine = append(commandLine, GetAlias(profileData))

	case "remove":
		commandLine[0] = "rm"
		commandLine = append(commandLine, GetAlias(profileData))

	case "connect":
		commandLine[0] = "exec"
		commandLine = append(commandLine, "-it")
		commandLine = append(commandLine, GetAlias(profileData))
		commandLine = append(commandLine, "/bin/bash")

	default:
		// Default action to define
	}

	//return commandLine
	return commandLine

}

func getImage(profileData model.Profile) string {

	if profileData.Image == "" {
		log.Fatal("Image name in the profile is not readable")
	}

	return profileData.Image
}

func getName(profileData model.Profile) string {
	return "--name=" + GetAlias(profileData)
}

// GetAlias is used to find the container alias that should be used by the application
func GetAlias(profileData model.Profile) string {
	return "jsdock-" + profileData.Alias
}

func getSourcesVolume(profileData model.Profile) string {
	host := config.GetCurrentPath()
	container := profileData.Sources

	return "-v=" + host + ":" + container
}
func getVolumes(commandLine []string, profileData model.Profile) []string {

	for _, volume := range profileData.Volumes {
		commandLine = append(commandLine, "-v="+volume.Host+":"+volume.Container)
	}
	return commandLine
}

func getEnv(commandLine []string, profileData model.Profile) []string {

	for keyItem, valueItem := range profileData.Env {
		commandLine = append(commandLine, "-e="+keyItem+":"+valueItem)
	}

	return commandLine
}

func getPorts(commandLine []string, profileData model.Profile) []string {

	for _, port := range profileData.Ports {
		commandLine = append(commandLine, "-p="+port.Host+":"+port.Container)
	}

	return commandLine
}
