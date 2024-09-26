# Exposing Service Outside Cluster

```bash
kubectl get po                                                   
NAME                     READY   STATUS    RESTARTS   AGE
nginx-7848d4b86f-kt9p6   1/1     Running   0          10m
nginx-7848d4b86f-rfvwk   1/1     Running   0          5m41s
nginx-7848d4b86f-5gzhc   1/1     Running   0          5m41s
```

Make sure that the `nginx` deployment has port set:

```bash
kubectl get deployments.apps nginx -o=yaml | grep -A2 ports
```

```yaml
ports:
    - containerPort: 80
      protocol: TCP
```

Expose service:

```bash
kubectl expose deployment nginx --type=LoadBalancer
```

Get service where service is listening on (32753):

```bash
kubectl get svc
NAME    TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
nginx   LoadBalancer   10.152.183.139   <pending>     80:32010/TCP   5s

hostname -i
10.100.102.95
```

Try to reach from machine outside k8s node:

```bash
curl 10.100.102.95 32010
```
