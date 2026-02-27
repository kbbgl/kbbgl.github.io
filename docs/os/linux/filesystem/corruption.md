---
slug: os-linux-filesystem-corruption
title: "Corruption and Recovery"
authors: [kbbgl]
tags: [os, linux, filesystem, corruption]
---

# Corruption and Recovery

If during the boot process, one or more filesystems fail to mount, `fsck` may be used to attempt repair. However, before doing that one should check that `/etc/fstab` has not been misconfigured or corrupted. Note once again that you could have a problem with a filesystem type the kernel you are running does not understand.

If the root filesystem has been mounted you can examine this file, but / may have been mounted as read-only, so to edit the file and fix it, you can run:

```bash
sudo mount -o remount,rw /
```

to remount it with write permission.

If `/etc/fstab`, seems to be correct, you can move to `fsck`. First, you should try:

```bash
sudo mount -a
```

to try and mount all filesystems. If this does not succeed completely, you can try to manually mount the ones with problems. You should first run fsck to just examine; afterwards, you can run it again to have it try and fix any errors found.
