package main

import (
	"fmt"
	"log"
	"strings"

	"tellaw.org/jsdock/services"
)

// Compiler : go build -ldflags "-X main.version='0.1'" -o bin/jsdock
var version string

// Basic command : jsdock start in the directory
func main() {

	services.CheckAndInitNetwork()

	pathParam := services.GetPathParam()
	actionName := services.GetAction()
	profileName := services.GetProfile()

	log.Println("JSDOCK version ", version)
	log.Println("Action : ", actionName)
	log.Println("Sources : ", pathParam)
	log.Println("Profile : ", profileName)
	log.Println("------------------------")

	switch strings.ToLower(actionName) {
	case "attach":
		services.Attach()

	case "start":
		services.Start(profileName)

	case "stop":
		services.Stop(profileName)

	case "connect":
		services.Connect(profileName)

	case "version":
		fmt.Println(version)

	default:
		// Default action to define
		// No command, so we have to find the correct command
		services.Start(profileName)
	}

}
