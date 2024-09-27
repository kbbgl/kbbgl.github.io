---
slug: tcpdump-container
title: TCP Dump in Container
authors: [kbbgl]
tags: [docker, network, debug, tcpdump]
---

[Set up sidecar container with `tcpdump`:](https://developers.redhat.com/blog/2019/02/27/sidecars-analyze-debug-network-traffic-kubernetes-pod/)

1) Add to container to deployment:

```yaml
- name: tcpdump
   image: corfr/tcpdump
   command:
     - /bin/sleep
     - infinity
```

![](https://developers.redhat.com/blog/wp-content/uploads/2019/02/Screenshot-2019-02-20-at-09.17.56.png)

2) Retrieve `docker` container IDs:

```bash
sudo docker ps -a -f "name=tcpdump" --format "{{.ID}}"
f306e8198bfa
91fff43bd3aa
573e90053c1f
```

3) Run container as `root`:

```bash
sudo docker exec -u root -it f306e8198bfa tcpdump -A -s 0
```
