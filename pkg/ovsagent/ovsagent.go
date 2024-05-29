package ovsagent

import (
	"log"
	"os/exec"
)

func CreateDistrubutedSwitch(bridge string) {
	ovsvsctl("create-br", bridge)
}

func ovsvsctl(params ...string) {
	baseArgs := []string{"--db=unix:/host/var/run/openvswitch/db.sock"}
	args := append(baseArgs, params...)

	command := exec.Command("ovs-vsctl", args...)
	log.Println(command.String())

	// output, err := command.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// 	log.Printf("%s\n", output)
	// 	return
	// }

	// log.Printf("%s\n", output)
}
