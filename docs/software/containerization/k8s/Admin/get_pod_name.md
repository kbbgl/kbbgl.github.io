# Get Pod Name

```bash
kubectl get pods -n app -l app=management --no-headers -o custom-columns=":metadata.name"
```
