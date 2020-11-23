package model

/*
This file should contain all func related to the data model.
*/

// Profile is the struct representing the data of the profile json file describing
type Profile struct {

	// https://yourbasic.org/golang/json-example/
	Alias   string            `json:"alias"`
	Sources string            `json:"sources"`
	Image   string            `json:"image"`
	Env     map[string]string `json:"env"`
	Ports   []Port            `json:"ports"`
	Volumes []Volume          `json:"volumes"`
	Network map[string]string
}

// Volume is the struct object for docker vulomes mapping
type Volume struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

// Port is the model for ports to listen to
type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

// Config is the struct describing the configuration file
// Located in the directory
type Config struct {
	Profile string `json:"profile"`
}
