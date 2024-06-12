# Beaver
## _**Distributed openvswitch agent for kubernetes.**_

> ### I works in combination with Multus CNI and OVS cni
> ### _It needs beaver, multus cni, ovs cni, to work with_


## Environment variables

|      Env             | Default value |
| -------------------- |:-------------:|
| DEFAULT_OVS_BRIDGE   | ext           |
| NODE_NIC             | eno33559296   |


you can also customize the daemonset to ensure running on prefered worker/compute nodes.

## Installation
Install the beaveragent daemonset using kubectl on compute nodes:
```bash
kubectl apply -f ./manifests/beaveragent.yaml
```

#### well tested with talos.
> The namespace must has the label ```pod-security.kubernetes.io/enforce: privileged```

e.g.
```yaml
apiVersion: v1
kind: Namespace
metadata:
  labels:
    kubernetes.io/metadata.name: ik8s-system
    pod-security.kubernetes.io/enforce: privileged
  name: ik8s-system
```