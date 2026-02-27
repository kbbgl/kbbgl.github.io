---
slug: os-linux-network-addresses-request-ip-change-from-dhcp-server
title: "Request IP Change from DHCP"
authors: [kbbgl]
tags: [os, linux, network, addresses, request_ip_change_from_dhcp_server]
---

# Request IP Change from DHCP

```bash
dhclient eth0
```

The `dhclient` command sends a `DHCPDISCOVER` request from the network interface specified (here, `eth0` ). It then receives an offer ( `DHCPOFFER` ) from the DHCP server and confirms the IP assignment to the DHCP server with a dhcp request.
