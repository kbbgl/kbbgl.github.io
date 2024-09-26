# Container Networking

([source](https://eevans.co/blog/deconstructing-kubernetes-networking/))

The technology at the heart of containerization is Linux namespacing, which allows for isolation of various resources without full OS-level virtualization. The kind of namespace we care about here is a network namespace (aka `netns`), which provides a full copy of the Linux networking stack that’s completely isolated from the “main” one.

Every Kubernetes pod gets its own network namespace (if there are multiple containers in a pod, they share the same namespace).

## Setting up network on Pod

we use two Linux networking features:

* **Bridges** - are virtual network switches that live within the Linux kernel.

* **veths** - virtual ethernet, similar to virtual network cables that attach two network devices.

These features combined enable us to communicate across namespaces.

To set up networking on a Pod:

1. Add a bridge (1 per host)

1. Create a `netns` for the Pod (1 per Pod)

1. Add `veth` pair with one end of the pair in the Pod's `netns` and the other connected to the bridge.

1. Assign IP addresses and add routes.

Diagram:
![bridgenet](https://eevans.co/blog/deconstructing-kubernetes-networking/bridgenet.svg)

```bash
# Create a netns named "test"
$ sudo ip netns add test
# Create a bridge named "test0"
$ sudo ip link add name test0 type bridge
# Create a veth pair with testveth0<->eth0 as the endpoints
$ sudo ip link add testveth0 type veth peer name eth0
# Move the eth0 side of the veth pair to the "test" netns
$ sudo ip link set eth0 netns test
# "Plug in" the testveth0 side of the veth pair to the test0 bridge
$ sudo ip link set testveth0 master test0
# Bring up the testveth0 side of the veth pair
$ sudo ip link set testveth0 up
# Bring up the eth0 side of the veth pair
$ sudo ip -n test link set eth0 up
# List network devices in the "main" namespace
$ ip link
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: ens3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether fa:16:3e:cf:81:3d brd ff:ff:ff:ff:ff:ff
11: test0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether 4a:13:0d:fb:9f:a4 brd ff:ff:ff:ff:ff:ff
15: testveth0@if14: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master test0 state UP mode DEFAULT group default qlen 1000
    link/ether 4a:13:0d:fb:9f:a4 brd ff:ff:ff:ff:ff:ff link-netnsid 1
# List network devices in the "test" namespace
$ sudo ip -n test link
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
14: eth0@if15: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP mode DEFAULT group default qlen 1000
    link/ether a6:4f:92:e2:4b:1e brd ff:ff:ff:ff:ff:ff link-netnsid 0
```

So any packet that goes across nodes will have to deal with three hops:

1. It will have to get from the pod to the relevant bridge (via the veth pair we discussed in the last post).

1. It will have to get to the destination node over the “real” network adapter, ens3 in my case. (Implicitly, the packet has to get from the bridge to the network adapter, but that’s handled internally by the kernel.)

1. Once the packet arrives at its destination host, it has to be routed through the relevant bridge to the destination pod.

![multi-node-net](https://eevans.co/blog/kubernetes-multi-node/multi-node-network.svg)
