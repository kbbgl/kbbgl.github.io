---
slug: container-mem
title: Container Memory
authors: [kbbgl]
tags: [docker, memory, resources]
---



### From Host

```bash
# Take full container ID
docker ps --no-trunc -q --filter "name=app/*"                
c023bb17cb48d1249e74437efdb585a3cd4432353b3fc8ebea1127377b556e2e

CONTAINER_ID=c023bb17cb48d1249e74437efdb585a3cd4432353b3fc8ebea1127377b556e2e

# watch changes
watch echo "container ID $CONTAINER_ID is using $(($(cat /sys/fs/cgroup/memory/docker/$CONTAINER_ID/memory.usage_in_bytes)/1000000))MB of memory"
```


### In Container

```bash
cat /sys/fs/cgroup/memory/memory.usage_in_bytes to get memory usage.
```