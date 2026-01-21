---
slug: certificate-troubleshooting
title: Troubleshooting Certificates
description: Troubleshooting Certificates
authors: [kbbgl]
tags: [cybersecurity,certificates,encryption,openssl,cli]
---

## Display Certificate Information

### Full Information

```bash
openssl x509 -in /tmp/kgal-localhost-kind-kubeconfig-cluster-cert-auth-data.pem -noout -text
```

The output will look something like this:
```
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 851...626 (0x763....1c3a)
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=kubernetes
        Validity
            Not Before: Sep 23 15:36:35 1970 GMT
            Not After : Sep 21 15:41:35 2070 GMT
        Subject: CN=kubernetes
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:b7:3f...:75
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment, Certificate Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier: 
                F3:A2:...:3C
            X509v3 Subject Alternative Name: 
                DNS:kubernetes
    Signature Algorithm: sha256WithRSAEncryption
    Signature Value:
        65:91:a9:....b0:9c:48:31
```


### Specific Subject and Dates

```bash
openssl x509 \
-in /tmp/kgal-localhost-kind-kubeconfig-cluster-cert-auth-data.pem \
-noout \
-subject \
-dates
```

```
subject=CN=kubernetes
Not Before: Sep 23 15:36:35 1970 GMT
Not After : Sep 21 15:41:35 2070 GMT
```


## Verifying Certificate

This command checks if a certificate is trusted by a given Certificate Authority (CA):

```bash
openssl verify -CAfile ca.pem crt.pem
```


## Public Keys

### Extract Public Key from Certificate

This extracts the public key from a certificate and displays it in PEM format:

```bash
openssl x509 \
-in /tmp/kgal-localhost-kind-kubeconfig-cluster-cert-auth-data.pem \
-noout \
-pubkey

-----BEGIN PUBLIC KEY-----
MIIBIj....AB
-----END PUBLIC KEY-----
```


### Displaying Public Key Information

```bash
openssl x509 \
-in /tmp/kgal-localhost-kind-kubeconfig-cluster-public-key-data.pem \
-noout \
-pubin \
-text
```


## Private Keys

### Displaying Private Key Information

```bash
openssl pkey -in privatekey.pem -noout -text
```

### Checking a Private Key

```bash
openssl rsa -in privatekey.pem -check -noout
```


## Certificate Signing Request (CSR)

Print the CSR content:

```bash
openssl req -in docs.kbbgl-gh-io.dev.key -text -noout -verify
```


## Add Self-Signed Certificate to CA Store in Ubuntu

```bash
CRT_PATH=/tmp/kbbgl-gh-io-tls.crt

sudo cp $CRT_PATH /usr/local/share/ca-certificates/

sudo update-ca-certificates

Updating certificates in /etc/ssl/certs...
1 added, 0 removed; done.
Running hooks in /etc/ca-certificates/update.d...
done.
```


## Checking Kubernetes TLS Secret

```bash
kubectl get secrets -n default tls-kbbgl-gh-io-dev-crt -o jsonpath='{.data.tls\.crt}' | base64 -d | openssl x509 -text -noout | grep -A 1 "Subject Alternative Name"
            X509v3 Subject Alternative Name: 
                DNS:kbbgl-gh-io.dev, DNS:*.kbbgl-gh-io.dev
```