package jsonparser

import (
	"encoding/json"
	"io/ioutil"
	"log"

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
func LoadConfigJSON(profileLocation string, fileName string) model.Config {

	//fmt.Println(services.GetProfileLocation() + fileName + ".json")
	file, _ := ioutil.ReadFile(profileLocation + fileName + ".json")

	var data model.Config

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Println(err)
	}

	return data

}
