## Upgrade Cluster

If the cluster was built using `kubeadm`, `kubeadm upgrade` can be used to upgrade a cluster.

To compare the installed version with the latest and if it's posssible to upgrade:
```bash
sudo kubeadm upgrade plan
```

To upgrade first control plane to specified version:
```bash 
sudo kubeadm apply
```

To show the differences applied during the upgrade (similar to `kubectl apply --dry-run`:
```bash
sudo kubeadm diff
```
This allows for updating the local `kubelet` configuration on worker nodes, or the control planes of other master nodes if there is more than one. Also, it will access a phase command to step through the upgrade process:
```bash
sudo kubeadm node
```

General upgrade process:

1) Update the software
2) Check the software version
3) Drain the control plane
4) View the planned upgrade
5) Apply the upgrade
6) Uncordon the control plane to allow pods to be scheduled.

[Full details](https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/)

