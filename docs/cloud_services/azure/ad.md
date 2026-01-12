---
slug: managing-users-azure-entra-id
title: Managing Users in Azure Entra ID
authors: [kgal-akl]
tags: [microsoft, azure, cloud, rbac, user, az, cli, active_directory, ad, entra_id]
---

## Creating a Batch of Groups

```bash
GROUP_COUNT=200
GROUP_NAME_PREFIX=kbbgl-dev-group
for i in {1..$GROUP_COUNT};do 
    az ad group create \
    --display-name "$GROUP_NAME_PREFIX-$i" \
    --mail-nickname "$GROUP_NAME_PREFIX-$i";
done

# Listing groups
az ad group list \
--query "[].displayName" \
-o tsv | \
grep "$GROUP_NAME_PREFIX" | \
wc -l

200
```

## Retrieve Groups the User is Assigned to

```bash
END_USER_ID="kbbgl+dev#@somezuretenant.onmicrosoft.com"

# Get the user's object ID first
USER_OBJECT_ID=$(az ad user show --id "$END_USER_ID" --query id -o tsv)
# List all groups the user belongs to
az ad user get-member-objects --id "$END_USER_ID" --security-enabled-only false
```

```json
[
  {
    "displayName": "test-inv-3",
    "id": "UUIDv4_3"
  },
  {
    "displayName": "test-inv-2",
    "id": "UUIDv4_2"
  },
  {
    "displayName": "test-inv-1",
    "id": "$UUIDv4_1"
  }
]
```

```bash
# Or get detailed group info
az ad user get-member-of --id "$END_USER_ID" --query "[].{id:id, displayName:displayName}" -o table
```

