---
slug: os-linux-network-network-filesystems-troubleshooting
title: "Troubleshooting"
authors: [kbbgl]
tags: [os, linux, network, network_filesystems, troubleshooting, filesystem]
---

# Troubleshooting

The most common errors found in setting up network filesystems are:

- Incorrect firewall settings
Older versions of NFS are more difficult to check; you need to open multiple ports to do so.
- Incorrect access control settings
Test from a different guest in the same network, or in a different network.
- Syntax errors in configuration files
Use testparm, showmount and the like for debugging​.
- NFS: `showmount -e <SERVER>`
- SMB: `smbclient`
