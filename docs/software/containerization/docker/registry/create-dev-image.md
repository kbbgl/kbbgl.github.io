---
slug: how-to-create-docker-image-cloud
title: How To Create Docker Image in Cloud Provider Registry
authors: [kgal-akl]
tags: [docker, registry, images, google_cloud, aws_ecr, ecr, artifacts]
---

When developing features for certain applications (where these application are run as containers, e.g. [KEDA](https://github.com/kedacore/keda), [Dapr](https://github.com/dapr/components-contrib)), we need to create a new Docker image to test out our changes. This guide explains how to do that using Google Cloud Artifact Registry and Amazon Web Services ECR.

## GCP

### Prerequisites

Enable the Artifact Registry service using the Google Cloud CLI:

```bash
gcloud services enable artifactregistry.googleapis.com
```

Configure the relevant Google Cloud Platform details:
```bash
export GCP_LOCATION=us-east
export GCP_PROJECT=keda-test-env
export GCP_REPO_NAME=keda
```

### Create a Repository

```bash
gcloud artifacts repositories create $GCP_REPO_NAME \
    --repository-format=docker \
    --location=$GCP_LOCATION \
    --project=$GCP_PROJECT \
    --description="Docker repository for $GCP_REPO_NAME-related images"
```
### Make Repository Public

```bash
gcloud artifacts repositories add-iam-policy-binding $GCP_REPO_NAME \
    --location=$GCP_LOCATION \
    --member='allUsers' \
    --role='roles/artifactregistry.reader' \
    --project=$GCP_PROJECT
```
### Make Repository Private

```bash
gcloud artifacts repositories remove-iam-policy-binding $GCP_REPO_NAME \
    --location=$GCP_LOCATION \
    --member='allUsers' \
    --role='roles/artifactregistry.reader' \
    --project=$GCP_PROJECT
```


## AWS ECR

### Configuration

```bash
export AWS_ACCOUNT_ID=1234
export AWS_REGION="us-east-1"  # e.g., us-east-1
export REPO_NAME=dapr
```

### Create Repository

```bash
aws ecr create-repository --repository-name $REPO_NAME --region $AWS_REGION
```

### Authenticate Docker

```bash
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
```

### List Images

```bash
aws ecr list-images --repository-name $REPO_NAME --no-cli-pager
```