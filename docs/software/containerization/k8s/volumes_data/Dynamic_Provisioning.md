## Dynamic Provisioning

The `StorageClass` API allows an administrator to define a persistent volume provisioner of a certain type, passing storage-specific parameters. With a `StorageClass` created, a user can request a claim, which the API server fills via auto-provisioning. Common providers are AWS and GCE.

An example of GCE `StorageClass` specification:

```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
```
