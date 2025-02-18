---
id: delete-all-evicted
title: Delete All Evicted Pods
slug: delete-all-evicted
tags: [kubectl, k8s, admin]
author: kgal-akl
---

## Delete All Evicted Pods

```bash
kubectl get pods -A | grep Evicted | awk '{print $1}' | xargs kubectl delete pod -A
```
