## Labels

Part of the metadata of an object is a label. Though labels are not API objects, they are an important tool for cluster administration. They can be used to select an object based on an arbitrary string, regardless of the object type. Labels are immutable as of API version `apps/v1`.

Every resource can contain labels in its metadata. By default, creating a Deployment with `kubectl create` adds a label.

View labels in new columns: 

```bash
kubectl get pods -l run=ghost
NAME                    READY  STATUS   RESTARTS  AGE
ghost-3378155678-eq5i6  1/1    Running  0         10m

kubectl get pods -L run
NAME                    READY  STATUS   RESTARTS  AGE  RUN
ghost-3378155678-eq5i6  1/1    Running  0         10m  ghost
nginx-3771699605-4v27e  1/1    Running  1         1h   nginx
```

To add a label:
```bash
kubectl label pods ghost-3378155678-eq5i6 foo=bar

kubectl get pods --show-labels
NAME                    READY  STATUS   RESTARTS  AGE  LABELS
ghost-3378155678-eq5i6  1/1    Running  0         11m  foo=bar, pod-template-hash=3378155678,run=ghost
```

For example, if you want to force the scheduling of a pod on a specific node, you can use a `nodeSelector` in a `Pod` definition, add specific labels to certain nodes in your cluster and use those labels in the pod. 

```yaml
....
spec:
    containers:
    - image: nginx
    nodeSelector:
        disktype: ssd
```