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