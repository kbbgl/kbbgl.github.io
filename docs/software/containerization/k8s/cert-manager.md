---
title: How to Deploy Cert Manager in Kubernetes Cluser
slug: how-to-deploy-cert-manager-k8s
app2or: kgal-akl
tags: [devops, certificate, clm, kubernetes, clm, k8s]
---

## Install Cert Manager

```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml
```

## Create Cluster Issues

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: selfsigned-issuer
spec:
  selfSigned: {}
```

## Create Certificate

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: kgal-localhost-kind-dev-tls
  namespace: akeyless-gateway
spec:
  secretName: kgal-localhost-kind-dev-tls-cert-manager
  issuerRef:
    name: selfsigned-issuer
    kind: ClusterIssuer
  commonName: kgal-localhost.kind.dev
  dnsNames:
    - kgal-localhost.kind.dev
    - localhost
```

```bash
kubectl apply -f issuer.yaml
kubectl apply -f cert.yaml
```

Wait for the TLS secret to get created:

```bash
kubectl get secrets kgal-localhost-kind-dev-tls-cert-manager -o yaml | yq '.data | keys'
- ca.crt
- tls.crt
- tls.key
```
