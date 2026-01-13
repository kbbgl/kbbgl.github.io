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
GKE_CLUSTER_ZONE=$GCP_REGION
GCP_SA=my-app-gsa
GCP_SA_NAMEPSACE=default

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


## Create Service Account
```bash
# Create a Google Service Account (GSA)
gcloud iam service-accounts create $GCP_SA \
--display-name="My Application GSA" \
--project $GCP_PROJECT

# Assign required IAM roles to `my-app-gsa@$GCP_PROJECT.iam.gserviceaccount.com`_
gcloud projects add-iam-policy-binding $GCP_PROJECT \
--member="serviceAccount:$GCP_SA@$GCP_PROJECT.iam.gserviceaccount.com" \
--role="roles/storage.objectViewer"
```

## Enable Workload Identity in GKE

```bash
# Check if it's already enabled
gcloud container clusters describe $GKE_CLUSTER_NAME \
--zone=$GKE_CLUSTER_ZONE \
--format="value(workloadIdentityConfig.workloadPool)"

# If it's not enabled, enable it
gcloud container clusters update $GKE_CLUSTER_NAME \
--zone=$GKE_CLUSTER_ZONE \
--workload-pool=$GCP_PROJECT.svc.id.goog
  
# Enable it in the Node Pools as well

gcloud container node-pools update NODE_POOL_NAME \
--cluster $GKE_CLUSTER_NAME \
--zone $GKE_CLUSTER_ZONE \
--enable-workload-identity

# Create an IAM binding allowing the Kubernetes ServiceAccount to impersonate the GCP service account:
gcloud iam service-accounts add-iam-policy-binding \
$GCP_SA@$GCP_PROJECT.iam.gserviceaccount.com \
--role roles/iam.workloadIdentityUser \
--member "serviceAccount:$GCP_PROJECT.svc.id.goog[$GCP_SA_NAMEPSACE/$GCP_SA]"
```
