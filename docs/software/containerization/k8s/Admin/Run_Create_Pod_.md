# Run and Create Pod

## Best way

```bash
kubectl run  backup-pod --rm -i --tty --image=quay.io/app/utils:L8.0.5.163 -- bash 
```

It would be deleted when you exit bash, if you want it to stay up remove the `--rm`  and then do

```bash
kubectl attach backup-pod -c backup-pod -i -t 
```

to renter the pod

### Create

```bash

kubectl run -i -n app  --tty app-backup-debug --image="quay.io/app/utils:L8.0.4.244" -- sh
```

### Resume session

```bash
kubectl attach app-backup-debug-68776c496b-l5pbc -c app-backup-debug -i -t
```
