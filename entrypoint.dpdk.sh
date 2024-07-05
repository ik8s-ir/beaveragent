#!/bin/sh
# create ovs database if not exist.
ovsdb-tool create 2> /dev/null
# run ovsdb server
ovsdb-server \
   --unixctl=/var/run/openvswitch/ovsdb-server.ctl \
   --pidfile=/var/run/openvswitch/ovsdb-server.pid \
   --remote=punix:/var/run/openvswitch/db.sock \
   --overwrite-pidfile \
   --detach
# Initialize OVS with DPDK
ovs-vsctl --no-wait set Open_vSwitch . other_config:dpdk-init=true
ovs-vsctl --no-wait set Open_vSwitch . other_config:dpdk-lcore-mask=0x1
ovs-vsctl --no-wait set Open_vSwitch . other_config:dpdk-socket-mem=1024,1024
ovs-vsctl --no-wait set Open_vSwitch . other_config:dpdk-hugepages-dir=/dev/hugepages
#run openvswitch daemon
ovs-vswitchd \
   --dpdk -vconsole:emer -vsyslog:err -vfile:info \
   --unixctl=/var/run/openvswitch/ovs-vswitchd.ctl \
   --pidfile=/var/run/openvswitch/ovs-vswitchd.pid \
   --overwrite-pidfile \
   --detach
# create default public external internet bridge
ovs-vsctl add-br $DEFAULT_OVS_BRIDGE -- set bridge $DEFAULT_OVS_BRIDGE datapath_type=netdev 2> /dev/null
ovs-vsctl add-port $DEFAULT_OVS_BRIDGE $NODE_NIC -- set Interface $NODE_NIC type=dpdk 2> /dev/null

# Isolataion for default public external network
ovs-ofctl del-flows $DEFAULT_OVS_BRIDGE
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "priority=0,actions=drop"
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "in_port=$NODE_NIC,actions=NORMAL"
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "actions=output=$NODE_NIC"

# run the agent
./beaveragent