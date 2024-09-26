---
slug: rabbitmq-add-rm-cluster-node
title: Add/Remove Node from Cluster
authors: [kbbgl]
tags: [rabbitmq, node, cluster, high_availability]
---

## Add

```bash
rabbitmqctl stop_app

rabbitmqctl join_cluster rabbit@$SOME_NODE_HOSTNAME

rabbitmqctl start_app
```


## Remove

Stop service (run on node to be removed):
```bash
rabbitmqctl stop_app
```

From other node, run:
```bash
rabbitmqctl forget_cluster_node rabbit@$SOME_NODE_HOSTNAME
```
