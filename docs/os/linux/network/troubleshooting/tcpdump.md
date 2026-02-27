---
slug: os-linux-network-troubleshooting-tcpdump
title: "Packet Capture"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, tcpdump]
---

# Packet Capture

Traffic on interface `eth0`, with ASCII decoding,

```bash
tcpdump -i eth0 -A port 5672 -s 0 -tttt -C 10 -w rabbit.pcap
```

All traffic between hosts `KievMNQ1/2`

```bash
tcpdump host KievMNQ1 and KievMNQ2 -i 2 -tttt -w
tcpdump.pcap
 ```

 All traffic to specific host:

```bash
tcpdump -A -s 0 dst app-mongod-0.app-mongodb-service.app.svc.cluster.local
```
