# Get List of Containers in Pod

```bash
kubectl get pods $POD_NAME -n $NAMESPACE -o jsonpath='{.spec.containers[*].name}*
```
