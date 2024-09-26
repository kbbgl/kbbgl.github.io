# Network

There is an `etcdctl` command to interrogate the database and `calicoctl` to view more of how the network is configured. We can see Felix, which is the primary Calico agent on each machine. This agent, or daemon, is responsible for interface monitoring and management, route programming, ACL configuration and state reporting.

BIRD is a dynamic IP routing daemon used by Felix to read routing state and distribute that information to other nodes in the cluster. This allows a client to connect to any node, and eventually be connected to the workload on a container, even if not the node originally contacted.

https://kubernetes.io/docs/concepts/cluster-administration/networking/

A `Pod` is a group of co-located containers that share the same IP address. From a networking perspective, a pod can be seen as a virtual machine of physical hosts. The network needs to assign IP addresses to pods, and needs to provide traffic routes between all pods on any nodes.

The three main networking challenges to solve in a container orchestration system are:

- Coupled container-to-container communication (solved by the `Pod` concept).

- Pod-to-pod communication.

- External-to-pod communication (solved by the `Service`s concept, which we will discuss later).

Kubernetes expects the network configuration to enable pod-to-pod communication to be available; it will not do it for you.

### CNI Network Configuration File

To provide container networking, Kubernetes is standardizing on the Container Network Interface (**CNI**) specification.

With CNI, you can write a network configuration file:

```json
{
    "cniVersion": "0.2.0",
    "name": "mynet",
    "type": "bridge",
    "bridge": "cni0",
    "isGateway": true,
    "ipMasq": true,
    "ipam": {
        "type": "host-local",
        "subnet": "10.22.0.0/16",
        "routes": [
            { "dst": "0.0.0.0/0" }
             ]
     }
}
```

This configuration defines a standard Linux bridge named cni0, which will give out IP addresses in the subnet `10.22.0.0./16`. The bridge plugin will configure the network interfaces in the correct namespaces to define the container network properly.

### `Pod`-to-`Pod` Communication

While a CNI plugin can be used to configure the network of a pod and provide a single IP per pod, CNI does not help you with pod-to-pod communication across nodes.

The requirement from Kubernetes is the following:

- All pods can communicate with each other across nodes.
- All nodes can communicate with all pods.
- No Network Address Translation (NAT).

this can be achieved with a software defined overlay with solutions like:

- Weave
- Flannel
- Calico
- Romana.

### Interacting with `etcd`

Checking health:

```bash
kubectl -n kube-system exec -it etcd-master -- sh \ 
-c "ETCDCTL_API=3 \ 
ETCDCTL_CACERT=/etc/kubernetes/pki/etcd/ca.crt \ ETCDCTL_CERT=/etc/kubernetes/pki/etcd/server.crt \ ETCDCTL_KEY=/etc/kubernetes/pki/etcd/server.key \
etcdctl endpoint health" 

https://127.0.0.1:2379 is healthy: successfully committed proposal: took = 11.942936ms
```

Checking how many databases there are:

```bash
kubectl -n kube-system exec -it etcd-master -- sh -c \
"ETCDCTL_API=3 etcdctl --cert=./peer.crt --key=./peer.key --cacert=./ca.crt \
endpoints=https://127.0.0.1:2379 member list"

1fb50b7ddbf4930ba, started, master, https://10.128.0.35:2380,2https://10.128.0.35:2379, false
```

Creating a backup:

```bash
kubectl -n kube-system exec -it etcd-master -- sh -c "ETCDCTL_API=3 \
ETCDCTL_CACERT=/etc/kubernetes/pki/etcd/ca.crt ETCDCTL_CERT=/etc/kubernetes/pki/etcd/server.crt \
ETCDCTL_KEY=/etc/kubernetes/pki/etcd/server.key  
etcdctl --endpoints=https://127.0.0.1:2379 snapshot save /var/lib/etcd/snapshot.db1

{"level":"info","ts":1598380941.6584022,"caller":"snapshot/v3_snapshot.go:110","msg":"created temporary db file","path":"/var/lib/etcd/snapshot.db.part"}{"level":"warn","ts":"2020-08-25T18:42:21.671Z","caller":"clientv3/retry_interceptor.go4:116","msg":"retry stream intercept"}{"level":"info","ts":1598380941.6736135,"caller":"snapshot/v3_snapshot.go:121","msg":"fetching snapshot","endpoint":"https://127.0.0.1:2379"}{"level":"info","ts":1598380941.7519674,"caller":"snapshot/v3_snapshot.go:134","msg":"fetched snapshot","endpoint":"https://127.0.0.1:2379","took":0.093466104}9{"level":"info","ts":1598380941.7521122,"caller":"snapshot/v3_snapshot.go:143","10msg":"saved","path":"/var/lib/etcd/snapshot.db"}
Snapshot saved at /var/lib/etcd/snapshot.db7.  Verify the snapshot exists from the node perspective, the file date should have been moments earlier.V2020-10-19
```

More details [here](https://drive.google.com/file/d/1ORuz5XftUmrMCSmN3Rl6okGu_ukaQlkZ/view?usp=sharing)
