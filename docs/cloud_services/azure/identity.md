---
title: Azure Identity
slug: azure-identity
app2or: kgal-akl
tags: [devops, azure, aks, az, imds, metadata, identity]
---

There are 2 authentication mechanisms available in Azure: Azure IMDS (e.g. Managed Identity) and Azure Workload Identity.

The fundamental difference is where the trust originates. IMDS trusts the physical/virtual network of the Azure host, while Workload Identity trusts the Kubernetes identity provider.

## IMDS

IMDS primary use case is using virtual machines, app services and functions. 

The trust source is the host network which is only accessible within the resource's network stack.

The flow is: Your code -> Local IP -> Azure Fabric -> Entra ID

The protocol is HTTP with a call to the IMDS endpoint available in the Azure resource, `169.254.169.254`:

```bash
curl -H "Metadata: true" \
"http://169.254.169.254/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https://management.azure.com/"
```

```powershell
Invoke-RestMethod -Headers @{"Metadata"="true"} -Method GET -Uri "http://169.254.169.254/metadata/instance?api-version=2025-04-07" | ConvertTo-Json -Depth 64
```

## Workload Identity

Workload Identity's primary use case is Kubernetes (AKS).

The trust source is OIDC federation, a relationship between Kubernetes and Entra ID. 

The flow is, workload/app -> K8s Service Account Token -> Entra ID (verifies via OIDC) -> Access Token.

The protocol is an OIDC token exchange, specifically a Kubernetes `ServiceAccount` token for the Entra token. Therefore, to enable Azure Workload Identity, the infrastructure requires an OIDC issuer and a `MutatingWebhook` on Kubernetes. 

To retrieve a token, the application (workload) must first have a signed Kubernetes token (projected into the `Pod`, usually placed in `/var/run/secrets/azure/tokens/azure-identity`). We can then call this endpoint to retrieve an Azure token:

```bash
curl -X POST https://login.microsoftonline.com/$AZURE_TENANT_ID/oauth2/v2.0/token \
-d "client_id=$AZURE_CLIENT_ID" \
-d "scope=https://management.azure.com/.default" \
-d "client_assertion_type=urn:ietf:params:oauth:client-assertion-type:jwt-bearer" \
-d "client_assertion=$(cat /var/run/secrets/azure/tokens/azure-identity)" \
-d "grant_type=client_credentials"
```

To check if the OIDC provider is enabled, see [Check Workload Identity](./aks.md#check-workload-identity).
To enable Workload Identity using the Azure CLI, see [Enable Workload Identity](./aks.md#enable-workload-identity).


## Configuration

### Environmental Variables

Azure credentials can be configured by environment variables. This credential is capable of authenticating as a service principal using a client secret or a certificate. Configuration is attempted in this order, using these environment variables:

#### Service principal with secret:

- `AZURE_TENANT_ID`: ID of the service principal's tenant. Also called its 'directory' ID.
- `AZURE_CLIENT_ID`: the service principal's client ID
- `AZURE_CLIENT_SECRET`: one of the service principal's client secrets
- `AZURE_AUTHORITY_HOST`: authority of a Microsoft Entra endpoint, for example "login.microsoftonline.com", the authority for Azure Public Cloud, which is the default when no value is given.

#### Service principal with certificate:

- `AZURE_TENANT_ID`: ID of the service principal's tenant. Also called its 'directory' ID.
- `AZURE_CLIENT_ID`: the service principal's client ID
- `AZURE_CLIENT_CERTIFICATE_PATH`: path to a PEM or PKCS12 certificate file including the private key.
- `AZURE_CLIENT_CERTIFICATE_PASSWORD`: (optional) password of the certificate file, if any.
- `AZURE_CLIENT_SEND_CERTIFICATE_CHAIN`: (optional) If True, the credential will send the public certificate chain in the x5c header of each token request's JWT. This is required for Subject Name/Issuer (SNI) authentication. Defaults to False.
- `AZURE_AUTHORITY_HOST`: authority of a Microsoft Entra endpoint, for example "login.microsoftonline.com", the authority for Azure Public Cloud, which is the default when no value is given.


## Authenticating using Azure Go SDK

To authenticate using the [Azure Go SDK](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go#section-readme), `azureidentity`, we can use the `DefaultAzureCredential` which automatically tries Workload Identity before falling back to IMDS.


```go
import (
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

func main() {
    // This will check Workload Identity first, then IMDS (Managed Identity)
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        panic(err)
    }

    // Use the credential with any Azure client
    client, _ := armresources.NewResourceGroupsClient("<subscription-id>", cred, nil)
}
```

If we want to explicitly target IMDS/Managed Identity:

```go
managedCred, err := azidentity.NewManagedIdentityCredential(nil)
```

Otherwise, if we want to explicitly target Workload Identity:

```go
// It reads AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_FEDERATED_TOKEN_FILE from environment
workloadCred, err := azidentity.NewWorkloadIdentityCredential(nil)
```

### Sources

- https://learn.microsoft.com/en-us/python/api/azure-identity/azure.identity.environmentcredential?view=azure-python