package services

import (
	"os"
	"strings"
)

// GetPathParam is a method used to find the source path in parameters
func GetPathParam() string {

	if os.Args[1] != "" {

		if strings.Contains(os.Args[1], "/") {
			return os.Args[1]
		}

	} else if os.Args[2] != "" {

		if strings.Contains(os.Args[2], "/") {
			return os.Args[2]
		}

	} else if os.Args[3] != "" {
		if strings.Contains(os.Args[3], "/") {
			return os.Args[3]
		}
	}
	path, _ := os.Getwd()
	return path
}

// GetAction is the method used to find the correct action to apply
func GetAction() string {

	actionVerbs := []string{"start", "stop", "attach", "", ""}

	if os.Args[1] != "" {

		if contains(actionVerbs, os.Args[1]) {
			return os.Args[1]
		}

	} else if os.Args[2] != "" {

		if contains(actionVerbs, os.Args[2]) {
			return os.Args[2]
		}

	} else if os.Args[3] != "" {
		if contains(actionVerbs, os.Args[3]) {
			return os.Args[3]
		}
	}

	// No Verb

	return ""
}

// GetProfile Used to find the profile to use
func GetProfile() string {

	if HasProfileFile(os.Args[1]) {
		return os.Args[1]
	} else if HasProfileFile(os.Args[2]) {
		return os.Args[2]
	} else if HasProfileFile(os.Args[3]) {
		return os.Args[3]
	}

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
