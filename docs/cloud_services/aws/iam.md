---
title: AWS IAM
slug: aws-iam
app2or: kgal-akl
tags: [devops, eks, aws, iam, imds]
---

## IMDS

### Get Token from IMDS

```shell
TOKEN=`curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"`
curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/
```

