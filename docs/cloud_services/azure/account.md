---
slug: managing-azure-account
title: Managing Account in Azure
authors: [kgal-akl]
tags: [microsoft, azure, cloud, account, az, cli]
---

## Retrieve Tenant ID

```bash
az account show \
--query tenantId \
--output tsv
```

## Retrieve Subscription ID

```bash
az account show \
--query id \
--output tsv
```


