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
#run openvswitch daemon
ovs-vswitchd \
   --unixctl=/var/run/openvswitch/ovs-vswitchd.ctl \
   --pidfile=/var/run/openvswitch/ovs-vswitchd.pid \
   --overwrite-pidfile \
   --detach
# create default public external internet bridge
ovs-vsctl add-br $DEFAULT_OVS_BRIDGE 2> /dev/null
ovs-vsctl add-port $DEFAULT_OVS_BRIDGE $NODE_NIC 2> /dev/null

# Isolataion for default public external network
ovs-ofctl del-flows $DEFAULT_OVS_BRIDGE
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "priority=0,actions=drop"
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "in_port=$NODE_NIC,actions=NORMAL"
ovs-ofctl add-flow $DEFAULT_OVS_BRIDGE "actions=output=$NODE_NIC"

# run the agent
./beaveragent