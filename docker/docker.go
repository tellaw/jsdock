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
func IsProfileRunning(profileName string) bool {

	cmd := exec.Command("docker", "ps")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined err:\n%s\n", string(out))
		log.Fatalf("Docker is not running !")

	}

	if strings.Contains(string(out), "jsdock-"+profileName) {
		return true
	}

	return false
}

// IsProfileStopped is used to check if the profile is stopped
func IsProfileStopped(profileName string) bool {
	cmd := exec.Command("docker", "ps", "-a")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined err:\n%s\n", string(out))
		log.Fatalf("Docker is not running !")

	}

	if strings.Contains(string(out), "jsdock-"+profileName) {
		return true
	}

	return false
}

// StartProfile is the method used to run a profile
func StartProfile(profileData model.Profile) {

	command := buildCommand(profileData, "start")

	dockerRun(command)

}

// dockerExec method used to start a docker command
func dockerRun(command []string) {

	log.Println("Command is : docker ", command)

	cmd := exec.Command("docker", command...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
