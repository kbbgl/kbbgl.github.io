---
slug: how-ssl-works
title: How SSL Works
description: How SSL Works
authors: [kbbgl]
tags: [cybersecurity,certificates,encryption,ssl]
---


## Why

1) Privacy

2) Integrity

3) Identification

## How

### Encryption

There are 2 types of encryption algorithms:

1. Symmetric Key: There is only one key to encrypt and decrypt a message

The encryption key is mixed in with the message.
To decrypt, the recipient takes the same steps taken during encryption (but in reverse order)

1. Assymetric Key: There are 2 keys, one public and one private. They are paired and work together.

Flow:

U1 sends their public key to U2 =>
U2 sends message back to U1 encrypting it with the public key =>
U1 received the message and decrypts it with the private key.

Only the private key can open a box locked with a public key pair.

### Handshake

The client communicates with server to establish a secure connection to transmit messages. The negotiation in order to establish this connection is referred to as the 'handshake'.

The steps:

1. Client Hello - the client sends a list of SSL/TLS versions and encryption algorthims (a.k.a cipher suite) that the client can work with the server.
1. Server Hello - the server chooses which encryption algorithm and SSL/TLS version to use and replies with the server certificate which includes the **public key**.

1. Client Key Exchange - the client checks whether the certificate is legit and generates a pre-master key. The pre-master key is encrypted with the server public key and sent to the server.

1. Change Cipher spec - The server uses their private key to decrypt the pre-master key. All communication up to this point is not secure and no messages have been shared. They both generate the same 'shared secret' that they are going to use as a symmetric key. Then both send a test message to certify encryption.

1. Secure connection.
