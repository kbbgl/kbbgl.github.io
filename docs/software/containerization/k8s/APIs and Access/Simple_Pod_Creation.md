## Simple Pod Creation

A simple `Pod` manifest will look like this:

```yaml
apiVersion: v1
kind: Pod
metadata:
    name: firstpod
spec:
    containers:
    - image: nginx
      name: its_me
```

To create the `Pod`:

```bash
kubectl create -f simple.yaml
```
