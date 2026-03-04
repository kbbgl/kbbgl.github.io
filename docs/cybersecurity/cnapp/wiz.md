---
slug: wiz-cli-cheatsheet
title: Wiz CLI Cheatsheet
description: Wiz CLI Cheatsheet
authors: [kgal-akl]
tags: [wiz, cnapp, cli, cheatsheet, scan, cybersecurity]
---

## Setup

To set up Wiz CLI for local use, go to _Settings > Deployments_ and create a new one using Wiz CLI CI/CD.
This will generate a `WIZ_CLIENT_{ID,SECRET}`.

## Scan Docker Images

```bash
IMAGE_TO_SCAN="${CONTAINER_REGISTRY}:${IMAGE_NAME}:${IMAGE_TAG}"

wizcli docker scan \
--name "$ISSUE_NAME-rc1" \
--image=$IMAGE_TO_SCAN\
--human-output-file=/tmp/${IMAGE_NAME}:${IMAGE_TAG}"
```


