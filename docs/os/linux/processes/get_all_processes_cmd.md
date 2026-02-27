---
slug: os-linux-processes-get-all-processes-cmd
title: "Get All Processes Command Line Arguments"
authors: [kbbgl]
tags: [os, linux, processes, get_all_processes_cmd]
---

# Get All Processes Command Line Arguments

```bash
# pgrep --list-full $PROCESS_NAME

me@me-xsup1757-rmncyzpy:~$ pgrep --list-full server

12328 /usr/local/app/server
14853 /usr/local/app/server -conf /usr/local/app/tenants/acc_Test/server.conf -logfile /var/log/app/acc_Test/server.log -public /usr/local/app/dist -loglevel info -tenant -onetimeconf /tmp/otc66844347
```
