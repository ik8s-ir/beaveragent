#!/bin/sh
 
ovsdb-tool create /host/var/lib/openvswitch/conf.db 2> /dev/null
ovsdb-server \
   --unixctl=/host/var/run/openvswitch/ovsdb-server.0.ctl \
   --pidfile=/host/var/run/openvswitch/ovsdb-server.pid \
   --remote=punix:/host/var/run/openvswitch/db.sock \
   --overwrite-pidfile \
   --detach \
   /host/var/lib/openvswitch/conf.db
 
ovs-vswitchd \
   --unixctl=/host/var/run/openvswitch/ovs-vswitchd.0.ctl \
   --pidfile=/host/var/run/openvswitch/ovs-vswitchd.pid \
   --overwrite-pidfile \
   --detach \
    unix:/host/var/run/openvswitch/db.sock

# create default public external internet bridge
ovs-vsctl --db=unix:/host/var/run/openvswitch/db.sock add-br $OVS_BRIDGE 2> /dev/null
ovs-vsctl --db=unix:/host/var/run/openvswitch/db.sock add-port $OVS_BRIDGE $NODE_NIC 2> /dev/null

# Isolataion for default public external network
ovs-ofctl --db=unix:/host/var/run/openvswitch/db.sock del-flows $OVS_BRIDGE
ovs-ofctl --db=unix:/host/var/run/openvswitch/db.sock add-flow $OVS_BRIDGE "priority=0,actions=drop"
ovs-ofctl --db=unix:/host/var/run/openvswitch/db.sock add-flow $OVS_BRIDGE "in_port=$NODE_NIC,actions=NORMAL"
ovs-ofctl --db=unix:/host/var/run/openvswitch/db.sock add-flow $OVS_BRIDGE "actions=output=$NODE_NIC"