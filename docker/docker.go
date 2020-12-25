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

	//log.Println("Looking for running profile : docker ps")

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
	//log.Println("Looking for stopped profile : docker ps -a")
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

// IsNetworkAvailable method used to check if network is set
func IsNetworkAvailable() bool {
	//log.Println("Looking for stopped profile : docker ps -a")
	cmd := exec.Command("docker", "network", "ls")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined err:\n%s\n", string(out))
		log.Fatalf("Docker is not running !")

	}

	if strings.Contains(string(out), "jsdock_net") {
		return true
	}

	return false
}

// InitNetwork method is used to setup the dev network
func InitNetwork() {
	commandLine := []string{""}
	commandLine[0] = "docker network create jsdock_net"
	dockerRun(commandLine)
}

// StopProfile is used to stop a profile
func StopProfile(profileData model.Profile) {
	command := buildCommand(profileData, "stop")
	dockerStopOrDown(command)
}

// StopDockerProcess method stop without control a docker process. Please prefer the StopProfile method.
func StopDockerProcess(processName string) {

	// Stop

	// Find process ID
	cmd := exec.Command("docker", "container", "ls", "-q", "--filter", "name="+processName)
	out, _ := cmd.CombinedOutput()

	fmt.Println("Process ID is : ", strings.TrimSpace(string(out)))

	// docker container stop $(docker container ls -q --filter name=myapp*)
	//cmd := exec.Command("docker", "stop", processName)
	cmdstop := exec.Command("docker", "stop", strings.TrimSpace(string(out)))
	outstop, _ := cmdstop.CombinedOutput()

	fmt.Println("Stopping container : ", strings.TrimSpace(string(outstop)))

	// Destroy
	cmd = exec.Command("docker", "rm", strings.TrimSpace(string(out)))
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	outstop, _ = cmd.CombinedOutput()

	fmt.Println("Removing container : ", strings.TrimSpace(string(outstop)))

}

// RemoveProfile is used to stop a profile
func RemoveProfile(profileData model.Profile) {
	command := buildCommand(profileData, "remove")
	dockerStopOrDown(command)
}

// Connect is used to connect to container
func Connect(profileData model.Profile) {
	command := buildCommand(profileData, "connect")
	fmt.Println("docker", command)
	//dockerRun(command)
}

// StartProfile is the method used to run a profile
func StartProfile(profileData model.Profile) {
	command := buildCommand(profileData, "start")
	dockerRun(command)
}

// dockerExec method used to start a docker command
func dockerStopOrDown(command []string) {

	cmd := exec.Command("bash", "-c", strings.Join(command, " "))
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	_ = cmd.Run()

}

// dockerExec method used to start a docker command
func dockerRun(command []string) {

	log.Println("Command is : ", command)

	cmd := exec.Command("bash", "-c", strings.Join(command, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Println("cmd.Run() failed with \n", err)
	}
}
