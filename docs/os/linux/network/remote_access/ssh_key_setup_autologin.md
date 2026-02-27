---
slug: os-linux-network-remote-access-ssh-key-setup-autologin
title: "SSH Autologin Setup"
authors: [kbbgl]
tags: [os, linux, network, remote_access, ssh_key_setup_autologin]
---

# SSH Autologin Setup

## Generate key

```bash
ssh-keygen
```

## Copy public key and permissions to remote server

```bash
ssh-copy-id -i /Users/me/.ssh/lnx12.pub me@lnx1
```
