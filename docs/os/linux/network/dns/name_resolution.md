---
slug: os-linux-network-dns-name-resolution
title: "Name Resolution"
authors: [kbbgl]
tags: [os, linux, network, dns, name_resolution]
---

# Name Resolution

Name resolution is **translating hostnames to IP address**.

There are two ways for translation:

- **Static** using `/etc/hosts`. It is checked before reaching the DNS server (can be modified in `/etc/nsswitch.conf`)
- **Dynamic** using DNS servers.

The tools we can use to check name resolution:

```bash
dig $hostname
host $hostname
# older
nslookup $hostname
```

Other host-related files are:

```bash
/etc/hosts.deny
/etc/hosts.allow

# rarely used
/etc/host.conf
```

## DNS

The server DNS is configured in `/etc/resolv.conf`. It specifies:

- Specify particular domains to search
- Define a strict order of nameservers to query
- May be manually configured/updated from a service (such as DHCP)
