---
title: AWS EC2 Management
slug: aws-ec2
app2or: kgal-akl
tags: [devops, eks, aws, load_balancer, network, vm]
---

## Load Balancers

AWS has 2 different types of Load Balancers:

- Application Load Balancer (ALB) which sits on Layer 7 (e.g. HTTP)
- Network Load Balancer (NLB) which sits at Layer 3/4 (e.g. TCP)

### Checking Load Balancer Listeners 

To check listeners for a specific Load Balancer:

```bash
aws elbv2 describe-listeners --load-balancer-arn "arn:aws:elasticloadbalancing:$AWS_REGION:$AWS_ACCOUNT_ID:loadbalancer/net/$LOAD_BALANCER_NAME"
```

## Targets

### Describing Target Health

```bash
aws elbv2 describe-target-health --target-group-arn "arn:aws:elasticloadbalancing:$AWS_REGION:$AWS_ACCOUNT_ID:targetgroup/$TARGET_GROUP_NAME"
```

## Security Groups

### Describe Security Groups

```bash
aws ec2 describe-security-groups --group-ids sg-$SECURITY_GROUP_ID
```