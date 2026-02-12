---
slug: tls-nodejs
title: TLS in NodeJS
description: How to solve common TLS problems in NodeJS environment
authors: [kgal-akl]
tags: [cybersecurity,certificates,encryption,openssl,cli,nodejs,js,javascript]
---

## UNABLE_TO_VERIFY_LEAF_SIGNATURE

When attempting to communicate with an HTTP server that has TLS enabled with a self-signed certificate, we will encounter an error `UNABLE_TO_VERIFY_LEAF_SIGNATURE`. 

The error indicates a problem with an SSL/TLS certificate chain validation, usually in a Node.js environment or a system using a similar certificate store. The issue often arises when a server fails to provide a complete certificate chain, or when a client does not inherently trust a self-signed or internal intermediate certificate, such as those used in corporate proxies. In the case of GitHub Actions, the error `UNABLE_TO_VERIFY_LEAF_SIGNATURE` can happen if the process that runs the action doesn’t use the same CA bundle as the one extended with `NODE_EXTRA_CA_CERTS`, or if the action’s dependencies use another TLS stack. Installing the certificate in the runner’s system trust store makes every TLS client on that machine (including the NodeJS-based client) trust the remote server.

To install the CA certificate on the host machine on [RHEL](../../../../cybersecurity/Encryption/certificate.md#rhel), [Ubuntu](../../../../cybersecurity/Encryption/certificate.md#ubuntu) or [MacOS Keychain](../../../../cybersecurity/Encryption/certificate.md#macos-keychain).

