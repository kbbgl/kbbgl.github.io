---
slug: zk-status
title: Get Status
authors: [kbbgl]
tags: [zk, status, k8s]
---
# Status

```bash
for i in $(kubectl get po -l app=zookeeper -o=json | jq -r '.items[] | .status.podIP'); do echo "Node: $i State: $(echo mntr| nc $i 2181)"; done
```

```bash
zk_version 3.4.8--1, built on 02/06/2016 03:18 GMT
zk_avg_latency 0
zk_max_latency 0
zk_min_latency 0
zk_packets_received 10
zk_packets_sent 9
zk_num_alive_connections 1
zk_outstanding_requests 0
zk_server_state standalone
zk_znode_count 4
zk_watch_count 0
zk_ephemerals_count 0
zk_approximate_data_size 27
zk_open_file_descriptor_count 25
zk_max_file_descriptor_count 1048576
```
