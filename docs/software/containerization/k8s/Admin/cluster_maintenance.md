---
slug: k8s-cluster-maintenance
title: Kubernetes Cluster Maintenance
description: How to Perform Maintenance on a Kubernetes Cluster
authors: [kbbgl]
tags: [k8s, kubernetes, maintenance]
---

## Prepare the Node (Cordon & Drain)
Before touching the OS, you need to tell the Kubernetes scheduler to stop sending new pods to this node and gracefully evict the ones currently running.

### Cordon

Marks the node as "SchedulingDisabled".

```bash
kubectl cordon <node-name>
```

### Drain

This is the "evacuation." It terminates the pods gracefully so they can move to other nodes (if you have them) or shut down cleanly.

```bash
kubectl drain <node-name> --ignore-daemonsets --delete-emptydir-data
```

## Stop the K3s Service
Even though you've drained the pods, the K3s engine is still running. It’s cleaner to stop the service before updating the underlying OS.

```bash
sudo systemctl stop k3s
```

## Update OS

```bash
sudo apt update && sudo apt upgrade -y
sudo reboot
```

## Bring the Cluster Back Online
Once the VM restarts, K3s usually starts automatically (unless you've disabled the service).

### Check Service Status

```bash
sudo systemctl status k3s
```

### Uncordon the Node

This tells Kubernetes that the node is healthy and ready to host pods again.

```bash
kubectl uncordon <node-name>
```