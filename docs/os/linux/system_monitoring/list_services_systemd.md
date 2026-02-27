---
slug: os-linux-system-monitoring-list-services-systemd
title: "List Services"
authors: [kbbgl]
tags: [os, linux, system_monitoring, list_services_systemd]
---

# List Services

## Active Services

```bash
systemctl list-units --type=service --state=running
```
