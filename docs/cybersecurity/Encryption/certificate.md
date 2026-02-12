---
slug: certificate-troubleshooting
title: Troubleshooting Certificates
description: Troubleshooting Certificates
authors: [kbbgl]
tags: [cybersecurity,certificates,encryption,openssl,cli,k8s]
---

The documentation below references a PEM-encoded certificate.

```bash
CRT_PATH=/tmp/kbbgl-gh-io-tls.crt.pem
```

## Verifying Certificates

### Trusted

This command checks if a certificate is trusted by a given Certificate Authority (CA):

```bash
openssl verify -CAfile ca.pem crt.pem
```

We can also get the certificate directly and verify it from `stdin`:

```bash
openssl s_client -connect $REMOTE_SERVER_HOSTNAME:${REMOTE_SERVER_PORT:-443} -showcerts </dev/null 2>/dev/null | \
openssl verify -CAfile /etc/pki/ca-trust/source/anchors/$REMOTE_SERVER_HOSTNAME -
```

### Display Certificate Information

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


### Public Keys

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


### Private Keys

### Displaying Private Key Information

```bash
openssl pkey -in privatekey.pem -noout -text
```

### Checking a Private Key

```bash
openssl rsa -in privatekey.pem -check -noout
```


### Certificate Signing Request (CSR)

Print the CSR content:

```bash
openssl req -in docs.kbbgl-gh-io.dev.key -text -noout -verify
```

### Checking Kubernetes TLS Secret

```bash
kubectl get secrets -n default tls-kbbgl-gh-io-dev-crt -o jsonpath='{.data.tls\.crt}' | base64 -d | openssl x509 -text -noout | grep -A 1 "Subject Alternative Name"
            X509v3 Subject Alternative Name: 
                DNS:kbbgl-gh-io.dev, DNS:*.kbbgl-gh-io.dev
```

### Certificate Expiration

When validating certificates, there could be a couple of other components in the certificate that might be relevant to understand whether a certificate has expired or not.

The most basic expiration methods use a static list of certificate IDs:

```bash
# To see it in hex (standard view)
openssl x509 -in client.crt -noout -serial

# To see it in decimal
openssl x509 -in client.crt -noout -serial -ext serial

# Or more reliably via text output:
openssl x509 -in client.crt -noout -text | grep "Serial Number" -A 1
```

Or the "Not Before" "Not After" attribute:

```bash
openssl x509 -in client.crt -noout -text | grep -E "Not Before|After"
```


## Adding Self-Signed Certificate to CA Store 

### Ubuntu

```bash
sudo cp $CRT_PATH /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

### RHEL

```bash
sudo cp $CRT_PATH /etc/pki/ca-trust/source/anchors/kbbgl-gh-io-tls.crt.pem
sudo update-ca-certificates
sudo update-ca-trust extract
```


### MacOS Keychain

```bash
sudo security add-trusted-cert -d -r trustRoot -k "/Library/Keychains/System.keychain" $CRT_PATH
```

## Certificate Revocation

### CRL (Certificate Revocation List)

A Certificate Revocation List is a "blacklist" file published by a Certificate Authority (CA) containing a list of serial numbers of certificates that have been revoked before their scheduled expiration date.

The client downloads this file periodically, parses it, and searches for the serial number of the certificate being presented.

A CRL is similar to a printed "Wanted" poster at a post office. You have to check the whole list to see if the person in front of you is on it.

To check if a certificate has a CRL:

```bash
openssl x509 -in client.crt -noout -text | grep -A 4 "CRL Distribution Points"
```

We can also manually verify a certificate against a CA with CRL:

```bash
# Convert CRL to PEM if needed
openssl crl -inform DER -in list.crl -out list.pem

# Verify
openssl verify -crl_check -CAfile ca.crt -CRLfile list.pem client.crt
```

### OCSP (Online Certificate Status Protocol)

The Online Certificate Status Protocol is a real-time protocol used to query the status of a specific certificate.

Instead of downloading a large list, the client sends a small request to an "OCSP Responder" (a server) asking: "Is certificate #12345 still valid?" The server responds with "Good," "Revoked," or "Unknown."

This is similar to a credit card swipe. The merchant doesn't check a book of stolen cards; they ping the bank instantly to see if this specific card is valid right now.

To check and retrieve the OCSP URL:

```bash
openssl x509 -in client.crt -noout -text | grep "OCSP - URI"
```
