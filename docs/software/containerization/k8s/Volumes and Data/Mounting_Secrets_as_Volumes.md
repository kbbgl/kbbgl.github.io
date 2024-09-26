## Mounting `Secret`s as `Volume`s

You can also mount secrets as files using a volume definition in a pod manifest. The mount path will contain a file whose name will be the key of the secret created with the kubectl create secret step earlier.

```yaml
...
spec:
    containers:
    - image: busybox
      command:
        - sleep
        - "3600"
      volumeMounts:
      - mountPath: /mysqlpassword
        name: mysql
      name: busy
    volumes:
    - name: mysql
        secret:
            secretName: mysql
```
Once the pod is running, you can verify that the secret is indeed accessible in the container:
```bash
kubectl exec -ti busybox -- cat /mysqlpassword/password
LFTr@1n
```