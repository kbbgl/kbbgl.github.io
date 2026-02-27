---
slug: os-linux-network-addresses-change-network-information
title: "Change Network Information"
authors: [kbbgl]
tags: [os, linux, network, addresses, change_network_information]
---

# Change Network Information

## New/Change IP address

```bash
ifconfig $INTERFACE $IP_ADDRESS
```

## Change net mask

```bash
ifconfig $INTERFACE $IP_ADDRESS netmask $NET_MASK broadcast $DEFAULT_GATEWAY_IP

# Example
```bash
ifconfig eth0 192.168.1.12 netmask 255.255.255.0 broadcast 192.168.1.1
```

## Change MAC address

```bash
kali > ifconfig eth0 down 
kali > ifconfig eth0 hw ether 00:11:22:33:44:55 
kali > ifconfig eth0 up
```

## Request new DCHP IP Address

```bash
dhclient $INTERFACE
```
