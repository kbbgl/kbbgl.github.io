---
slug: os-linux-network-protocols-ntp
title: "Network Time Protocol (NTP)"
authors: [kbbgl]
tags: [os, linux, network, protocols, ntp]
---

# Network Time Protocol (NTP)

Many protocols require consistent, if not accurate time to function properly.

The security of many encryption systems is highly dependent on proper time. Industries such as commodities or stock trading require highly accurate time, as a difference of only seconds can mean hundreds if not thousands of dollars lost or earned. NTP time sources are divided up into strata.

A `strata 0` clock is a special purpose time device (atomic clock, GPS radio, etc).
A `strata 1` server is any NTP server connected directly to a `strata 0` source (over serial or the like).
A `strata 2` server is any NTP server which references a `strata 1` server using NTP.
A `strata 3` server is any NTP server which references a `strata 2` server using NTP.
NTP may function as a client, a server, or a peer:

Client: Acquires time from a server or a peer.
Server: Provides time to a client.
Peers: Synchronize time between other peers, regardless of the defined servers.
Below you can see the Network Time Protocol illustrated.
