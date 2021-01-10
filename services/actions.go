package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"tellaw.org/jsdock/docker"
	"tellaw.org/jsdock/jsonparser"
	"tellaw.org/jsdock/model"
	"tellaw.org/jsdock/prompt"
	"tellaw.org/jsdock/yamlparser"
)

// Attach is used to configure a directory to work with this app
func Attach() string {

	fmt.Println("Select the profile to attach to this directory")

	// Grab available profiles
	profiles := GetProfileList()

	for _, profile := range profiles {
		fmt.Println(profile)
	}

	// Prompt for the profile name
	result := prompt.InList(profiles, "Select a Profile")

	// Create json file at root containing the name of the profile
	content := "{\"profile\":\"" + result + "\"}"

	d1 := []byte(content)
	err := ioutil.WriteFile(".jsdock", d1, 0644)

	if err != nil {
		panic(err)
	}

	return result

}

// Stop function is used to stop started continers
func Stop(profileName string) {

	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use for container")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := getProfileFactory(getProfilesPath(), profileName)
	// Check if a profile with same alias is already runnin
	if docker.IsProfileRunning(docker.GetAlias(profileData)) {
		fmt.Println("Profile [" + profileName + "] is running, stopping...")
		// Stop it
		docker.StopProfile(profileData)
	} else {
		fmt.Println("profile is not running")
	}
	//log.Println("profile not running anymore")

}

// Connect Method is used to open a bash shell in the container
func Connect(profileName string) {
	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use for container")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := getProfileFactory(getProfilesPath(), profileName)
	docker.Connect(profileData)
}

// Start is the method used to start a docker kit
func Start(profileName string, pathParam string) {

	if strings.TrimSpace(profileName) == "" {
		log.Fatal("Missing the name of the profile to use")
	}
	// Check if profile exists and load it
	if !HasProfileFile(profileName) {
		log.Fatal("The requested profile " + profileName + " does not exists")
	}

	profileData := getProfileFactory(getProfilesPath(), profileName)

	// Add Pathparam to profileData model
	profileData.PathParam = pathParam

	CheckAndResolveProfileConflicts(profileData)

	// Check if a profile with same alias is already runnin
	if docker.IsProfileStopped(docker.GetAlias(profileData)) {
		// Then remove the profile using a simple docker rm
		fmt.Println("Profile [" + profileName + "] is stopped, removing...")
		docker.RemoveProfile(profileData)
	} else {
		fmt.Println("profile is not stopped")
	}
	fmt.Println("profile stopped")
	docker.StartProfile(profileData)

}

// CheckAndResolveProfileConflicts is a method used to check if there is any conflict ( profile stated, port in use )
func CheckAndResolveProfileConflicts(profileData model.Profile) bool {

	// Loop over ps to check used ports
	cmd := exec.Command("docker", "ps", "--format", "\"{{.Names}}|{{.Ports}}\"")
	out, _ := cmd.CombinedOutput()

	// Split by line & look for the information
	outputLines := strings.Split(string(out), "\n")
	for _, port := range profileData.Ports {

		requiredPort := port.Container
		for _, line := range outputLines {

			outLineCols := strings.Split(line, "|")
			if len(outLineCols) > 1 {

				processName := outLineCols[0][1:]
				fmt.Println("checking conflict with running container : ", processName, " - ", requiredPort, " - ", outLineCols[1])
				if strings.Contains(outLineCols[1], requiredPort) {
					docker.StopDockerProcess(processName)
				}
			}

		}

	}

	return true

}

// Help method generates the help output
func Help() {

	fmt.Println(" ")
	fmt.Println("Help...")
	fmt.Println(" ")
	fmt.Println("Command line example : jsdock <action|optionnal> <profileName|optionnal> <pathAsSource|optionnal>")
	fmt.Println(" ")
	fmt.Println("Command line attributes are the following :")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	// "start", "stop", "attach", "connect", "version", "?", "help"
	fmt.Fprintln(w, "Action\tDefine the action to execute. See below")
	fmt.Fprintln(w, "Profile\tServer profile to use")
	fmt.Fprintln(w, "Source\tSource directory to use. (default is the current directory)") // trailing tab
	w.Flush()

	fmt.Println(" ")
	fmt.Println("Available actions are :")
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	// "start", "stop", "attach", "connect", "version", "?", "help"
	fmt.Fprintln(w, "start\tAction ued to start a server. (Default action)")
	fmt.Fprintln(w, "stop\tAction used to stop a running server")
	fmt.Fprintln(w, "attach\tSet te default profile for a directory") // trailing tab
	fmt.Fprintln(w, "connect\tOutput command to connect to server")
	fmt.Fprintln(w, "list\tList available profiles")
	fmt.Fprintln(w, "version\tJSDock current version")
	fmt.Fprintln(w, "help\tDisplay this help")
	fmt.Fprintln(w, "?\tDisplay this help")
	w.Flush()
	fmt.Println(" ")
}

// List method will print the detail of existing profiles
func List() {

	profiles := GetProfileList()

	fmt.Println("")
	fmt.Println("Available profiles in the JSon repository (" + getProfilesPath() + "):")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)

	for _, profileName := range profiles {

		profileData := getProfileFactory(getProfilesPath(), profileName)

		fmt.Fprintln(w, profileName+"\t Image => "+profileData.Image+"\t Alias => "+profileData.Alias)

	}
	w.Flush()
	fmt.Println(" ")
}

func getProfileFactory(profileLocation string, fileName string) model.Profile {

	jsonfullFileName := getProfilesPath() + fileName + ".json"
	yamlfullFileName := getProfilesPath() + fileName + ".yaml"

	var profileData model.Profile

	if _, err := os.Stat(jsonfullFileName); err == nil {
		// Unable to find the JSON profile

		profileData = jsonparser.LoadProfileJSON(getProfilesPath(), fileName)

	} else if _, err := os.Stat(yamlfullFileName); err == nil {
		// Unable to also find the yaml profile
		profileData = yamlparser.LoadProfileYAML(getProfilesPath(), fileName)
	}

	return profileData

}

// CheckAndInitNetwork method is used to check if jsdock network is available
func CheckAndInitNetwork() {
	if !docker.IsNetworkAvailable() {
		docker.InitNetwork()
	}
}
