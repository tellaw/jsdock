package main

import (
	"fmt"
	"os"
	"strings"

	"tellaw.org/jsdock/config"
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

	// Argument 1 is the action
	actionName := os.Args[1]
	profileName := os.Args[2]

	fmt.Println("Action : ", actionName)
	fmt.Println("Profile Name : ", profileName)

	fmt.Println("Path actuel : ", config.GetCurrentPath())

	//fmt.Println("Path to Config dir : ", config.GetConfigDir())

	switch strings.ToLower(actionName) {
	case "attach":
		services.Attach()

	case "start":
		services.Start(profileName)

	default:
		// Default action to define
	}

}
