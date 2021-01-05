package yamlparser

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	"tellaw.org/jsdock/model"
)

// LoadProfileYAML method used to load JSon file
func LoadProfileYAML(profileLocation string, fileName string) model.Profile {

	// https: //github.com/go-yaml/yaml
	file, _ := ioutil.ReadFile(profileLocation + fileName + ".yaml")

	var data model.YAMLProfile

	err := yaml.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Println(err)
	}

	out := convertYAMLToProfile(data)

	return out

}

func convertYAMLToProfile(data model.YAMLProfile) model.Profile {

	var out model.Profile

	out.Alias = data.Alias
	out.Sources = data.Sources
	out.Image = data.Image
	out.Env = data.Env
	out.Ports = data.Ports
	out.Volumes = data.Volumes
	out.Network = data.Network

	return out

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
