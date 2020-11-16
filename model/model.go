package model

/*
This file should contain all func related to the data model.
*/

// Profile is the struct representing the data of the profile json file
type Profile struct {
	Alias    string
	Location string
	Image    string
	Env      map[string]string
	Ports    map[string]string
	Volumes  map[string]string
	Network  map[string]string
}

// Config is the struct describing the configuration file
type Config struct {
}
