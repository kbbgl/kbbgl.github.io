## Nodes

Nodes are physical or VM servers where the cluster is deployed. Nodes have roles, either master or worker. A `Node` is an API object created outside the cluster representing an instance.

We can create a master node with:

```bash
kubeadm init
``` 
and worker nodes:
```bash
kubeadm join
```

To remove nodes:

```bash
kubectl delete node $nodename

# remove cluster-specific information
kubeadm reset 
```

To view current node status:
```bash
kubectl describe node $nodename
```

All nodes run the following components:

- `kubelet`: Ensures that containers are running in a `Pod`.
    - Receives resource specifications (`PodSpec`) and ensures that the node meets the desired specs working with `container-runtime`
    - Ensures that a `Pod` has access to specified storage, `Secrets` or `ConfigMaps`.
    - Reports status of `Pod`s to cluster.
- `kube-proxy`: Network proxy implemented as a `Service`. It maintains network rules on the node based on IP tables so that internal and external clients can communicate with the `Pod`s.
- `container-runtime`: Works with `kubelet` to ensure specified containers are running and healthy. It's an interface between Kubernetes and the container/Docker Engine.

## Master Node

The Master node runs various management services for the whole cluster, mainly:

- `kube-apiserver`: All calls (internal and external traffic) are handled with this agent. It is the only connection to the `etcd` database. Acts as a master process for the entire cluster.
- `kube-scheduler`: Responsible for assigning Pods to Nodes by provided specification (labels, taints, toletations, number of replicas, state)
- `etcd`: key-value database where the state of the cluster information, as well as other persistent data and networking is kept. Requests to update DB is load-balanced by `kube-apiserver` and distributed to database in series. `etcdctl` can be used to interact with db.
- `kube-controller-manager`: Determines the state of the cluster by interaction with `kube-apiserver`. If any action is needed, it alerts other controllers (endpoints, namespaces, replicasets) of need to add/remove resources.

### Worker Node

The Worker Node hosts the applications workload.

