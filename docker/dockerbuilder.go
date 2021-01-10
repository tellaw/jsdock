package docker

import (
	"fmt"
	"log"
	"strings"

	"tellaw.org/jsdock/model"
	"tellaw.org/jsdock/tools"
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
		commandLine[0] = "docker run"
		commandLine = append(commandLine, "-d")
		commandLine = append(commandLine, getName(profileData))
		if sourceVolume := getSourcesVolume(profileData); sourceVolume != "" {
			commandLine = append(commandLine, getSourcesVolume(profileData))
		}
		commandLine = getPorts(commandLine, profileData)
		commandLine = getEnv(commandLine, profileData)
		commandLine = getVolumes(commandLine, profileData)
		commandLine = append(commandLine, "--network jsdock_net")

		commandLine = append(commandLine, getImage(profileData))

	case "stop":
		commandLine[0] = "docker stop"
		commandLine = append(commandLine, GetAlias(profileData))

	case "remove":
		commandLine[0] = "docker rm"
		commandLine = append(commandLine, GetAlias(profileData))

	case "connect":
		commandLine[0] = "docker exec"
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
	return "--name " + GetAlias(profileData)
}

// GetAlias is used to find the container alias that should be used by the application
func GetAlias(profileData model.Profile) string {
	return profileData.Alias
}

func getSourcesVolume(profileData model.Profile) string {

	if profileData.Sources != "" {

		host := profileData.PathParam
		container := profileData.Sources

		return "-v " + host + ":" + container

	} else {
		return ""
	}

}

func getVolumes(commandLine []string, profileData model.Profile) []string {

	for _, volume := range profileData.Volumes {

		var conditionChecked bool = true

		// Check if Volume has conditions
		for _, fileExist := range volume.Conditions.FileExists {
			if tools.FileExists(fileExist) == false {
				fmt.Println("File exist condition is not true", fileExist)
				conditionChecked = false
			}
		}

		// Check if Volume has conditions
		for _, dirExist := range volume.Conditions.DirExists {
			if tools.DirExists(dirExist) == false {
				fmt.Println("Directory exist condition is not true", dirExist)
				conditionChecked = false
			}
		}

		// Check if Volume has conditions
		for _, fileContain := range volume.Conditions.FileContains {
			if tools.FileContains(fileContain.File, fileContain.Value) == false {
				fmt.Println("File contain condition is not true", fileContain)
				conditionChecked = false
			}
		}

		if conditionChecked == true {
			commandLine = append(commandLine, "-v "+volume.Host+":"+volume.Container)
		}
	}
	return commandLine
}

func getEnv(commandLine []string, profileData model.Profile) []string {

	for keyItem, valueItem := range profileData.Env {
		commandLine = append(commandLine, "-e "+keyItem+"='"+valueItem+"'")
	}

	return commandLine
}

func getPorts(commandLine []string, profileData model.Profile) []string {

	for _, port := range profileData.Ports {
		commandLine = append(commandLine, "-p "+port.Host+":"+port.Container)
	}

	return commandLine
}
