---
slug: rosa-cheatsheet
title: Red Hat OpenShift Service on AWS (ROSA) CLI Cheat Sheet
description: Red Hat OpenShift Service on AWS (ROSA) CLI Cheat Sheet
authors: [kgal-akl]
tags: [openshift,redhat,rhel,kubernetes,k8s,oc,cheatsheet,rosa]
---

## Install/Configure

```bash
brew install rosa-cli aws-cli openshift-cli
aws sso login --profile $AWS_PROFILE
rosa login --use-auth-code # can also use --token="$TOKEN" from https://console.redhat.com/openshift/token 
```

## Prepare AWS Account

We need to create the IAM users/roles for the ROSA support role, control plane and worker: 
```bash
rosa create account-roles --mode auto
```

We then need to set up the VPCs and subnets needed for the cluster:
```bash
rosa create network
```

## Create Cluster

We can create a cluster interactively:
```bash
rosa create cluster
```

Or using specific arguments using the outputs from `rosa list account-roles`:
```bash
rosa create cluster \
--cluster-name $CLUSTER_NAME \
--domain-prefix kbbgl-gh \
--sts \
--cluster-admin-password $CLUSTER_ADMIN_PASSWORD \
--role-arn arn:aws:iam::$AWS_ACCOUNT_ID:role/ManagedOpenShift-Installer-Role \
--support-role-arn arn:aws:iam::$AWS_ACCOUNT_ID:role/ManagedOpenShift-Support-Role \
--controlplane-iam-role arn:aws:iam::$AWS_ACCOUNT_ID:role/ManagedOpenShift-ControlPlane-Role \
--worker-iam-role arn:aws:iam::$AWS_ACCOUNT_ID:role/ManagedOpenShift-Worker-Role \
--operator-roles-prefix $CLUSTER_NAME \
--region $AWS_REGION \
--version 4.21.4 \
--ec2-metadata-http-tokens required \
--replicas 2 \
--compute-machine-type t3.xlarge \
--machine-cidr $IP/16 \
--service-cidr $IP_2/16 \
--pod-cidr $IP_3/14 \
--host-prefix 23 \
--worker-disk-size 128GiB
```

Once the cluster is ready (takes between 30-60 minutes), we can see the cluster details:

```bash
rosa list clusters

ID       NAME         STATE       TOPOLOGY
1234     $CLUSTER_NAME  installing  Classic (STS)
```

## Setting Up ROSA Cluster Access

We can either create a temporary admin:

```bash
rosa create admin --cluster=$CLUSTER_NAME
oc login <api_url> --username cluster-admin --password <password>
```

Or long term using an Identity Provider such as GitHub or Google:

```bash
rosa create idp --cluster=$CLUSTER_NAME --interactive
rosa grant user cluster-admin --user=<your_username> --cluster=<cluster_name>
```

