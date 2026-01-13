---
title: Managing Local Kubernetes Clusters using Kubernetes in Docker (kind)
slug: managing-local-k8s-clusters-with-kind
app2or: kgal-akl
tags: [devops, k8s, kubernetes, kind, local_dev, local, dev, development]
---

`kind` is a tool that we use to quickly create local development Kubernetes cluster.

## Configuration

Here's a sample configuration to

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: kgal-localhost-prod-legacy
networking:
  apiServerAddress: "127.0.0.1"
  apiServerPort: 6443

nodes:
- role: control-plane
  
extraPortMappings:
- containerPort: 30000
  hostPort: 8000
  protocol: TCP
- containerPort: 30080
  hostPort: 8080
  protocol: TCP
- containerPort: 30081
  hostPort: 8081
  protocol: TCP
- containerPort: 30088
  hostPort: 18888
  protocol: TCP
```

## Create Cluster

```bash
kind create cluster \
--config config.yaml \
--kubeconfig .kubeconfig
```
