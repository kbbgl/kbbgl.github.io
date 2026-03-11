---
slug: managing-azure-keyvault
title: Managing Key Vault in Azure
authors: [kgal-akl]
tags: [microsoft, azure, cloud, vm, az, cli, keyvault, vault, secrets, cybersecurity]
---


## List Key Vaults

```bash
az keyvault list \
--resource-group $AZURE_RESOURCE_GROUP
```

## Secrets in Vault

```bash
az keyvault secret list \
--vault-name "$AKV_NAME" \
--query "length(@)"
```

## Create Bulk Secrets

```bash
for i in {1..$COUNT}; do
rand=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 256 | head -n 1)
az keyvault secret set \
--vault-name "$AKV_NAME" \
--name "loadtest-secret-$i" \  
--value "$rand"
done
```