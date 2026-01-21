---
slug: managing-active-connections
title: How to Manage Active Network Connections
description: How to Manage Active Network Connections
authors: [kgal-akl]
tags: [network, testing, development, vm, linux, kernel, traffic, latency, performance, chaos, connectivity]
---


```bash
sudo ss -K src :60032 dst $IP

sudo ss -K dst $IP dport = $PORT
```

If `ss -K` isn't available on your system, you can use `tcpkill`. It works by sniffing the traffic and injecting "RST" (reset) packets to tear down the connection.

```bash
sudo apt install dnsiff
sudo tcpkill -i $IFACE_ host $IP
```