---
slug: os-linux-processes-get-all-ports-listened-by-process
title: "List All Ports Listened by Process"
authors: [kbbgl]
tags: [os, linux, processes, get_all_ports_listened_by_process]
---

# List All Ports Listened by Process

```bash
lsof -i -P -n | head -1;lsof -i -P -n | grep app
COMMAND PID         USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
cmd    12328         app   37u  IPv6  61967      0t0  TCP *:443 (LISTEN)
cmd    12328         app   51u  IPv6 144140      0t0  TCP 10.180.189.145:443->10.196.103.127:50130 (ESTABLISHED)
cmd    12328         app   55u  IPv6 150225      0t0  TCP 10.180.189.145:443->10.196.103.127:51276 (ESTABLISHED)
cmd    12328         app   56u  IPv6 151870      0t0  TCP 10.180.189.145:443->10.196.103.127:51461 (ESTABLISHED)
cmd    12328         app   59u  IPv4 148170      0t0  TCP 127.0.0.1:54134->127.0.0.1:38731 (ESTABLISHED)
cmd    12328         app   60u  IPv6 154461      0t0  TCP 10.180.189.145:443->10.196.103.127:51891 (ESTABLISHED)
cmd    12328         app   62u  IPv6 155675      0t0  TCP 10.180.189.145:443->10.196.103.127:51587 (ESTABLISHED)
cmd    12328         app   67u  IPv6 156368      0t0  TCP 10.180.189.145:443->10.196.103.127:52154 (ESTABLISHED)
cmd    12328         app   68u  IPv4 144145      0t0  TCP 127.0.0.1:36034->127.0.0.1:38731 (ESTABLISHED)
cmd    12328         app   71u  IPv4 155364      0t0  TCP 127.0.0.1:58592->127.0.0.1:38731 (ESTABLISHED)
cmd    14853         app   13u  IPv6  75169      0t0  TCP *:38731 (LISTEN)
cmd    14853         app   40u  IPv6 155649      0t0  TCP 127.0.0.1:38731->127.0.0.1:54134 (ESTABLISHED)
cmd    14853         app   44u  IPv6 146022      0t0  TCP 127.0.0.1:38731->127.0.0.1:36034 (ESTABLISHED)
cmd    14853         app   55u  IPv6 157193      0t0  TCP 127.0.0.1:38731->127.0.0.1:58592 (ESTABLISHED)
```
