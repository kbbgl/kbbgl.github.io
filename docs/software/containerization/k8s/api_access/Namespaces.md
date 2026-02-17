## Namespaces

Namespaces refer to both the Linux kernel feature and the segregation of API objects by Kubernetes.

Every API call includes a namespace:

```
https://10.128.0.3:6443/api/v1/namespaces/default/pods
```

Namespaces are intended to isolate multuple groups and the resources they have access to work with via quotas.

There are 4 namespaces when the cluster is first created:

- `default`: Default namespace where all resources are assumed to be stored if not specified.

- `kube-node-lease`: Where worker node lease information is kept.

- `kube-public`: A namespace readable by all, even those not authenticated.

- `kube-system`: Contains infrastructure pods.

To specify a resource's namespace:

```yaml
# redis.yaml
apiVersion: V1
kind: Pod
metadata:
  name: redis
  namespace: linuxcon
...
```
