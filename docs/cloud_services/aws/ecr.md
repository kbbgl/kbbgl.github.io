---
slug: aws-ecr
title: Manage Elastic Container Registry
authors: [kbbgl]
tags: [cloud, aws, cli, registry, docker, images, containers]
---

### Configuration

```bash
export AWS_ACCOUNT_ID=1234
export AWS_REGION="us-east-1"
export REPO_NAME=dapr
```

### Create ECR Repository

```bash
aws ecr create-repository --repository-name $REPO_NAME --region $AWS_REGION
```

### Authenticate Docker with ECR Registry

```bash
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
```

### List Images in ECR Registry

```bash
aws ecr list-images --repository-name $REPO_NAME --no-cli-pager
```