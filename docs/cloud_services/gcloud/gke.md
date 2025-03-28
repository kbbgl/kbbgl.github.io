---
title: Google Kubernetes Engine
slug: gke
tags: [kubectl, cloud, cli, gke, k8s, kubernetes]
authors: [kgal-akl]
---



## Configure GKE cluster with `kubectl` 
After creating a GKE cluster, we can configure `kubectl` to interact with the GKE cluster using:

```bash
GKE_CLUSTER_NAME=kgal-gke-cluster
GCP_PROJECT=kgal-dev
GCP_REGION=us-east4

gcloud container clusters get-credentials $GKE_CLUSTER_NAME --region $GCP_REGION --project $GCP_PROJECT
```

This will append a new user, context and cluster to `~/.kube/config`.

## Configure Ingress

```bash
kubectl create clusterrolebinding cluster-admin-binding \
  --clusterrole cluster-admin \
  --user $(gcloud config get-value account)
```

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.1/deploy/static/provider/cloud/deploy.yaml
```

See [this](https://kubernetes.github.io/ingress-nginx/deploy/#gce-gke) for more info.