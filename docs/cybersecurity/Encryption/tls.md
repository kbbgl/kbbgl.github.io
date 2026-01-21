---
slug: tls-troubleshooting
title: TLS/SSL Troubleshooting
authors: [kgal-akl]
tags: [tls, network, handshake, encryption, ciphersuite]
---

## Test Connectivity with CA/certs

```bash
openssl s_client \
-connect kbbgl.github.io.dev:8443 \
-cert kbbgl.gh.io.crt.pem \
-key kbbgl.gh.io.key.pem \
-CAfile kbbgl.gh.io.ca.pem \
-verify-hostname kbbgl.github.io.dev \
-tls1_2 \
-showcerts
```

## Cipher Suites Mismatch

When there's mismatch between what cipher suites the client offers and what the server supports we'll see an error such as `tls: handshake failure`.

We can find the cipher suite that the server chooses if we compare two PCAPs:

```bash
apt install -y tshark net-tools

tshark \
-i eth0 \
-f "host $HOSTNAME and port $PORT" \
-w /tmp/capture.pcap \
-V \
-c 500
```

One for the success and one for the failure and look at the TLS > Handshake Protocol (Server/Client Hello) > Cipher Suites section to see what ciphersuites are sent by the client and accepted by the server.



