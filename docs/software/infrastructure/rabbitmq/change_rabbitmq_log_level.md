---
slug: rabbitmq-change-log-level
title: Change Log Level in Kubernetes
authors: [kbbgl]
tags: [rabbitmq, k8s, logs, debug]
---

Edit `ConfigMap`:

```bash
kubectl edit cm app-rabbitmq
```

Add
`log.console.level = debug`.

Recycle `Pod`s:

```bash
kubectl delete pod app-rabbitmq
```
