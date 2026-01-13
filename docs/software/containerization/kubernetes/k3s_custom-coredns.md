---
title: How to Add Custom Host to K3S
slug: how-to-add-custom-host-k3s
app2or: kgal-akl
tags: [devops, kubernetes, dns, k8s]
---

Create a new file `coredns-custom.yaml`:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns-custom
  namespace: kube-system
data:
  default.server: |
    kbbgl.github.dev {
      hosts {
        123.123.123.123 kbbgl.github.dev
        fallthrough
      }
    }
```

```bash
kubectl apply -f coredns-custom.yaml
kubectl -n kube-system rollout restart deployment coredns
```

Source:
- https://metalcoder.dev/add-custom-dns-entries-to-k3s/