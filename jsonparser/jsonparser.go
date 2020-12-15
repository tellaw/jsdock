package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"tellaw.org/jsdock/model"
)

/*
Package used to write and read JSon
*/
// LoadProfileJSON method used to load JSon file
func LoadProfileJSON(profileLocation string, fileName string) model.Profile {

	//fmt.Println(services.GetProfileLocation() + fileName + ".json")
	file, _ := ioutil.ReadFile(profileLocation + fileName + ".json")

	var data model.Profile

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Println(err)
	}

	return data

}

// LoadConfigJSON method used to load the config JSon file
func LoadConfigJSON(configFileLocation string) model.Config {

	file, _ := ioutil.ReadFile(configFileLocation + "/.jsdock")

	var data model.Config

	err := json.Unmarshal([]byte(file), &data)

	//log.Println(".jsdock file contain : ", data.Profile)
	if err != nil {
		log.Println(err)
	}

	return data

}

// HasConfigFile Method used to know if a config file exists
func HasConfigFile(configFileLocation string) bool {

	fullFileName := ""
	fullFileName = configFileLocation + "/.jsdock"

	_, err := os.Stat(fullFileName)

	if os.IsNotExist(err) {
		//log.Println("Unable to find profile in : " + fullFileName)
		return false
	}
	return true
}
