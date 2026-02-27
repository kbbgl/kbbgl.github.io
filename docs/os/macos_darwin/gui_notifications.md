---
slug: os-macos-darwin-gui-notifications
title: "Displaying Native GUI Notifications from Bash"
authors: [kbbgl]
tags: [os, macos, darwin, gui_notifications]
---

# Displaying Native GUI Notifications from Bash

```bash
#!/bin/bash
sleep 10
osascript -e "display notification \"Task #1 was completed successfully\" with title \"notify.sh\""
```
