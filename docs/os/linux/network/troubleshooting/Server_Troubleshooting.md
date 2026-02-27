---
slug: os-linux-network-troubleshooting-server-troubleshooting
title: "Server Troubleshooting"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, server_troubleshooting]
---

# Server Troubleshooting

```bash
netstat -taupe | grep httpd

sudo ss -ltp | grep httpd


# Access control. 

man host_acccess

cat /etc/hosts.allow
cat /etc/hosts.deny

```

For advanced server troubleshooting, the `/proc` filesystem has settings that affect the network stack:

- `/proc/sys/net/ipv4/ip_forward`
Allows for network traffic to be forwarded from one interface to another.

- `/proc/sys/net/ipv4/conf/*/accept_redirects`
Accepting Internet Control Message Protocol (ICMP) redirects from a router to find better routes. This setting has the potential to be exploited by a malicious party to redirect your traffic.

- `/proc/sys/net/ipv4/icmp_echo_ignore_all`
Changing this setting will affect the host's visibility to ICMP ping packets.

- `/proc/sys/net/ipv4/icmp_echo_ignore_broadcasts`
This setting will change the host's visibility to broadcast ICMP ping packets.

- `/proc/net/arp`
Contains the current arp table.

These settings are not persistent across reboots. To make the changes persistent, edit the `/etc/sysctl.conf` configuration file, or a `.conf` file in the `/etc/sysctl.d` directory.

The syntax for `/etc/sysctl.conf` matches the path for the file in `/proc/sys` with the . character instead of /.

## Common Server-Side Problems

Common server problems include broken DNS, overzealous firewall rules, incorrect network settings, and the daemon not listening on the right interface/port.

Some access control systems require that Reverse DNS be properly set up.

When enabling new traffic to pass through a firewall, pay attention to the type of protocol (for example, UDP over TCP) used.

Some protocols break when return traffic comes back from a different IP address. Verify that your egress route is correct.
