## Persistent Volumes and Claims

### Persistent Volumes

A **persistent volume (pv)** is a storage abstraction used to retain data longer than the `Pod` using it. `Pod`s define a volume of type `persistentVolumeClaim (pvc)`. The cluster then attaches the `persistentVolume`. PVs are not namespaced while PVCs are.

The following are the phases of Persistent Storage:

1) `Provision`: Provisioning can be from the PVs created in advance by the cluster administrator or requested from a dynamic provider (Google/AWS, etc).

2) `Bind`: Binding occurs when a control loop on the master notices the PVC and locates a matching PV.

3) `Use`: When the bound volume is mounted for the `Pod` to use and is kept in this phase as long as the `Pod` requires.

4) `Release`: Happens when the `Pod` is done with the volume and an API request to delete the PVC is sent.

5) `Reclaim`: could be:
    - `Retain`: keeps data intact allowing an admin to handle storage and data.
    - `Delete`: tells volume plugin to delete the API object as well as storage.
    - `Recycle`: run `rm -rf /mountpoint` and then makes it available to a new claim.

We can see the PV and PVC:

```bash
kubectl get pv
kubectl get pvc
```

The following specification is a basic declaration of a PV using `hostPath`:

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: 10Gpv01
  labels:
    type: local
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/somepath/data01"
```

### Persistent Volume Claim

With a created PV in the cluster, we can write a manifest for a claim and use that claim in the `Pod` specification using `persistentVolumeClaim`.

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: myclaim
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 8Gi
```

In the `Pod`:

```yaml
spec:
  containers:
...
  volumes:
    - name: test-volume
      persistentVolumeClaim:
        claimName: myclain
```
