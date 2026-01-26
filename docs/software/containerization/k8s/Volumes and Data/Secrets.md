## Secrets

Some data should not be read and can be encrypted/encoded using the `Secret` API.

To create a `Secret`:

```bash
kubectl create secret generic --help

kubectl create secret generic mysql --from-literal=password=root
```

To create an encoded `Secret` manually:

```bash
echo safyryaifa | base64
c2FmeXJ5YWlmYQo=
```

```yaml
# secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
data:
  password: c2FmeXJ5YWlmYQo=
```

By default, `Secret`s are base-64 encoded. To encrypt, an `EncryptionConfiguration` resource must be created with a key and identity.


### Copy Secrets Between Namespaces

```bash
#!/bin/bash

# Usage: ./copy_secret.sh <secret_name> <source_ns> <dest_ns>

SECRET_NAME=$1
SOURCE_NS=$2
DEST_NS=$3

if [[ -z "$SECRET_NAME" || -z "$SOURCE_NS" || -z "$DEST_NS" ]]; then
    echo "Usage: $0 <secret-name> <source-namespace> <destination-namespace>"
    exit 1
fi

echo "Copying secret '$SECRET_NAME' from '$SOURCE_NS' to '$DEST_NS'..."

# 1. Get the secret as JSON
# 2. Strip out system-generated metadata (uid, resourceVersion, etc.)
# 3. Change the namespace field
# 4. Apply it to the new namespace
kubectl get secret "$SECRET_NAME" --namespace="$SOURCE_NS" -o json | \
jq 'del(.metadata.namespace,.metadata.resourceVersion,.metadata.uid,.metadata.creationTimestamp,.metadata.selfLink,.metadata.managedFields)' | \
kubectl apply --namespace="$DEST_NS" -f -

if [ $? -eq 0 ]; then
    echo "Success: Secret '$SECRET_NAME' is now in '$DEST_NS'"
else
    echo "Error: Failed to copy secret."
fi
```