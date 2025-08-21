---
slug: hashicorp-vault-cheatsheet
title: HashiCorp Vault Cheat Sheet
authors: [kgal-akl]
tags: [secrets, vault, hashicorp, cheatsheet, cli]
---

## Initialize Vault

After the HashiCorp Vault service deployment is complete, we need to initialize it:

```bash
vault operator init \
	-key-shares=1 \
	-key-threshold=1 \
	-format=json > vault-keys.json
```

The resulting `vault-keys.json` looks like:

```json
{
	"unseal_keys_b64": ["[REDACTED]+fQpVFVHQdYi9woCL1TYLY="],
	"unseal_keys_hex": ["7131540f02[REDACTED]"],
	"unseal_shares": 1,
	"unseal_threshold": 1,
	"recovery_keys_b64": [],
	"recovery_keys_hex": [],
	"recovery_keys_shares": 0,
	"recovery_keys_threshold": 0,
	"root_token": "hvs.[REDACTED]"
}
```

## Unseal Vault

To do anything (e.g. create key-vault store, create policies) with the vault, we need to first unseal it

```bash
UNSEAL_KEY=$(jq -r '.unseal_keys_b64[0]' vault-keys.json)

vault operator unseal $UNSEAL_KEY
```

We can confirm it's unsealed:

```bash
 vault status
Key             Value
---             -----
Seal Type       shamir
Initialized     true
Sealed          false
```

## Log Into Vault

We can use the root token to log in and manage the vault:

```bash
ROOT_TOKEN=$(jq -r '.root_token' vault-keys.json)

vault login $ROOT_TOKEN
```

Once we're logged in, we can [Enable Key/Value Store](#enable-key-value-store), [Add Secrets to Key/Value Store](#add-secrets-to-key-value-store), [Create Token](#create-token).

## Enable Key/Value Store

```bash
vault secrets enable -version=2 kv
```

## Add Secrets to Key/Value Store

```bash
vault kv put myservice/api api_key=t-1234 api_secret=null endpoint=https://api.staging.myservice.dev

vault kv put myservice/database username=admin password=securepassword123 host=staging.myservice.dev port=5432

vault kv list myservice
Keys
----
api
database
```

## Create Policy

We can use a policy to define what capabilities users have for certain vault paths. Here's an example:
```hcl
# Allow reading all secrets
path "myservice/*" {
  capabilities = ["read", "list"]
}

# Allow reading auth methods
path "auth/*" {
  capabilities = ["read", "list"]
}

# Allow reading sys/mounts to see enabled secrets engines
path "sys/mounts" {
  capabilities = ["read"]
}
```

To create the policy:
```bash
vault policy write myservice-migration-policy migration.hcl
```


## Create Token

Once we have a policy, we can create a token that will have the capabilities defined in the policy:

```bash
vault token create \
	-policy=myservice-migration-policy \
	-ttl=24h \
	-display-name="myservice-migration-token"
```

```
Key                  Value
---                  -----
token                hvs.[REDACTED]
token_accessor       [REDACTED]
token_duration       768h
token_renewable      true
token_policies       ["default" "myservice-migration-policy"]
identity_policies    []
policies             ["default" "myservice-migration-policy"]
```


## Accessing Vault UI

When the vault is deployed in Kubernetes, we can access the vault UI by navigating to http://localhost:8200 after running:

```bash
VAULT_POD=$(kubectl get pods -n vault -l app.kubernetes.io/name=vault -o jsonpath='{.items[0].metadata.name}')
kubectl port-forward $VAULT_POD 8200:8200 -n vault
```

Log in with the token created in [Create Token](#create-token).
