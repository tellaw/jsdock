package main

import (
	"log"
	"strings"

	"tellaw.org/jsdock/services"
)

// Basic command : jsdock start in the directory
/*
Command line parameters :

No parameter = Valid
Action
Action & path

*/
func main() {

	pathParam := services.GetPathParam()
	actionName := services.GetAction()
	profileName := services.GetProfile()

	log.Println("************************************")
	log.Println("Action to apply : ", actionName)
	log.Println("Sources Path : ", pathParam)
	log.Println("Profile Name : ", profileName)
	log.Println("************************************")

	//fmt.Println("Path to Config dir : ", config.GetConfigDir())

	switch strings.ToLower(actionName) {
	case "attach":
		services.Attach()

	case "start":
		services.Start(profileName)

	default:
		// Default action to define
		// No command, so we have to find the correct command
	}

}
