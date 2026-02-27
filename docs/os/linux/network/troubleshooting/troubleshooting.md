---
slug: os-linux-network-troubleshooting-troubleshooting
title: "Troubleshooting"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting]
---

# Troubleshooting

## IP Configuration

Use `ifconfig` or `ip` to see if the interface is up, and if so, if it is configured.

## Network Driver

If the interface can't be brought up, maybe the correct device driver for the network card(s) is not loaded. Check with lsmod if the network driver is loaded as a kernel module, or by examining relevant pseudo-files in `/proc` and `/sys`, such as `/proc/interrupts` or `/sys/class/net`.

## Connectivity

Use `ping` to see if the network is visible, checking for response time and packet loss. `traceroute` can follow packets through the network, while `mtr` can do this in a continuous fashion. Use of these utilities can tell you if the problem is local or on the Internet.

## Default Gateway and Routing Configuration

Run `route -n` and see if the routing table makes sense.

## Hostname Resolution

Run `dig` or `host` on a URL and see if DNS is working properly.
