---
slug: os-linux-network-dns-change-dns-server
title: "Change DNS Server"
authors: [kbbgl]
tags: [os, linux, network, dns, change_dns_server]
---

# Change DNS Server

Entries are kept in `/etc/resolv.enf`.

To change the DNS server, we would modify the entry:

```bash
➜  ~ cat /etc/resolv.conf
nameserver $SOME_DNS_SERVER
```

When using a DHCP address and the DHCP server provides a DNS setting, the DHCP server will replace the contents of the file when it renews the DHCP address.
