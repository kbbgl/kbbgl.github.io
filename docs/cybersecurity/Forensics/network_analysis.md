---
slug: network-capture-techniques
title: Network Capture Techniques
tags: [cybersecurity, network, hacking]
---

## Passive Network Traffic Capture

### Wireshark

### System Call Tracing

kernel exposes syscalls to application running in user-mode. 

We can monitor these calls directly to passively extract data from an application.

unix syscalls:

`socket` - creates new socket fd
`connect` - connects socket to IP address and port
`bind` - bind socket to local IP address and port.
`recv, read, recfrom` - receives data from network via the socket
`send, write, sendfrom` - sends data over network via the socket


#### strace/dtrace

```bash
strace -e trace=network,read,write /path/to/app args
```

#### Process Monitor on Windows

Windows implements its user-mode network functions without direct system calls. The network stack is exposed through a driver and establishing a connection uses the file `open`, `read` and `write` system calls to configure a network socket to use.

We can use the Process Monitor tool to analyze network system calls.

## Active Network Traffic Capture

Usually is done using a Man-in-the-Middle technique where the capturing device is a bridge between the client and server.

It's the most valuable technique to analyze and exploit application network protocols.

### Network Proxies

#### Port-Forwarding Proxies

Port forwarding is the easiest way to proxy a connection. We set up a TCP port-forwarding proxy server that listens for new connections. When a new connection is made to the proxy server, it will open a forwarding connection to the real service and logically connect the two.

We then send the request that was meant to the intended server to the port-forwarding proxy server first by sending the client request to `http://localhost:$LOCALPORT/resource` if the original request was headed to `http://domain.com/resource`.

#### SOCKS Proxies

SOCKS proxy are port-forwarding proxies on steroids. It has version 4, 4a and 5.

In Java, we can specify the `socksProxyHost` and `socksProxyPort` arguments to redirect clients to the SOCKS proxy:

```bash
java -DsocksProxyHost=localhost -DsocksProxyPort=1080 SocketClientApplication
```

#### HTTP Proxies

The two main types of HTTP proxies are forwarding and reverse proxy.

## Network Protocol Structures

