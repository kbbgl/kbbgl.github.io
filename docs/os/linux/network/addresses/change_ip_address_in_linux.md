---
slug: os-linux-network-addresses-change-ip-address-in-linux
title: "Change IP Address"
authors: [kbbgl]
tags: [os, linux, network, addresses, change_ip_address_in_linux]
---

# Change IP Address

```bash
ifconfig $INTERFACE $NEW_IP 
ifconfig eth0 192.168.1.12 
```

```bash
ifconfig eth0 192.168.181.115 netmask 255.255.0.0 broadcast 192.168.1.255
```
