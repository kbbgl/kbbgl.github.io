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

See [ECR](../../../../cloud_services/aws/ecr.md)