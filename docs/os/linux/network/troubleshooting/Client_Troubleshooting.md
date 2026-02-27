---
slug: os-linux-network-troubleshooting-client-troubleshooting
title: "Client Troubleshooting"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, client_troubleshooting]
---

# Client Troubleshooting

## Simple Client Troubleshooting

The basics of network troubleshooting usually deal with basic connectivity testing. You can use the tools ping, traceroute, and nmap to test connectivity. Remember to test both DNS hostnames and IP addresses to diagnose DNS-related issues.

### `ping` for connectivity

- as a reminder, `ping` uses the ICMP protocol, normal traffic uses TCP
- use the IP address to your adapter
- use the IP address to your gateway address
- use the IP address of the DNS
- use the DNS name to see if name resolution is working.

### `traceroute` or `mtr`

These two will show the connectivity path to the destination. mtr will show statistics of the connection and packets drops or failures.

### `nmap`

It scans the server to see if the required ports are available to you.

## Intermediate Client Troubleshooting

Test plain-text protocols by using the `telnet` command:

```bash
telnet example.com 80


Trying 192.0.43.10...
Connected to example.com.
Escape character is '^]'.
GET /
<html>
<head><title>welcome to example.com</title></head>
<body>
  <h1>welcome to example.com</h1>
</body>
</html>
```

You can also do the same with SSL or TLS protocols.

```bash
openssl s_client -connect www.example.com:443
```

Use the `arp` command to check the link-layer connectivity.

## Advanced Client Troubleshooting

The `tcpdump` command and `wireshark` tool are useful when you need to dig deeper into a protocol. The command line-based `tcpdump` truncates packets by default and generates pcap files.

`wireshark` uses the graphical interface to capture packets. It can capture and analyze packets in real time. It is useful to analyze `.pcap` files, but you may not want `wireshark` installed on the system you are troubleshooting.

To capture packets with tcpdump for use with wireshark, use:

```bash
sudo tcpdump -i eth0 -s 65535 -w capture.pcap port 22
```

## Common Client-Side Problems

Some common networking issues found at the client side include:

DNS issues

- Can you ping the IP address but not the hostname?

Firewall issues

- A firewall on the client side which is rejecting the return traffic from a network request will cause problems.

Incorrect network settings

- Make sure the IP address is correct. Does it match the DNS host name?
- If the route is wrong or missing, traffic will not get to the other network node.
- Netmasks determine network routes, thus it is important to have the netmask of your host correct.
