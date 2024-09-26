```bash
kubectl get cm -n kube-system -oyaml cm-logmon-app-env
```

```yaml
apiVersion: v1
data:
  efk: "false"
  external_monitoring: "false"
  internal_monitoring: "true"
  app_ownerid: SOME_OWNER_ID
```
