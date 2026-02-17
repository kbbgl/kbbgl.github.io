---
slug: aws-cli
title: How to Configure AWS CLI
authors: [kbbgl]
tags: [cloud, aws, cli]
---

## Configure AWS CLI Profile

```bash
AWS_PROFILE=dev
aws configure --profile $AWS_PROFILE
```

Access Key ID and Secret Access Key can be found in:

```bash
cat ~/.aws/credentials
```

## Login

```bash
aws sso login --profile $AWS_PROFILE
```

## Configuration File

The configuration file can be found in:

```bash
cat ~/.aws/config
```
