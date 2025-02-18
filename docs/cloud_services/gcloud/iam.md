---
slug: gcloud-iam-cheatsheet
title: Google Cloud CLI IAM Cheatsheet
authors: [kgal-akl]
tags: [cloud, google, gcloud, iam, security, cli, cheatsheet]
---

## Create IAM Role

```bash
gcloud iam roles create $IAM_ROLE_NAME \
	--title=akl_min_permissions \
	--project=$GCP_PROJECT \
	--description="Some description" \
	--permissions=iam.serviceAccounts.get,iam.serviceAccountKeys.get,compute.instances.get,compute.instanceGroups.list
```

## Create a Service Account JSON

```bash
gcloud iam service-accounts keys create svc.json --iam-account=kbbgl-gcp-dev@dev.gserviceaccount.com
```