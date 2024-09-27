## Managing API resources with `kubectl`

Kubernetes exposes resources via RESTful API calls.

We can see the API call using `curl` by setting highest verbosity:

```bash
kubectl --v=10 get pods firstpod | grep curl


curl -k -v -XGET  -H "User-Agent: kubectl/v1.20.2 (linux/arm64) kubernetes/faecb19" -H "Accept: application/json;as=Table;v=v1;g=meta.k8s.io,application/json;as=Table;v=v1beta1;g=meta.k8s.io,application/json" -H "Authorization: Bearer <masked>" 'https://10.100.102.95:16443/api/v1/namespaces/test/pods/firstpod'
```

We can see the server information is stored in `/home/student/.kube/config`:

```bash
kubectl --v=10 config view

I0124 19:26:56.264075 3256697 loader.go:379] Config loaded from file:  /home/ubuntu/.kube/config
```

```yaml
apiVersion: v1 # instructs `kube-apiserver` which API version to use to retrieve this resource, e.g https://host:16443/api/${apiVersion}/namespace...
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED # passed to authenticate `curl`, `kubectl` requests
    server: https://10.100.102.95:16443
  name: microk8s-cluster
contexts: # setting used to configure different clusters/namespaces
- context:
    cluster: microk8s-cluster
    namespace: test
    user: admin
  name: microk8s
current-context: microk8s # current used context 
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: REDACTED
```

### Configuring TLS Access

To set up access to the cluster resources with `curl`:

```bash
# get kubernetes server
k8s_server=$(kubectl config view | grep server | cut -d":" -f2,3,4 | tr -d [:blank:])

export client=$(grep client-cert ~/.kube/config | cut -d" " -f 6)
export key=$(grep client-key-data ~/.kube/config | cut -d" " -f 6)
export auth=$(grep certificate-authority-data ~/.kube/config | cut -d" " -f 6)
# encode keys for use with `curl`
echo  $client | base64 -d - > client.pem
echo $key | base64 -d - > ./client-key.pem
echo $auth | base64 -d -> ./ca.pem

# find server ifno
kubectl config view | grep https | cut -d":" -f2,3,4

# get all Pods status in JSON from API
curl --cert ./client.pem --key ./client-key.pem --cacert ./ca.pem $k8s_server/api/v1/pods

# create a Pod from a file
curl --cert ./client.pem --key ./client-key.pem --cacert ./ca.pem 
-X POST -H "Content-Type: application/json" -d@pod.yaml $k8s_server/api/v1/namespaces/default/pods
```
