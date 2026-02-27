---
slug: os-linux-network-firewalls-tcp-wrappers
title: "TCP Wrappers"
authors: [kbbgl]
tags: [os, linux, network, firewalls, tcp_wrappers]
---

# TCP Wrappers

The TCP Wrappers system is a host-based network firewall and ACL. Originally, it only protected the `inetd` system, but has now been extended with the shared object library `libwrap`.

The configuration for `tcpwrappers` is handled by two files, `/etc/hosts.allow` and `/etc/hosts.deny`. Both files have the same syntax:

```bash
# <DAEMON>:<CLIENT>
```

The `<DAEMON>` should match the name of the binary of the service (e.g. `sshd`). The `<CLIENT>` pattern can be:

- An IP address: `10.30.21.7`.
- A network/netmask: `10.30.21.0/255.255.255.0`.
- A domain name: `.foo.example.com`.
- A partial address: `10.30.21`.
- A file name full of the above patterns: `/etc/ssh-hosts.allow`.

When traffic comes to a `libwrap-enabled` daemon, those two files are consulted to see the following:

- If the pattern matches in `/etc/hosts.allow`, the traffic is permitted.
- If the pattern is not found in `/etc/hosts.allow`, and it matches in `/etc/hosts.deny`, the traffic will be denied.
- If the pattern does not match in either file, the traffic will be permitted.

Below you will find some examples of TCP wrappers:

```bash
hosts.allow
vsftpd:ALL
ALL:LOCAL
ALL:10
ALL:.example.com EXCEPT untrusted.example.com

hosts.deny
ALL:ALL
```
