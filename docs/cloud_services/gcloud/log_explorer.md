---
slug: gcp-log-explorer
title: Google Cloud Platform Log Explorer
authors: [kgal-akl]
tags: [gcp, cloud, log_explorer, logs, monitoring, observability]
---

## Show Load Balancer API Calls 

Only shows requests with 4XX:

```
jsonPayload.@type="type.googleapis.com/google.cloud.loadbalancing.type.LoadBalancerLogEntry"
```

## Show API Calls Reaching from Specific IP

```
jsonPayload.remoteIp="123.123.123.123"
```