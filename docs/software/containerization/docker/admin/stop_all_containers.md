---
slug: stop-all-containers
titie: Stop All Containers
authors: [kbbgl]
tags: [docker, containers]
---

```bash
docker stop $(docker ps -a -q)
```
