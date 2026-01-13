---
slug: managing-aks-azure
title: Managing AKS in Azure
authors: [kgal-akl]
tags: [microsoft, azure, cloud, aks, kubernetes, az, cli, k8s]
---

## Workload Identity

### Check Workload Identity

```bash
az aks show \
--resource-group $AZURE_RESOURCE_GROUP \
--name $AKS_CLUSTER_NAME \
--query "oidcIssuerProfile" -o json
```

```json
{
  "enabled": false,
  "issuerUrl": null
}
```

### Enable Workload Identity

```bash
az aks update \
--resource-group $AZURE_RESOURCE_GROUP \
--name $AKS_CLUSTER_NAME \
--enable-oidc-issuer \
--enable-workload-identity
```


### Create Workload Identity

```bash
AZURE_IDENTITY_NAME="kbbgl-azure-test-wid"
az identity create \
--name $AZURE_IDENTITY_NAME \
--resource-group $AZURE_RESOURCE_GROUP
{
  "clientId": "$AZURE_CLIENT_ID",
  "id": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourcegroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.ManagedIdentity/userAssignedIdentities/$AZURE_MANAGED_IDENTITY",
  "location": "$AZURE_LOCATION",
  "name": "$AZURE_MANAGED_IDENTITY",
  "principalId": "$AZURE_MANAGED_IDENTITY_SP_ID",
  "resourceGroup": "$AZURE_RESOURCE_GROUP",
  "systemData": null,
  "tags": {},
  "tenantId": "$AZURE_TENANT_ID",
  "type": "Microsoft.ManagedIdentity/userAssignedIdentities"
}
```

### Link Kubernetes Service Account to AKS Managed Account

```bash
AZURE_FEDERATED_CRED_NAME="kbbgl-aks-fed-cred-test"
az identity federated-credential create \
--name "$AZURE_FEDERATED_CRED_NAME" \
--identity-name $AZURE_IDENTITY_NAME \
--resource-group $AZURE_RESOURCE_GROUP \
--issuer $OIDC_ISSUER_URL \
--subject "system:serviceaccount:${AKEYLESS_K8S_GATEWAY_SERVICE_ACCOUNT_NAMESPACE}:${AKEYLESS_K8S_GATEWAY_SERVICE_ACCOUNT_NAME}" \
--audiences api://AzureADTokenExchange
{
  "audiences": [
    "api://AzureADTokenExchange"
  ],
  "id": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourcegroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.ManagedIdentity/userAssignedIdentities/$AZURE_MANAGED_IDENTITY/federatedIdentityCredentials/$AZURE_FEDERATED_CRED_NAME",
  "issuer": "https://$AZURE_LOCATION.oic.prod-aks.azure.com/$AZURE_TENANT_ID/$AZURE_SUBSCRIPTION_ID/",
  "name": "$AZURE_FEDERATED_CRED_NAME",
  "resourceGroup": "$AZURE_RESOURCE_GROUP",
  "subject": "system:serviceaccount:$AKS_NAMESPACE:$AKS_SERVICE_ACCOUNT_NAME$",
  "type": "Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials"
}
```

### Grant Permissions to Azure Key Vault

```bash
AZURE_KEYVAULT_NAME="kbbgl-azure-kv-test"
az role assignment create \
--role "Key Vault Secrets Officer" \
--assignee $AZURE_MANAGED_IDENTITY_SP_ID \
--scope "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME"
```

```json
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "1970-01-01T00:00:00.00000+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME/providers/Microsoft.Authorization/roleAssignments/$AZURE_ASSIGNMENT_UUID",
  "name": "$AZURE_ASSIGNMENT_UUID",
  "principalId": "$AZURE_MANAGED_IDENTITY_SP_ID",
  "principalType": "ServicePrincipal",
  "resourceGroup": "$AZURE_RESOURCE_GROUP",
  "roleDefinitionId": "/subscriptions/$AZURE_SUBSCRIPTION_ID/providers/Microsoft.Authorization/roleDefinitions/$UUID",
  "scope": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME",
  "type": "Microsoft.Authorization/roleAssignments",
  "updatedOn": "1970-01-01T00:00:00.00000+00:00"
}

az role assignment create \
--role "Key Vault Certificates Officer" \
--assignee $AZURE_MANAGED_IDENTITY_SP_ID \
--scope "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME"
```

```json
{
  "condition": null,
  "conditionVersion": null,
  "createdBy": null,
  "createdOn": "1970-01-01T00:00:00.00000+00:00",
  "delegatedManagedIdentityResourceId": null,
  "description": null,
  "id": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME/providers/Microsoft.Authorization/roleAssignments/$UUID",
  "name": "$UUID",
  "principalId": "$AZURE_MANAGED_IDENTITY_SP_ID",
  "principalType": "ServicePrincipal",
  "resourceGroup": "$AZURE_RESOURCE_GROUP",
  "roleDefinitionId": "/subscriptions/$AZURE_SUBSCRIPTION_ID/providers/Microsoft.Authorization/roleDefinitions/a4417e6f-fecd-4de8-b567-7b0420556985",
  "scope": "/subscriptions/$AZURE_SUBSCRIPTION_ID/resourceGroups/$AZURE_RESOURCE_GROUP/providers/Microsoft.KeyVault/vaults/$AZURE_KEYVAULT_NAME",
  "type": "Microsoft.Authorization/roleAssignments",
  "createdOn": "1970-01-01T00:00:00.00000+00:00"
}
```