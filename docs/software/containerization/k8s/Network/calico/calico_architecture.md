### What Calico does

It uses Layer 3 (network) to manage Kubernetes networking.

### Architecture

![](https://miro.medium.com/max/700/1*XvDzIPDcDrJqK4WJsR28UA.png)

#### Datastore (`etcd`)

`etcd` holds all cluster status information and is the entrypoint for `calicoctl`.

#### Container Network Interface (CNI)

A CNI plugin is a set of libraries that allow creation of network interfaces in containers.

In Calico, the CNI assigns a virtual ethernet pair and an IP address to a container/Pod. This manages communication between the container and the host.

The Pod will have a default route to the host's ethernet interface. 

```bash
❯ kubectl exec calico-node-tlzzh -it -- bash

ip route | grep default

# default via 10.50.0.1 dev ens160 proto dhcp src 10.50.43.145 metric 100
```

The packet sent from container to host machine will hit the host's routing table.

#### BIRD (BGP Client)

BGP (Border Gateway Protocol) is a standard protocol for exchanging routing information between two routers in a network. Each router running BGP has one or more BGP peers - other routers which they are communicating with over BGP. 

Calico uses BIRD as the GDP client to establish sessions between nodes and share their routes to allow communication to and from nodes within a Kubernetes cluster.

When the CNI assigns an IP address to a Pod, it will also add them to the routing table. When this addition is done, BIRD will also advertise the new route to the BGP peers.


#### FELIX

Felix is used in Calico for policy enforcement. It will use the `etcd` to manage Node status, endpoints and network policies. It's a daemon running on each host node or container, communicates with Calico, BGP clients and kernel routing.

#### IPAM (IP Address Management)

Used during the process of creating a Calico network. It will monitor the IP addresses assigned and in use. 

