package jsonparser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"tellaw.org/jsdock/model"
	"tellaw.org/jsdock/services"
)

/*
Package used to write and read JSon
*/

// LoadJSON method used to load JSon file
func LoadJSON(fileName string) {

	fmt.Println(services.GetProfileLocation() + fileName + ".json")
	file, _ := ioutil.ReadFile(services.GetProfileLocation() + fileName + ".json")

	var data model.Profile

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data.Alias)

}
