package model

/*
This file should contain all func related to the data model.
*/

// Profile is the struct representing the data of the profile json file describing
type Profile struct {

	// https://yourbasic.org/golang/json-example/
	Alias     string            `json:"alias"`
	Sources   string            `json:"sources"`
	Image     string            `json:"image"`
	Env       map[string]string `json:"env"`
	Ports     []Port            `json:"ports"`
	Volumes   []Volume          `json:"volumes"`
	PathParam string            // This is for internal usage
	Network   map[string]string
}

// YAMLProfile is the struct representing the data of the profile json file describing
type YAMLProfile struct {

	// https://yourbasic.org/golang/json-example/
	Alias   string            `yaml:"alias"`
	Sources string            `yaml:"sources"`
	Image   string            `yaml:"image"`
	Env     map[string]string `yaml:"env"`
	Ports   []Port            `yaml:"ports"`
	Volumes []Volume          `yaml:"volumes"`
	Network map[string]string
}

// Conditions is a wrapper for the volume inject conditions
type Conditions struct {
	FileExists   []string               `json:"fileExists"`
	DirExists    []string               `json:"dirExists"`
	FileContains []FileContainCondition `json:"fileContains"`
}

// FileContainCondition represent a contant of a file to test
type FileContainCondition struct {
	File  string `json:"file"`
	Value string `json:"value"`
}

// Volume is the struct object for docker vulomes mapping
type Volume struct {
	Host       string     `json:"host"`
	Container  string     `json:"container"`
	Conditions Conditions `json:"conditions"`
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
