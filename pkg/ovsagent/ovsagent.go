package ovsagent

import (
	"log"
	"os/exec"
)

func CreateDistrubutedSwitch(bridge string) (string, error) {
	output, stdErr := ovsvsctl("add-br", bridge)
	if stdErr != nil {
		log.Printf("Error: %v, %v", stdErr,output)
		return string(output), stdErr
	}
	return string(output), nil
}

func ovsvsctl(params ...string) ([]byte, error) {
	baseArgs := []string{"--db=unix:/host/var/run/openvswitch/db.sock"}
	args := append(baseArgs, params...)

	command := exec.Command("ovs-vsctl",args...)
	log.Printf("Running: %s \n",command.String())

	return command.CombinedOutput()
}
