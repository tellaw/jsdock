package model

/*
This file should contain all func related to the data model.
*/

// Profile is the struct representing the data of the profile json file
type Profile struct {

	// https://yourbasic.org/golang/json-example/
	Alias   string `json:"alias"`
	Source  string `json:"sources"`
	Image   string `json:"image"`
	Env     map[string]string
	Ports   map[string]string
	Volumes map[string]string
	Network map[string]string
}

// Config is the struct describing the configuration file
type Config struct {
}
