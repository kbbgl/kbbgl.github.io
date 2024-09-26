## API Groups

We can discover all API groups by running:

```bash
curl https://$(hostname -i):6443 \ 
--key ./client-key.pem \
--cacert ./ca.pem \
--cert ./client.pem \
jq '.groups[].name'

"apiregistration.k8s.io"
"extensions"
"apps"
"events.k8s.io"
"authentication.k8s.io"
"authorization.k8s.io"
"autoscaling"
"batch"
"certificates.k8s.io"
"networking.k8s.io"
"policy"
"rbac.authorization.k8s.io"
"storage.k8s.io"
"admissionregistration.k8s.io"
"apiextensions.k8s.io"
"scheduling.k8s.io"
"coordination.k8s.io"
"node.k8s.io"
"discovery.k8s.io"
"monitoring.coreos.com"
"metrics.k8s.io"
```

