## Scaling and Rolling Updates

To scale a deployment:
```bash
kubectl scale deploy/dev-web --replicas=4
```

To trigger a rolling update, we can change the image of a container:
```bash
kubectl edit deployment nginx

...
containers:
  - image: nginx: 1.8 # change to older value
...
```

### Rollbacks

using the `--record` option of the `kubectl create` command, which allows annotation in the resource definition.

```bash
kubectl create deploy ghost --image=ghost --record

$ kubectl get deployments ghost -o yaml
deployment.kubernetes.io/revision: "1" 
kubernetes.io/change-cause: kubectl create deploy ghost --image=ghost --record
```

Should an update fail, due to an improper image version, for example, you can roll back the change to a working version with `kubectl rollout` undo:

```bash
kubectl set image deployment/ghost ghost=ghost:09 --all

kubectl rollout history deployment/ghost deployments "ghost":
REVISION   CHANGE-CAUSE
1 ​         kubectl create deploy ghost --image=ghost --record
2          kubectl set image deployment/ghost ghost=ghost:09 --all

kubectl get pods
NAME                    READY  STATUS            RESTARTS  AGE
ghost-2141819201-tcths  0/1    ImagePullBackOff  0         1m​

kubectl rollout undo deployment/ghost ; kubectl get pods

NAME                    READY  STATUS   RESTARTS  AGE
ghost-3378155678-eq5i6  1/1    Running  0         7s
```

You can roll back to a specific revision with the `--to-revision=2` option.

You can also edit a Deployment using the `kubectl edit` command.

You can also pause a Deployment, and then resume.

```bash
kubectl rollout pause deployment/ghost

kubectl rollout resume deployment/ghost
```

Can still do a rolling update on `ReplicationControllers` with the `kubectl rolling-update` command, but this is done on the client side. Hence, if you close your client, the rolling update will stop.