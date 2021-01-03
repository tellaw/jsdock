package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

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

	fmt.Println("")
	fmt.Println("JSDOCK (version ", version, ")")
	fmt.Println("------------------------")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	fmt.Fprintln(w, "Requested action\t"+actionName)

	fmt.Fprintln(w, "Sources directory\t"+pathParam)
	if profileName == "" {
		fmt.Fprintln(w, "Selected profile\tundefined")
	} else {
		fmt.Fprintln(w, "Selected profile\t"+profileName)
	}
	w.Flush()
	fmt.Println("------------------------")

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

	case "help":
		services.Help()

	case "list":
		services.List()

	case "?":
		services.Help()

	default:
		// Default action to define
		// No command, so we have to find the correct command
		services.Start(profileName)
	}

}
