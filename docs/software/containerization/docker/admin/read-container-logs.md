---
slug: read-container-logs
title: "Read Container Logs"
authors: [kbbgl]
tags: [docker, logs, containers]
---

```bash
#!/bin/bash

docker_container_id=$1
echo "getting logs for Docker container $docker_container_id"

docker logs $docker_container_id  -f | jq '.args.args[0]'
```
