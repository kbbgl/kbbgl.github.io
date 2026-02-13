---
slug: managing-azure-container-app
title: Managing Azure Container Apps using in Azure CLI
authors: [kgal-akl]
tags: [microsoft, azure, cloud, user, managed_identity, federation, az, cli, ad, entra_id]
---

## Deploy/Upgrade a Container App using Bicep

To deploy or upgrade an Azure Container App it using a Bicep template and parameters:

```bash
az deployment group create \
--resource-group "$AZURE_RESOURCE_GROUP" \
--template-file main.bicep \
--parameters main.parameters.json \

# OR
# --parameters \
    # location="$AZURE_LOCATION" \
    # acaEnvName="$ACA_ENV_NAME" \
    # acaAppName="$ACA_APP_NAME" \
    # uamiName="$ACA_UAMI_NAME" \
    # parameter1="$PARAM_1"
```

`main.parameters.json` uses [`deployParameters` JSON schema](https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json). Here's a sample:

```json
{
    "$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentParameters.json#",
    "contentVersion": "1.0.0.0",
    "parameters": {
      "location": { "value": "eastus" },
      "acaEnvName": { "value": "kgal-aca-wid-env" },
      "acaAppName": { "value": "kgal-aca-wid-gw" },
      "uamiName": { "value": "kgal-aca-wid-uami" },
      "containerImage": { "value": "docker.registry.kbbgl.dev/app:tag" },
      "appUrl": { "value": "https://kbbgl.github.io/app" },
      "cpu": { "value": "2.0" },
      "memory": { "value": "4Gi" }
    }
  }
```

And `main.bicep`:

```
targetScope = 'resourceGroup'

@description('Azure location')
param location string = resourceGroup().location

@description('Container Apps environment name')
param acaEnvName string

@description('Container App name')
param acaAppName string

@description('User assigned managed identity name')
param uamiName string

@description('Container image')
param containerImage string = 'docker.registry.kbbgl.dev/app:tag'

@description('App URL')
param appUrl string = 'https://app.kbbgl.github.io'

@description('CPU cores')
param cpu string = '2.0'

@description('Memory')
param memory string = '4Gi'

resource logAnalytics 'Microsoft.OperationalInsights/workspaces@2022-10-01' = {
  name: '${acaEnvName}-law'
  location: location
  properties: {
    sku: {
      name: 'PerGB2018'
    }
    retentionInDays: 30
  }
}

resource uami 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' = {
  name: uamiName
  location: location
}

resource acaEnv 'Microsoft.App/managedEnvironments@2024-03-01' = {
  name: acaEnvName
  location: location
  properties: {
    appLogsConfiguration: {
      destination: 'log-analytics'
      logAnalyticsConfiguration: {
        customerId: logAnalytics.properties.customerId
        sharedKey: listKeys(logAnalytics.id, '2022-10-01').primarySharedKey
      }
    }
  }
}

resource acaApp 'Microsoft.App/containerApps@2024-03-01' = {
  name: acaAppName
  location: location
  identity: {
    type: 'UserAssigned'
    userAssignedIdentities: {
      '${uami.id}': {}
    }
  }
  properties: {
    managedEnvironmentId: acaEnv.id
    configuration: {
      ingress: {
        external: true
        targetPort: 80
        transport: 'auto'
      }
    }
    template: {
      containers: [
        {
          name: 'app'
          image: containerImage
          env: [
            { name: 'APP_URL', value: appUrl }
          ]
          resources: {
            cpu: json(cpu)
            memory: memory
          }
        }
      ]
      scale: {
        minReplicas: 1
        maxReplicas: 1
      }
    }
  }
}

output uamiResourceId string = uami.id
output uamiClientId string = uami.properties.clientId
output acaFqdn string = acaApp.properties.configuration.ingress.fqdn
output acaResourceId string = acaApp.id
```


## Show Container App

```bash
az containerapp show \
--name "$ACA_APP_NAME" \
--resource-group "$AZURE_RESOURCE_GROUP" \
--output yaml
```

## Show Container App Logs

```bash
az containerapp logs show \
--name "$ACA_APP_NAME" \
--resource-group "$AZURE_RESOURCE_GROUP" \
--follow
```

## Access Container App

```bash
az containerapp exec \
--name "$ACA_APP_NAME" \
--resource-group "$AZURE_RESOURCE_GROUP" \
--container app \
--command "/bin/sh"
```

