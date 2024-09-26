---
slug: docker-shim-to-interface
title: Get Docker Shim to Network Interface
authors: [kbbgl]
tags: [docker, network]
---

Printing all socket information:

```bash
ss | grep 4cfb89dd2d4b11f0ea3cad2d11f66543513149a2f45ca876dc497a7ac965b021

u_str             ESTAB               0                    0                    @/containerd-shim/moby/4cfb89dd2d4b11f0ea3cad2d11f66543513149a2f45ca876dc497a7ac965b021/shim.sock@ 561241
```

Which the container logs can be found in:

```bash
sudo ls /var/lib/docker/containers | grep 4cfb89dd2d4b11f0ea3cad2d11f66543513149a2f45ca876dc497a7ac965b021

4cfb89dd2d4b11f0ea3cad2d11f66543513149a2f45ca876dc497a7ac965b021
```
