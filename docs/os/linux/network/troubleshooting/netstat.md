---
slug: os-linux-network-troubleshooting-netstat
title: "Netstat Cheat Sheet"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, netstat]
---

# Netstat Cheat Sheet

Best command, will provide `pid` and name, listening, network stats:

```bash
netstat -tulpn
```

list all listening ports, UDP/TCP:

```bash
netstat -a
```

list only tcp:

```bash
netstat -at
```

list only udp:

```bash
netstat -au
```

list all actively listening ports:

```bash
netstat -l
```

`LISTENING` means that a service is listening for connections on that port.

Once a connection is established it will be `ESTABLISHED`, and you'll have a matching foreign address on the line.

add pid output:

```bash
netstat -p
```

show statistics by protocol:

```bash
netstat -s
```

show statistics for tcp:

```bash
netstat -st
```

raw network stats for tcp:

```bash
netstat -t --statistics --raw
```

show I/O by interface:

```bash
netstat -i

Kernel Interface table
Iface      MTU    RX-OK RX-ERR RX-DRP RX-OVR    TX-OK TX-ERR TX-DRP TX-OVR Flg
cali008a  1500   461097      0      0 0        712668      0      0      0 BMRU
cali00d7  1500    22246      0      0 0         35126      0      0      0 BMRU
cali01ee  1500  1424714      0      0 0       2091930      0      0      0 BMRU
cali0261  1500    51412      0      0 0         90721      0      0      0 BMRU
cali0354  1500    48491      0      0 0         98009      0      0      0 BMRU
cali0b41  1500  2416976      0      0 0       2451894      0      0      0 BMRU
cali0c0e  1500 80759328      0      0 0      74711267      0      0      0 BMRU
cali0c65  1500    40621      0      0 0         91110      0      0      0 BMRU
cali14e9  1500    21601      0      0 0         47504      0      0      0 BMRU
cali1ac2  1500    65225      0      0 0        136848      0      0      0 BMRU
```
