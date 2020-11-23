package docker

/*
This file should contain all func related to docker directly and only.
Parameters line building is done by dockerbuilder.
*/
import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"tellaw.org/jsdock/model"
)

// IsProfileRunning is used to check with docker that a profile is already running
func IsProfileRunning(profileAlias string) bool {
	log.Println("Looking for running profile : docker ps")
	cmd := exec.Command("docker", "ps")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined err:\n%s\n", string(out))
		log.Fatalf("Docker is not running !")

	}
	if strings.Contains(string(out), profileAlias) {
		return true
	}

	return false
}

// IsProfileStopped is used to check if the profile is stopped
func IsProfileStopped(profileAlias string) bool {
	log.Println("Looking for stopped profile : docker ps -a")
	cmd := exec.Command("docker", "ps", "-a")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined err:\n%s\n", string(out))
		log.Fatalf("Docker is not running !")

	}

	if strings.Contains(string(out), profileAlias) {
		return true
	}

	return false
}

// StopProfile is used to stop a profile
func StopProfile(profileData model.Profile) {
	command := buildCommand(profileData, "stop")
	dockerStopOrDown(command)
}

// RemoveProfile is used to stop a profile
func RemoveProfile(profileData model.Profile) {
	command := buildCommand(profileData, "remove")
	dockerStopOrDown(command)
}

// StartProfile is the method used to run a profile
func StartProfile(profileData model.Profile) {

	command := buildCommand(profileData, "start")
	dockerRun(command)

}

// dockerExec method used to start a docker command
func dockerStopOrDown(command []string) {

	cmd := exec.Command("docker", command...)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	_ = cmd.Run()

}

// dockerExec method used to start a docker command
func dockerRun(command []string) {

	log.Println("Command is : docker ", command)

	cmd := exec.Command("docker", command...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Println("cmd.Run() failed with \n", err)
	}
}
