---
slug: openshift-cheatsheet
title: OpenShift CLI Cheat Sheet
description: OpenShift CLI Cheat Sheet
authors: [kgal-akl]
tags: [openshift,redhat,rhel,kubernetes,k8s,oc]
---

Once you are logged in as a cluster-admin, you can use the OpenShift CLI (`oc`) to inspect the different layers of your cluster. OpenShift is organized into Nodes (infrastructure), Cluster Operators (the brains), and Pods/Projects (the workloads).

## List OpenShift Nodes

```bash
oc get nodes
```

## List Cluster Operators

```bash
oc get clusteroperators
```

## List OpenShift Projects

```bash
oc projects
```

## Show Logged In User
```bash
oc whoami
```

## Miscellaneous
```bash
oc adm top nodes	# Shows CPU and Memory usage for each AWS instance.
oc get events -A	# Shows a stream of recent system logs/errors across the cluster.
oc get routes -A	# Lists all URL endpoints (like the Web Console or your apps).
```