---
title: Azure Identity
slug: azure-identity
app2or: kgal-akl
tags: [devops, azure, aks, az, imds, metadata, identity]
---

Azure credentials can be configured by environment variables. This credential is capable of authenticating as a service principal using a client secret or a certificate. Configuration is attempted in this order, using these environment variables:

## Service principal with secret:

- `AZURE_TENANT_ID`: ID of the service principal's tenant. Also called its 'directory' ID.
- `AZURE_CLIENT_ID`: the service principal's client ID
- `AZURE_CLIENT_SECRET`: one of the service principal's client secrets
- `AZURE_AUTHORITY_HOST`: authority of a Microsoft Entra endpoint, for example "login.microsoftonline.com", the authority for Azure Public Cloud, which is the default when no value is given.

## Service principal with certificate:

- `AZURE_TENANT_ID`: ID of the service principal's tenant. Also called its 'directory' ID.
- `AZURE_CLIENT_ID`: the service principal's client ID
- `AZURE_CLIENT_CERTIFICATE_PATH`: path to a PEM or PKCS12 certificate file including the private key.
- `AZURE_CLIENT_CERTIFICATE_PASSWORD`: (optional) password of the certificate file, if any.
- `AZURE_CLIENT_SEND_CERTIFICATE_CHAIN`: (optional) If True, the credential will send the public certificate chain in the x5c header of each token request's JWT. This is required for Subject Name/Issuer (SNI) authentication. Defaults to False.
- `AZURE_AUTHORITY_HOST`: authority of a Microsoft Entra endpoint, for example "login.microsoftonline.com", the authority for Azure Public Cloud, which is the default when no value is given.

### Sources

- https://learn.microsoft.com/en-us/python/api/azure-identity/azure.identity.environmentcredential?view=azure-python