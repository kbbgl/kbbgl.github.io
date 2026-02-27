---
slug: os-linux-network-firewalls-firewalls
title: "Firewalls"
authors: [kbbgl]
tags: [os, linux, network, firewalls]
---

# Firewalls

A firewall is a network security system that monitors and controls network traffic. It applies to bot incoming and outgoing traffic. Can be implemented on HW and/or SW.

## Packet Filtering

Each transmitted packet has:

- **Header**
- **Payload**
- **Footer**

The header and footer include information about:

- the destination
- the source
- kind of packet
- protocol
- flags
- which packet number it is in the stream
- other metadata

The actual data is the Payload.

Packet filtering intercepts packets at one or more layers of the transmission (application, transport, network, datalink).

A firewall estabilishes a set of rules for each packet to be:

- Accept or reject based on content, address.
- Mangled
- Redirected to another address
- Inspected for security reasons

## Interface and Tools

Configuring a firewall can be done by:

- Using low-level tools and manually modifying contents of `/etc`.
- Using GUI such as `system-config-firewall`, `firewall-config`, `gufw`, `yast`.

### `firewalld` and `firewall-cmd`

`firewalld` is a dynamic firewall manager.

To enable it:

```bash
sudo apt install firewalld
sudo systemctl [enable\disable] firewalld
sudo systemctl [start\stop] firewalld
```

It's configuration can be found in:

```bash
# override other directories, sysadmin should work on this path
/etc/firewalld

/usr/lib/firewalld
```

To interact with `firewalld` we use:

```bash
firewall-cmd --help
```

**`iptables` service should be disabled when using `firewalld`**
