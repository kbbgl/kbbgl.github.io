# Accessing API using `curl`

```bash
# kubernetes server
k8s_server=$(kubectl config view | grep server | cut -d":" -f2,3,4 | tr -d [:blank:])

# retrieve token (option 1)
token=$(kubectl get secrets -o jsonpath="{.items[?(@.metadata.annotations['kubernetes\.io/service-account\.name']=='default')].data.token}"|base64 --decode)

# retrieve token (option 2)
token=$(kubectl exec firstpod -- cat /var/run/secrets/kubernetes.io/serviceaccount/token)

# send request
curl $k8s_server/apis -H "Authorization: Bearer $token" -k
curl $k8s_server/api/v1 -H "Authorization: Bearer $token" -k
```
