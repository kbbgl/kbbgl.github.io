---
slug: passive-net-cspture-techniques
title: Passive Network Capture Techniques
tags: [cybersecurity, network]
---

## System Call Tracing

kernel exposes syscalls to application running in user-mode. 

We can monitor these calls directly to passively extract data from an application.

unix syscalls:

`socket` - creates new socket fd
`connect` - connects socket to IP address and port
`bind` - bind socket to local IP address and port.
`recv, read, recfrom` - receives data from network via the socket
`send, write, sendfrom` - sends data over network via the socket


### strace/dtrace

```bash
strace -e trace=network,read,write /path/to/app args
```

### Process Monitor on Windows