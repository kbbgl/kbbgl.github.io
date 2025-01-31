---
slug: aws-cli-cheatsheet
title: AWS CLI Cheatsheet
description: A cheatsheet for Amazon Web Services CLI. See https://aws.amazon.com/cli/ for more information.
authors: [kgal-akl]
tags: [cheatsheet, aws, cli]
---

## Authentication

## Log In Using Profile

```bash
aws sso login --profile dev
```

## EC2

### Stop EC2 Instance

```bash
aws ec2 stop-instances --instance-ids $EC2_INSTANCE
```

### Stop EC2 Instance


```bash
aws ec2 start-instances --instance-ids $EC2_INSTANCE
```

### Associate EC2 with IAM Role

```bash
aws ec2 associate-iam-instance-profile --instance-id "$EC2_INSTANCE" --iam-instance-profile Name="$AWS_IAM_ROLE_NAME"
```
```