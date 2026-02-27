---
slug: os-linux-network-troubleshooting-finding-process-listening-port
title: "How to Find Process listening on port"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, finding_process_listening_port]
---

# How to Find Process listening on port

```bash
PORT=4445
lsof -n -i :$PORT
```
