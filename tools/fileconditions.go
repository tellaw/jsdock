package tools

import (
	"io/ioutil"
	"os"
	"strings"
)

var (
	fileInfo *os.FileInfo
	err      error
)

// Find the full filename of the path (if relative, then add the current dir as prefix)
func getFullFilename(fileName string) string {

	if strings.HasPrefix(string(fileName[0]), "/") {
		path, _ := os.Getwd()
		fileName = path + fileName
	}

	return fileName

}

// FileExists method used to valid if a file exists
func FileExists(fileName string) bool {

	fileName = getFullFilename(fileName)

	info, err := os.Stat(fileName)

	// File or directory doesn't exists
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}

// DirExists method is used to check if dir exists
func DirExists(fileName string) bool {

	fileName = getFullFilename(fileName)

	info, err := os.Stat(fileName)

	// File or directory doesn't exists
	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() == false {
		return false
	}

	return true

}

// FileContains check if a file contains a string
func FileContains(fileName string, value string) bool {

	fileName = getFullFilename(fileName)

	if FileExists(fileName) == false {
		return false
	}

	// read the whole file at once
	b, _ := ioutil.ReadFile(fileName)

	s := string(b)

	// //check whether s contains substring text
	if strings.Contains(s, value) == false {
		return false
	}

	return true
}
