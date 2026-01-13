---
slug: akeyless-cli-cheatsheet
title: Akeyless CLI Cheatsheet
description: Akeyless CLI Cheatsheet
authors: [kgal-akl]
tags: [akeyless, cli, cheatsheet]
---

## Authentication
### SAML

```bash
akeyless auth --access-id $SAML_ACCESS_ID --access-type saml
```

### AWS IAM

```bash
akeyless auth --access-id $AWS_IAM_ACCESS_ID --access-type aws_iam --debug
```


### GCP IAM/SA
```bash
GOOGLE_APPLICATION_CREDENTIALS=/path/to/gcp/app.json 
akeyless auth --access-id $GCP_IAM_ACCESS_ID --access-type gcp
```


### Azure

#### Create Auth Method

Get the `AZURE_TENANT_ID` by [retreiving the Tenant ID](../../../cloud_services/azure/account.md#retrieve-tenant-id).
```bash
AZURE_TENANT_ID=$(az account show \
--query tenantId \
--output tsv)

akeyless create-auth-method-azure-ad \
--name "$AUTH_METHOD_NAME" \
--bound-tenant-id "$AZURE_TENANT_ID" \
--profile kgal-azure-dev-gw-admin;
```

Make sure to [enable a managed idenitity](../../../cloud_services/azure/vm.md#assign-a-managed-identity-to-vm).

Then authenticate:

```bash
akeyless auth \
--access-id $AZURE_AD_AKEYLESS_ACCESS_ID \
--access-type azure_ad
```

#### Get JWT from T-Token

```bash
akeyless validate-token --token $AKEYLESS_T_TOKEN --debug
```

#### Get Cloud Identity

We can run this within a cloud provider environment to retrieve the cloud ID and use it to authenticate to [AWS](./cli.md#aws-iam), [Azure](./cli.md#azure) or [GCP](./cli.md#gcp-iamsa)
```bash
akeyless get-cloud-identity --describe-sub-claims
akeyless auth gcp --access-id $GCP_IAM_ACCESS_ID --cloud-id $CLOUD_ID
```

### LDAP

```bash
akeyless auth \
--username "$LDAP_USERNAME" \
--password "$LDAP_PASSWORD" \
--access-type ldap \
--access-id $LDAP_ACCESS_ID \
--ldap_proxy_url http://localhost:8000
```


### Authenticate using Universal ID

```bash
AKEYLESS_GATEWAY_URL=https://kgal-ec2.dev/api/v1
akeyless uid-generate-token \
--auth-method-name "$AUTH_METHOD_NAME"\
--profile $AKEYLESS_PROFILE
```

### Kubernetes

If we're authenticating from within a cluster (e.g inside an Akeyless Gateway container)
```bash
akeyless auth \
--access-type k8s \
--access-id "$ACCESS_ID" \
--k8s-auth-config-name "$K8S_AUTH_CONF_NAME" \
--gateway-url "http://localhost:8000"
```

Or using the service account JWT when authenticating outside the cluster:

```bash
B64_ENCODED_JWT=(echo "eJ..." | base64 | tr -d '\n')
 akeyless auth \
--access-type k8s \
--access-id "$ACCESS_ID" \
--k8s-auth-config-name "$K8S_AUTH_CONF_NAME" \
--k8s-service-account-token  "$B64_ENCODED_JWT" \
--gateway-url "http://localhost:8000"
```


## Configuration
### Use Gateway

To make calls to the Gateway, set:

```bash
AKEYLESS_GATEWAY_URL="$GATEWAY_FQDN/api/v1" 

akeyless get-secret-value --profile $MY_PROFILE --name $SECRET_NAME
```

### Define a Profile

```bash
AKEYLESS_PROFILE="global-staging-api-key"
akeyless configure \
--profile "$AKEYLESS_PROFILE" \
--access-id "$AKEYLESS_ACCESS_ID" \
--access-key "$AKEYLESS_ACCESS_KEY" \
--gateway-url "http://localhost:8000"

cat /Users/kgal/.akeyless/profiles/$AKEYLESS_PROFILE.toml
```

### Configure CLI to use SAML

```bash
akeyless configure --profile $AKEYLESS_PROFILE --access-id $SAML_ACCESS_ID --access-type saml
```


### Configure CLI to use AWS IAM

```bash
akeyless configure --profile $AKEYLESS_PROFILE --access-id $AWS_IAM_ACCESS_ID --access-type aws_iam
```

### Configure CLI to use different tenant

set the `dns` field in `~/.akeyless/settings`.

## Secrets
### Get Secret Value

```bash
akeyless get-secret-value --name "$SECRET_NAME"
```

### Create Secret

```bash
SECRET_NAME="test"
akeyless create-secret --name $SECRET_NAME --value noSecret --type generic

A new secret named /test was successfully created
```


### Delete Secret

```bash
akeyless delete-item --name $SECRET_NAME

Item /test was successfully deleted
```


### Create Google Workspace Dynamic Secret
```bash
GWORKSPACE_DS_NAME="test-gworkspace-ds"
akeyless dynamic-secret create google-workspace \
--name $GWORKSPACE_DS_NAME \
--target-name "$GWORKSPACE_DS_NAME-target" \
--access-mode role \
--admin-email kbbgl@github.io \
--role-name "$GCP_ROLE_NAME" \
--role-scope CUSTOMER \
--user-ttl 60m \
--profile $AKEYLESS_PROFILE \
--gateway-url http://localhost:8000

Dynamic secret test-gworkspace-ds successfully created
- Dynamic secret ID: 1
```

## SRA

### Connect to Target
```bash
akeyless connect \
--target $VM_USER@$TARGET_IP \
--via-sra 11.22.33.44 \
--cert-issuer-name /ssh-issuer-name \
--profile $AKEYLESS_PROFILE
```

### Connect to EKS Target
```bash
akeyless connect \
--target default@$EKS_CLUSTER_ID.$AWS_REGION.eks.amazonaws.com \
--name /eks/eks-ds \
--cert-issuer /8-ssh/issuer-compose \
--via-sra localhost:2222
--gateway-url http://localhost:8000 \
--profile $AKEYLESS_PROFILE
```

We can add `--debug` to see some details about the command that is being run to set up the connection.

## Certificates

### Create a Certificate
```bash
CERT_NAME=kbbgl-gh-dev-tls
akeyless create-certificate \
--name "$CERT_NAME" \
--certificate ./certs/kbbgl-gh-dev.crt \
--private-key ./certs/kbbgl-gh-dev.key \
--profile $AKEYLESS_PROFILE
```

### Get Certificate and Private Key
```bash
akeyless get-certificate-value \
--name "$CERT_NAME" \
--certificate-file-output "/tmp/$CERT_NAME-crt.pem" \
--private-key-file-output "/tmp/$CERT_NAME-key.pem" \
--profile $AKEYLESS_PROFILE
```

#### Targets

### Get Target Details

```bash
akeyless target get-details --name target-1 --profile $AKEYLESS_PROFILE
```
