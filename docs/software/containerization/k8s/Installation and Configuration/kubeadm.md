## `kubeadm`

`kubeadm` is used to build clusters.

[Full process to create a cluster](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/)

The basic steps to join other nodes to the cluster:

1) Retrieve a token and a SHA256 hash returned by:

```bash
# on to be master node
kubeadm init
```

2) Once the master is initialized, we would apply a network plugin (e.g. `calico`).

3) Create a network for IP-per-Pod criteria. For example, using Weave:

```bash
kubectl create -f https://git.io/weave-kube
```

4) Run the following command on the worker nodes:

```bash
# on worker node
kubeadm join --token $token $master_node_ip
```

