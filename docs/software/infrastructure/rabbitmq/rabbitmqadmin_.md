---
slug: rabbitmq-admin
title: RabbitMQ Admin
authors: [kbbgl]
tags: [rabbitmq, admin, cheatsheet, k8s]
---

[Source](https://stackoverflow.com/a/53379019/1710131)

```bash
# Find
ip=$(kubectl get endpoints app-rabbitmq-ha -o=jsonpath='{.subsets[0].addresses[0].ip}')

# Download CLI
wget http://$ip:15672/cli/rabbitmqadmin

# Change permission
chmod 777 rabbitmqadmin

./rabbitmqadmin -H $ip list queues name messages consumers

./rabbitmqadmin -H $ip show overview
+------------------+--------------------------------------------------------------------------------------+-----------------------+----------------------+
| rabbitmq_version |                                     cluster_name                                     | queue_totals.messages | object_totals.queues |
+------------------+--------------------------------------------------------------------------------------+-----------------------+----------------------+
| 3.7.12           | rabbit@app-rabbitmq-ha-0.app-rabbitmq-ha-discovery.app.svc.cluster.local | 0                     | 284                  |
+------------------+--------------------------------------------------------------------------------------+-----------------------+----------------------+
```
