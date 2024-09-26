---
slug: zk-get-key-in-znode
title: Get Value of Znode
authors: [kbbgl]
tags: [zk, query, k8s, windows]
---

## Windows

```powershell
zkCli.cmd -server localhost:2181 get /app/path/to/key.nested 
```

## Linux

```bash
zk=$(kubectl get pod -l app=app-zookeeper -o=jsonpath='{.items[0].metadata.name}')
kubectl exec $zk -- zkCli.sh -server localhost:2181 get /app/path/to/key.nested
```
