## Accessing API using Proxy

Another way to interact with the API is via a proxy. The proxy can be run from a node or from within a Pod through the use of a sidecar.

```bash
# start a proxy in the background for the whole Kubernetes API
kubectl proxy --api-prefix=/ &

# [1] 28633
# Starting to serve on 127.0.0.1:8001

curl http://127.0.0.1:8001/apis
```