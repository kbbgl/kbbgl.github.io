---
slug: check-cert-cn
title: Check CN of Certificate
description: Check CN of Certificate
authors: [kbbgl]
tags: [cybersecurity,certificates,encryption]
---


```bash
openssl x509 -noout -subject -in certificate.pem
#subject=CN = some.domain.com
```
