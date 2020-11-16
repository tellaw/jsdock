package docker

/*
This file should contain all func related to docker directly and only.
Parameters line building is done by dockerbuilder.
*/
import (
	"log"
	"os/exec"
	"strings"
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
