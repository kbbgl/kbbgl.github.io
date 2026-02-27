---
slug: os-linux-account-management-ssh
title: "SSH"
authors: [kbbgl]
tags: [os, linux, account_management, ssh]
---

# SSH

Enables remote login to servers. The important files are in:

```bash
$ ls ~/.ssh

authorized_keys
config
id_rsa
id_rsa.pub
known_hosts
```

- `config`: A configuration file for specifying various options. See `man ssh_config`
- `id_rsa`: private encryption key
- `id_rsa.pub`: public encryption key
- `authorized_keys`: A list of public keys that are permitted to login.
- `known_hosts`: A list of hosts from which logins have been allowed in the past
