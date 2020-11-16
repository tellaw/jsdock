package config

import (
	"os"
)

// ProfilesDirName is the name of the directory where profile are stored
var ProfilesDirName string = "jsdock"

// GetCurrentPath return the path where the program is executed
func GetCurrentPath() string {
	path, _ := os.Getwd()
	return path
}
