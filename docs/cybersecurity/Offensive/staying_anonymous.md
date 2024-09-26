---
slug: staying-anon
title: Staying Anonymous
description: Staying Anonymous
authors: [kbbgl]
tags: [cybersecurity,offensive,anonymous,shell,netcat]
---

there are 4 methods:

* Onion network
* Proxy servers
* VPN
* Private encrypted email

Every request sent across internet includes the source IP. The request hops between different routers before reaching the destination. To see all hops, we can use `traceroute`.

## Tor

uses the Onion Router System, a network of 7k routers that encrypt the data, destination and source address of each packet. Each hop is a cycle of encryption/decryption. So if someone intercepts one of the messages at one of the hops, they can only see the previous hop information rather than the original source address.

![alt](https://i.insider.com/558968d9eab8eaf3664c3a5f?width=1300&format=jpeg&auto=webp)

## Proxy servers

Traffic will look like it came from the source address but will in fact look like a different address.

`proxychains` is a tool that can be used to run a command through a chain of proxies:

```bash
proxychains nmap -T4 -A 10.10.10.1
```

example  `proxychains` configuration (`/etc/proxychains.conf`) has 3 proxies set up:

```bash
# ProxyList section, defaults to tor
[ProxyList] 
# add proxy here... 
socks4 114.134.186.12 22020 
socks4 188.187.190.59 8888 
socks4 181.113.121.158 335551
```

uncomment `dynamic_chain` and `strict_chain` to have `proxychains` proxy requests through each proxy and skip if one fails.

uncomment `random_chain` (and comment `strict_chain`) so each time we use `proxychains` , the proxy will look different to the target, making it harder to track our traffic from its source. We can also set the number of proxies by changing the value next to `chain_len`
