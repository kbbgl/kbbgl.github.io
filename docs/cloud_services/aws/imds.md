---
title: AWS IAM Authentication using IMDSv2
slug: aws-iam-auth-imdsv2
app2or: kgal-akl
tags: [devops, eks, aws, iam, imds]
---

```shell
TOKEN=`curl -X PUT "http://169.254.169.254/latest/api/token" -H "X-aws-ec2-metadata-token-ttl-seconds: 21600"`
curl -H "X-aws-ec2-metadata-token: $TOKEN" http://169.254.169.254/latest/meta-data/
```

