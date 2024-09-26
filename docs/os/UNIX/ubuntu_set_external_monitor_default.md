---
slug: how-to-set-external-monitor-default-ubuntu
title: How to Set Up External Monitor in Ubuntu
authors: [kbbgl]
tags: [ubuntu,monitor]
---

## Copy Monitor Configuration to GNOME Display Manager

```bash
sudo cp ~/.config/monitors.xml ~gdm/.config/monitors.xml

sudo chown gdm:gdm ~gdm/.config/monitors.xml

reboot
```
