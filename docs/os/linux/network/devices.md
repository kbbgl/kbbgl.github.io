---
slug: os-linux-network-devices
title: "Devices"
authors: [kbbgl]
tags: [os, linux, network, devices]
---

# Devices

- `ethX`: Ethernet devices.
- `wlanX`: Wireless devices.
- `brX`: Bridge interfaces
- `vmnetZ`: virtual devices

## Configuring and Controlling Devices

`ip` utility is used to configure, control and query interface parameters and routing. It replaced the `ifconfig` as it's faster and more versatile. **Settings changed using `ip` are not persistent.**

Format:

```bash
ip [options] object {command | help}

# for reading commands from file
ip [-force] -batch /path/to/file
```

Popular `object`:

- `address`: ipv4/6 protocol device address.
- `link`: network devices
- `maddress`: multicast address
- `monitor`: watch for `netlink` messages
- `route`: routing table entry
- `rule`: rule in routing policy db
- `tunnel`: tunnel over ip

Examples:

```bash
# show info for all interfaces
ip link show

# show stats (rx/tx) about eth0 
ip -s link show eth0 
# OR
cat /proc/net/dev

# set ip address
sudo ip addr add 192.168.1.7 dev eth0

# take down interface
sudo ip link set eth0 down

# set mtu to 1480 bytes
sudo ip link set eth0 mtu 1480

# set networking route
sudo ip route add 172.16.1.0/24 via 192.168.1.5
```

The device configuration, information and statistics are held in the following filesystem locations:

```bash
less /proc/net/dev

head -n 4 /proc/net/dev
Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
califf0535d9c94: 98940135  518876    0    0    0     0          0         0 60376340  463276    0    0    0     0       0          0
calie8c5a45cbcd: 2673411   33886    0    0    0     0          0         0  2801482   20678    0    0    0     0       0          0
```

Device configuration and statistics can be found in:

```bash
cat /sys/class/net/docker0/statistics/rx_missed_errors
0
```

Each distro has it's own set of files that include persistent network device configuration (**NIC**).

- RedHat:
  - `/etc/sysconfig/network`
  - `/etc/sysconfig/network-scripts/ifcfg-ethX`
  - `/etc/sysconfig/network-scripts/ifcfg-ethX:Y`
  - `/etc/sysconfig/network-scripts/route-ethX`
- Debian:
  - `/etc/network/interfaces`
- SUSE:
  - `/etc/sysconfig/network`

## Network Manager

The preferred/modern method to interact with network devices is using the Network Manager.

We can change the network configuration using `nmcli`.

Installation and enablement:

```bash
sudo apt install network-manager
sudo /etc/init.d/network-manager start
```

```bash
# contains many useful examples
man nmcli-examples 

# see list of wifi APs
nmcli device wifi list
```
