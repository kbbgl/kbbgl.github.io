# Backup and Restore Deployment

```bash
# 1. backup deployment
kubectl get deployment nginx -o yaml > nginx.yaml

# 2. remove `creationTimestamp`, `resourceVersion`, `selfLink`, `uid`
kubectl get deployments.apps nginx -o yaml | grep -Ev "creationTimestamp|resourceVersion|selfLink|uid" > nginx_2.yaml

# 3. create resource
kubectl apply -f nginx_2.yaml
```
