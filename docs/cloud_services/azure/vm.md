---
slug: managing-azure-vm
title: Managing VM in Azure
authors: [kgal-akl]
tags: [microsoft, azure, cloud, vm, az, cli, vms, virtual_machine]
---

## Get VM Resource ID

```bash
az vm show \
--name $VM_NAME \
--resource-group $RESOURCE_GROUP \
--query id \
--output tsv
```

## Assign A Managed Identity to VM

```bash
az vm identity assign \
--name $AZURE_VM_NAME \
--resource-group $AZURE_RESOURCE_GROUP
```