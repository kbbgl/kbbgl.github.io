---
title: How to Create an EKS Cluster using eksctl
slug: how-to-create-eks-cluster-ekctl
app2or: kgal-akl
tags: [$AWS_PROFILEops, k8s, kubernetes, kubectl, eksctl, eks]
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


## Interacting With Cluster

In cases when we have an existing EKS cluster that we want to interact with once the AWS access token has expired, we need to first reauthenticate with AWS:

```bash
aws sso login --profile $AWS_PROFILE
```

Once we're authenticated, we can run the following command to let `kubectl` know that it needs to retrieve the cluster access token using the AWS CLI:

```bash
aws eks update-kubeconfig --region $AWS_REGION --name $EKS_CLUSTER_NAME --profile $AWS_PROFILE
```

This will update the `.kubeconfig` file to include the necessary AWS command to retrieve the secret and allow interaction with the EKS cluster:
```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS[redacted]
    server: https://abcdefg.hi2.$AWS_REGION.eks.amazonaws.com
  name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
contexts:
- context:
    cluster: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
    user: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
  name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
current-context: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
kind: Config
users:
- name: arn:aws:eks:$AWS_REGION:$AWS_ACCOUNT_ID:cluster/$EKS_CLUSTER_NAME
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
      - --region
      - $AWS_REGION
      - eks
      - get-token
      - --cluster-name
      - $EKS_CLUSTER_NAME
      - --profile
      - $AWS_PROFILE
      command: aws
```

We can then run `kubectl` commands.