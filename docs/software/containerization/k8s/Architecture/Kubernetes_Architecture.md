## Kubernetes Architecture

Kubernetes main components:

- Master/worker nodes
- Controllers
- Services
- Pods of containers
- Namespaces and quotas
- Network and policies
- Storage

A Kubernetes cluster is composed of a master node and a set of worker nodes. The cluster is all driven via API calls to controllers (interior/exterior traffic).

![k8s-architecture](https://d36ai2hkxl16us.cloudfront.net/course-uploads/e0df7fbf-a057-42af-8a1f-590912be5460/ww1wo482ah07-Kubernetes_Architecture2.png)

### Example, request to create a `Pod`

1. User request to create a new `Pod` received by `kube-apiserver`.

1. `kube-apiserver` queries `etcd` for number of current `Pod`s.

1. `etcd` responds to `kube-apiserver` with number of current `Pod`s.

1. `kube-apiserver` sends request to `kube-controller-manager` to create a new `Pod`.

1. `kube-controller-manager` compares current and desired states and responds to `kube-apiserver` with need to spawn a new `Pod`.

1. `kube-apiserver` sends request to `kube-scheduler` to schedule a new `Pod`. `kube-scheduler` responds to `kube-apiserver` with specific worker node where the `Pod` needs to be scheduled.

1. `kube-apiserver` sends request to `kubelet` on Worker Node and to `kube-proxy` to change network configuration/routing (`iptables`, `ipvs`)

1. `kubelet` communicates with `container engine` which communicates with the containers. A `Service` will expose the containers externally.

1. `kubelet` and `kube-proxy` respond to `kube-apiserver` with changes made and current state.
