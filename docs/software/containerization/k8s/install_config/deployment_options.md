## Deployment Options

To help choose the type of deployment for use case, see https://kubernetes.io/docs/setup/.

The 4 main types are:

1) Single-node
2) Single head node, multiple workers
3) Multiple head nodes with HA, multiple workers
4) HA etcd, HA head nodes, multiple workers


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

## Install from Source
The list of binary releases is available on GitHub. Together with `gcloud`, `minikube`, and `kubeadmin`, these cover several scenarios to get started with Kubernetes.

Kubernetes can also be compiled from source relatively quickly. You can clone the repository from GitHub, and then use the Makefile to build the binaries. You can build them natively on your platform if you have a Golang environment properly set up, or via Docker containers if you are on a Docker host.

To build natively with Golang, first install Golang. Download files and directions can be found online.

Once Golang is working, you can clone the kubernetes repository, around 500MB in size. Change into the directory and use make:
```bash
$ cd $GOPATH

$ git clone https://github.com/kubernetes/kubernetes

$ cd kubernetes

$ make
```
On a Docker host, clone the repository anywhere you want and use the make quick-release command. The build will be done in Docker containers. 

The `_output/bin` directory will contain the newly built binaries.