## Installation Tools

```
kubespray
kops
kube-aws
kubicorn
```


### Manual Installation
To install Kubernetes manually, see https://github.com/kelseyhightower/kubernetes-the-hard-way.

Kubernetes is a set of daemons/binaries:

- `kube-apiserver` (AKA the `master`),
- `kubelet` (start/stop containers, sync conf.),
- `kube-scheduler` (resources manager)
- `kube-controller-manager` (monitor RC, and maintain the desired state)
- `kube-proxy` (expose services on each node)
- `kubectl` (CLI)

The `hyperkube` binary is an all in one binary (in a way similar to `busybox`), combining all the previously separate binaries.

The following command:

```bash
hyperkube kubelet \
  --api-servers=http://localhost:8080 \
  --v=2 \
  --address=0.0.0.0 \
  --enable-server \
  --hostname-override=127.0.0.1 \
  --config=/etc/kubernetes/manifests-multi \
  --cluster-dns=10.0.0.10 \
  --cluster-domain=cluster.local
```
runs the daemon `kubelet`.