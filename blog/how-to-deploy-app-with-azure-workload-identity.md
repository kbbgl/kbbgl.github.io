---
slug: how-to-deploy-app-with-azure-workload-identity
title: How To Deploy Application with Azure Workload Identity
description: This blog post describes how to set up split tunneling with a VPN on MacOS.
authors: [kgal-akl]
tags: [cloud, azure, aks, workload_identity, k8s, kubernetes, federation]
date: 2026-02-24
---

This tutorial is a guide on how to deploy an application in Kubernetes that will authenticate using Azure Workload Identity on Azure Kubernetes Services (AKS).

## Prerequisites

- Access to the Azure CLI and an Azure account.
- `kubectl` installed and access to the AKS cluster.
- `helm`.

See links for more information about [Azure Identity](/docs/cloud_services/azure/azure-identity) and [AKS](/docs/cloud_services/azure/managing-aks-azure).

## Enable OIDC on AKS

1. [Check if the OIDC issuer is enabled in the AKS cluster](/docs/cloud_services/azure/managing-aks-azure#check-workload-identity). [Enable it](/docs/cloud_services/azure/managing-aks-azure#enable-workload-identity) if it's not.

## (Optional) Enable Workload Identity plugin

```bash
az aks update \
--resource-group "$AZURE_RESOURCE_GROUP" \
--name "$AKS_CLUSTER_NAME" \
--enable-workload-identity
```

This will deploy a `Deployment` named `azure-wi-webhook-controller-manager` in the `kube-system` namespace:

```bash
❯ kubectl get deploy -n kube-system
NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
azure-wi-webhook-controller-manager   2/2     2            2           48d
```

This step is optional since we can explicitly specify the application that will use Azure Workload Identity to mount the Azure token as a volume. More on that in a bit.

## Create User Assigned Managed Identity for Application

```bash
# Replace with your preferred names and location
IDENTITY_NAME="app-wi"
IDENTITY_RG="$AZURE_RESOURCE_GROUP"
LOCATION="${AZURE_LOCATION:-eastus}"

az identity create --resource-group "$IDENTITY_RG" --name "$IDENTITY_NAME" --location "$LOCATION"
CLIENT_ID=$(az identity show --resource-group "$IDENTITY_RG" --name "$IDENTITY_NAME" --query clientId -o tsv)
PRINCIPAL_ID=$(az identity show --resource-group "$IDENTITY_RG" --name "$IDENTITY_NAME" --query principalId -o tsv)
TENANT_ID=$(az account show --query tenantId -o tsv)
OIDC_ISSUER=$(az aks show --resource-group "$AZURE_RESOURCE_GROUP" --name "$AKS_CLUSTER_NAME" --query "oidcIssuerProfile.issuerUrl" -o tsv)
```

## Create a Federated Credential

```bash
# namespace and service account name that your test app will use
NAMESPACE="default"
SA_NAME="app-wi-sa"

az identity federated-credential create \
--resource-group "$IDENTITY_RG" \
--name "${IDENTITY_NAME}-fc" \
--identity-name "$IDENTITY_NAME" \
--issuer "$OIDC_ISSUER" \
--subject "system:serviceaccount:${NAMESPACE}:${SA_NAME}"
```

## Install Azure Workload Identity Webhook

This is what injects `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_FEDERATED_TOKEN_FILE`, and the projected token volume into pods that use the label. See [Service Principal](/docs/cloud_services/azure/azure-identity#service-principal-with-secret) for more info on those environmental variables.

```bash
helm repo add azure-workload-identity https://azure.github.io/azure-workload-identity/charts
helm repo update
kubectl create namespace azure-workload-identity-system 2>/dev/null || true
helm upgrade --install workload-identity-webhook azure-workload-identity/workload-identity-webhook \
--namespace azure-workload-identity-system \
--set azureTenantId="$TENANT_ID"s
```

## Create a Kubernetes ServiceAccount

Here is where the link between Kubernetes and Azure Workload Identity happens:

```bash
kubectl create namespace "$NAMESPACE" 2>/dev/null || true
kubectl apply -f - <<EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: $SA_NAME
  namespace: $NAMESPACE
  annotations:
    azure.workload.identity/client-id: "$CLIENT_ID"
EOF
```

As we can see, we annotate the `ServiceAccount` with `azure.workload.identity/client-id: "$CLIENT_ID"`.

## Deploy Application with Workload Identity

```bash
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-wid
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-wid
  template:
    metadata:
      labels:
        app: hello-wid
        azure.workload.identity/use: "true"
    spec:
      serviceAccountName: $SA_NAME
      containers:
      - name: alpine
        image: alpine
        command:
          - "sh"
          - "-c"
          - "echo "Workload Identity tutorial done! Sleeping..." && sleep 10000"
EOF
```

The main things we're doing here are:

1. We set the application `Deployment` to use the `ServiceAccount` we created in the previous step and that is linked to an Azure Workload Identity.
2. We set the `Deployment` `Pod` specification to use Azure Workload Identity by setting the label `azure.workload.identity/use: "true"`.
