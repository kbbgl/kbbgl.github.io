---
slug: os-linux-system-monitoring-check-number-of-file-handlers
title: "Check Number of File Handlers"
authors: [kbbgl]
tags: [os, linux, system_monitoring, check_number_of_file_handlers]
---

# Check Number of File Handlers

```bash
# check file descriptors for each process
set -o errexit
set -o pipefail
lsof +c 0 -n -P -u root \
        | awk '/inotify$/ { gsub(/[urw]$/,"",$4); print $1" "$2" "$4 }' \
        | while read name pid fd; do \
                exe="$(readlink -f /proc/$pid/exe || echo n/a)"; \
                fdinfo="/proc/$pid/fdinfo/$fd" ; \
                count="$(grep -c inotify "$fdinfo" || true)"; \
                echo "$name $exe $pid $fdinfo $count"; \
        done
```

```
# output sample
# ....
# crond /usr/sbin/crond 22841 /proc/22841/fdinfo/5 3
# fluent-bit n/a 24196 /proc/24196/fdinfo/20 7294             <----
# fluent-bit n/a 27210 /proc/27210/fdinfo/20 36
# fluent-bit n/a 27210 /proc/27210/fdinfo/62 36
# ....
```
