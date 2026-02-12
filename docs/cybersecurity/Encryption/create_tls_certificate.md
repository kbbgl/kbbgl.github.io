---
slug: how-to-create-self-signed-tls-certificate
title: How to Create a Self-Signed TLS Certificate
description: How to Create a Self-Signed TLS Certificate
authors: [kgal-akl]
tags: [cybersecurity,certificates,encryption,openssl,cli]
---

In some cases we'd want to set up TLS on a webserver for testing purposes.

In Kubernetes we can utilize the [CertManager](../../software/containerization/k8s/cert-manager.md) to create it for us.

This tutorial explains how to create one manually using `openssl`.

Let's say we have an EKS cluster named `kgal-eks.dev` and we'd like to create a new Kubernetes `Ingress` (e.g. `IngressClass` is installed `alb`) that would terminate TLS for the host `app.kgal-eks.dev`. The `Ingress` will use a Kubernetes `Secret` to populate the TLS server certificate and private key.

## Generate Server Key and CSR

Create a server configuration file:

```
[req]
distinguished_name = req_distinguished_name
req_extensions = v3_req
prompt = no

[req_distinguished_name]
CN = app.kgal-eks.dev

[v3_req]
keyUsage = keyEncipherment, dataEncipherment, digitalSignature
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1 = app.kgal-eks.dev
```

Then create the server key and CSR:

```bash
openssl genrsa -out app.kgal-eks.dev.key.pem 2048
openssl req -new -key app.kgal-eks.dev.key.pem -out app.kgal-eks.dev.csr -config app-kgal-eks-dev.cnf
```

and sign it with a Certificate Authority (CA):

```bash
openssl x509 -req -in app.kgal-eks.dev.csr \
  -CA ca.pem -CAkey ca.key.pem -CAcreateserial \
  -out app.kgal-eks.dev.crt.pem -days 825 \
  -extensions v3_req -extfile app-kgal-eks-dev.cnf
```

## Create Kubernetes TLS Secret

The `Ingress` expects the `tls.crt` to be the server certificate first and then the CA so clients can verify the chain:

```bash
cat app.kgal-eks.dev.crt.pem ca.pem > app.kgal-eks.dev.chain.pem
```

```bash
kubectl create secret tls tls-kgal-eks-dev-crt \
--cert=app.kgal-eks.dev.chain.pem \
--key=app.kgal-eks.dev.key.pem
```