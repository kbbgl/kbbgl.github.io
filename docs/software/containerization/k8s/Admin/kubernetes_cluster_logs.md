# Log locations

## Master

`/var/log/kube-apiserver.log` - API Server, responsible for serving the API

`/var/log/kube-scheduler.log` - Scheduler, responsible for making scheduling decisions

`/var/log/kube-controller-manager.log` - Controller that manages replication controllers

## Worker

`/var/log/kubelet.log` - Kubelet, responsible for running containers on the node

`/var/log/kube-proxy.log` - Kube Proxy, responsible for service load balancing

`api-server` log location:

```bash
sudo find / -name "*apiserver*.log
/var/log/containers/kube-apiserver-master_kube-system_kube-apiserver-4232d25701998f68b503e64d41dd786e657fc09504f13278044934d79a4019e3c.log
```

## Container logs

```bash
ls /var/log/containers
```

`Pod` logs:

```bash
ls /var/log/pods
```

```bash
journalctl -xe -u kubelet --no-pager
```

Check cluster status (deprecated)

```bash
kubectl get cs
NAME                 STATUS    MESSAGE             ERROR
controller-manager   Healthy   ok                  
scheduler            Healthy   ok                  
etcd-2               Healthy   {"health":"true"}   
etcd-1               Healthy   {"health":"true"}   
etcd-0               Healthy   {"health":"true"} 
```

Check `kubelet` service status and logs:

```bash
systemctl status kubelet --no-pager
```

Run `kubelet` to check system calls

```bash
strace /usr/local/bin/kubelet
```
