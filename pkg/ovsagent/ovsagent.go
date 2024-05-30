package ovsagent

import (
	"log"
	"github.com/ik8s-ir/beaveragent/pkg/types"
	"strconv"
	"os/exec"
)

func CreateDistrubutedSwitch(bridge string, topology []types.MeshTopology) (string, error) {
	output, stdErr := createOVSBridge(bridge)
	if stdErr !=nil {
		return string(output), stdErr
	}
	for index, item := range topology {
		o,e:=addVXLANtoBridge(bridge,index,item.NodeIP,item.VNI)
		if e!=nil {
			log.Printf("error on vxlan %v",o)
		}
	}
	return string(output), stdErr
}

func createOVSBridge(bridge string) (string, error) {
	output, stdErr := ovsvsctl("add-br", bridge)
	return string(output), stdErr
}

func addVXLANtoBridge(bridge string, vxlanIndex int, remoteIP string, vni int32) (string, error) {
	output, stdErr := ovsvsctl("add-port", bridge, "vxlan"+strconv.Itoa(vxlanIndex), "--", "set", "interface", "vxlan"+strconv.Itoa(vxlanIndex), "type=vxlan", "options:remote_ip="+remoteIP, "options:key="+strconv.Itoa(int(vni)))
	return string(output), stdErr
}

func ovsvsctl(params ...string) ([]byte, error) {
	baseArgs := []string{"--db=unix:/host/var/run/openvswitch/db.sock"}
	args := append(baseArgs, params...)

	command := exec.Command("ovs-vsctl",args...)
	log.Printf("Running: %s \n",command.String())

	return command.CombinedOutput()
}

func DeleteDistrubutedSwitch(bridge string) (string, error) {
	o,e:=ovsvsctl("del-br","bridge")
	return string(o),e
}