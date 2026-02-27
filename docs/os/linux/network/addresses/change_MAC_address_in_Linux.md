---
slug: os-linux-network-addresses-change-mac-address-in-linux
title: "Change MAC Address"
authors: [kbbgl]
tags: [os, linux, network, addresses, change_mac_address_in_linux]
---

# Change MAC Address

```bash
ifconfig eth0 down 
ifconfig eth0 hw ether 00:11:22:33:44:55 
ifconfig eth0 up
```
