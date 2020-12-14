package services

import (
	"os"
	"strings"

	"tellaw.org/jsdock/jsonparser"
)

// GetPathParam is a method used to find the source path in parameters
func GetPathParam() string {

	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {
		if strings.Contains(os.Args[1], "/") {
			return os.Args[1]
		}

	}
	if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" {
		if strings.Contains(os.Args[2], "/") {
			return os.Args[2]
		}

	}
	if len(os.Args) > 3 && strings.TrimSpace(os.Args[3]) != "" {
		if strings.Contains(os.Args[3], "/") {
			return os.Args[3]
		}
	}
	path, _ := os.Getwd()
	return path
}

// GetAction is the method used to find the correct action to apply
func GetAction() string {

	actionVerbs := []string{"start", "stop", "attach", "connect", ""}

	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {

		if contains(actionVerbs, os.Args[1]) {
			return os.Args[1]
		}

	}
	if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" {

		if contains(actionVerbs, os.Args[2]) {
			return os.Args[2]
		}

	}
	if len(os.Args) > 3 && strings.TrimSpace(os.Args[3]) != "" {
		if contains(actionVerbs, os.Args[3]) {
			return os.Args[3]
		}
	}

	// No Verb
	return "start"

}

// GetProfile Used to find the profile to use
func GetProfile() string {

	if len(os.Args) > 1 && strings.TrimSpace(os.Args[1]) != "" {
		if HasProfileFile(os.Args[1]) {
			return os.Args[1]
		}
	}

	if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" {
		if HasProfileFile(os.Args[2]) {
			return os.Args[2]
		}
	}

	if len(os.Args) > 3 && strings.TrimSpace(os.Args[3]) != "" {
		if HasProfileFile(os.Args[3]) {
			return os.Args[3]
		}
	}

	// Check if there is a config of attached directory
	path, _ := os.Getwd()
	if jsonparser.HasConfigFile(path) {

		configData := jsonparser.LoadConfigJSON(path)
		return configData.Profile
	}

	// At the end, data empty returned
	return ""
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
