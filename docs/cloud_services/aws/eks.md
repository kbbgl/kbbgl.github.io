---
title: How to Create an EKS Cluster using eksctl
slug: how-to-create-eks-cluster-ekctl
app2or: kgal-akl
tags: [devops, k8s, kubernetes, kubectl, eksctl, eks]
---


## Log Into AWS Using CLI

```bash
aws sso login --profile $AWS_PROFILE
```

## Create Cluster
```bash
eksctl create cluster \
--name $EKS_CLUSTER_NAME \
--region $AWS_REGION \
--nodegroup-name "kbbgl-nodegroup" \
--node-type "t3.medium" \
--nodes 1 \
--managed
```


## Enable CloudWatch Logging on the EKS for authentication events

```bash
aws eks update-cluster-config \
--name $EKS_CLUSTER_NAME \
--region $AWS_REGION \
--logging '{"clusterLogging":[{"types":["authenticator"],"enabled":true}]}'
```
