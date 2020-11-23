package docker

/*
This file should contain all func related to docker directly and only.
Parameters line building is done by dockerbuilder.
*/
import (
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
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	//fmt.Printf("combined out:\n%s\n", string(out))

	if !isDockerWorking(string(out)) {
		panic("Docker is not running")
	}

	if strings.Contains(string(out), "jsdock-"+profileName) {
		return true
	}

	return false
}

func isDockerWorking(response string) bool {
	if strings.Contains(response, "CONTAINER ID") {
		return true
	}
	return false
}

// StartProfile is the method used to run a profile
func StartProfile(profileData model.Profile) {

	command := buildCommand(profileData, "start")

	dockerExec(command)

}

// dockerExec method used to start a docker command
func dockerExec(command [10]string) {
	cmd := exec.Command(command[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
