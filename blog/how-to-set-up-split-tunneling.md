---
slug: how-to-set-up-split-tunneling
title: How To Set Up Split Tunneling with VPN
description: This blog post describes how to set up split tunneling with a VPN on MacOS.
authors: [kgal-akl]
tags: [vpn,network,route,routing,macos,unix]
date: 2025-10-19
---

## Introduction

A lot of our work nowadays requires using and connecting to a Virtual Private Networks (VPNs) in order to access certain resources (e.g. databases, websites, REST APIs) that were deemed important to protect from the public internet. When we connect to the VPN, we're able to access these resources. 

The nature of our modern digital work requires simultaneous access to a plethora of services. Some of these services require an active VPN connection and some can be accessed without. 

At times, the VPN we need to connect to is geolocated far from us. In addition, the VPN can be one that serves the entire company and is not very performant. These factors result in an experience of collective latency accessing resources, restricted and unrestricted ones alike. 

If you find/found yourself in this type of situation before, this post will explain how you can circumvent that and suffer latency only when accessing the restricted resources instead of all resources by modifying the operating system routing tables in what's called 'Split Tunneling'. 

To do this you will need to have root/administrator access to the UNIX operating system.

We begin by collecting the relevant information and later performing the modifications.

<!-- truncate -->

## Collection

The first thing we will need to do is to collect the IP addresses of the restricted services that we want to only access using the VPN. In the example below I will use `global.svc.dev`, `us.svc.dev`, `eu.svc.dev` as the hostnames for the services that require access.

### Service Hostname to CIDR IP Address
If we don't have those readily-available, we'll need to connect to the VPN as we usually do and lookup the service:

```bash
❯ RESTRICTED_SERVICES_HOSTNAMES=("global.svc.dev" "us.svc.dev" "eu.svc.dev")
❯ for HOSTNAME in "${RESTRICTED_SERVICES_HOSTNAMES[@]}"; do
    CIDR_ADDRESS="$(nslookup HOSTNAME | grep Address | tail -1 | cut -d" " -f2)/32";
    echo "$HOSTNAME => $CIDR_ADDRESS
done

global.svc.dev => 123.123.123.123/32
us.svc.dev => 213.213.213.213/32
eu.svc.dev => 122.133.111.222/32
```

### Network Interfaces

Now that we have the IP addresses of the restricted services, we need to find the names of the network interfaces which we will manipulate. The relevant network interfaces are the VPN and the local one (the one that is used when the VPN is not needed).

To find both interfaces we use `ifconfig`. 
For the VPN interface it'll usually be named something like `utun[1-4]` or `ppp[0-4]`. As for the local interface, it'll either be `en0` or `en1`. Here's a shortened example output:

```bash
❯ ifconfig

en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
        inet 10.100.1.165 netmask 0xffffff00 broadcast 10.100.1.255
        status: active
...
utun4: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1340
        inet 192.168.111.222 --> 192.168.111.222 netmask 0xffffffff
```

We can see from the output above that the local interface is named `en01` and is currently assigned IP address 10.100.1.165 and the VPN interface is named `utun4` and has an internal IP address of 192.168.111.222. The VPN client application will usually print out this IP address so you can always compare it to the one you see from the `ifconfig` output.

Not necessary but we can also see the default gateway IP (usually the ISP-provided router) used by the local network interface:

```bash
❯ netstat -nr | grep default | grep en0
default            192.168.1.1        UGScIg                en0
```

## Change Routing Tables

Now that we have all the information we need, we can modify the routing tables.

The routing instructions needed are:

- Reset routing of requests to the VPN interface on the default network:

```bash
❯ sudo route delete -net default -interface utun4
```

- Set route requests to the local interface on the default network.

```bash
❯ sudo route add -net 0.0.0.0 -interface en0
```

- Route each request sent to a restricted IP address to the VPN interface.

```bash
❯ for ip in "${RESTRICTED_SITES[@]}";do
        sudo route add -net $ip -interface $VPN_INTERFACE;
done
```

Putting it all together in a script `/path/to/split_tunnel.sh`:

```bash
#!/usr/bin/env bash

❯ VPN_INTERFACE="utun4"
❯ LOCAL_INTERFACE="en0"
❯ RESTRICTED_SERVICES_HOSTNAMES=("global.svc.dev" "us.svc.dev" "eu.svc.dev")
❯ RESTRICTED_SERVICES_CIDR_ADDRESSES=()
❯ for HOSTNAME in "${RESTRICTED_SERVICES_HOSTNAMES[@]}"; do
    CIDR_ADDRESS="$(nslookup HOSTNAME | grep Address | tail -1 | cut -d' ' -f2)/32";
    RESTRICTED_SERVICES_CIDR_ADDRESSES+=(CIDR_ADDRESS)
done

❯ sudo route delete -net default -interface $VPN_INTERFACE
# output delete net default: gateway $VPN_INTERFACE

❯ sudo route add -net 0.0.0.0 -interface $LOCAL_INTERFACE
# output add net 0.0.0.0: gateway $LOCAL_INTERFACE

❯ for ip in "${RESTRICTED_SERVICES_CIDR_ADDRESSES[@]}";do
        sudo route add -net $ip -interface $VPN_INTERFACE;
done

# output
# add net 123.123.123.123/32: gateway $VPN_INTERFACE
# add net 213.213.213.213/32: gateway $VPN_INTERFACE
# add net 122.133.111.222/32: gateway $VPN_INTERFACE
```

After running this script, we'll be able to access the non-restricted services without going through the VPN. 

Keep in mind that these changes will not persistent and will likely be reset every time you disconnect/connect to the VPN. The following section explains how to set up automatically running the script when the VPN interface is detected.

## (Optional) Split Tunnel on VPN Connection on MacOS

The goal is for us to run the script above every time we connect to the VPN. 

We first need to confirm whether the VPN interface name assigned is constant. To do that, we can list the names of the interfaces before and after we connect and disconnect from the VPN:

```bash
❯ watch ifconfig -l

# not connected to VPN
lo0 en1 bridge0 ap1 en0 utun0 utun1 utun2 utun3

# connected to VPN
lo0 en1 bridge0 ap1 en0 utun0 utun1 utun2 utun3 utun4
```

As we can see from the output above, `utun4` was added when we connected to the VPN so we can be pretty confident that the name of the VPN interface is constant.

Now that we confirmed the VPN interface name is constant, we can set up a `launchd` agent to monitor for changes in the specific VPN interface device (e.g. `/dev/utun4`).

To do this, we need to create a `launchd` property list (`plist`) file in `~/Library/LaunchAgents/com.kbbgl.split_tunnel.plist`:

```xml title="~/Library/LaunchAgents/com.kbbgl.split_tunnel.plist"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.kbbgl.split_tunnel</string>
    <key>ProgramArguments</key>
    <array>
        <!-- TODO change path to script -->
        <string>/path/to/split_tunnel.sh</string>  
    </array>
    <key>WatchPaths</key>
    <array>
        <!-- TODO change name of VPN interface -->
        <string>/dev/utun4</string>  </array>
    <key>RunAtLoad</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/split_tunnel.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/split_tunnel.log</string>
</dict>
</plist>
```

And load the new agent:
```bash
❯ launchctl load ~/Library/LaunchAgents/com.kbbgl.split_tunnel.plist
```
