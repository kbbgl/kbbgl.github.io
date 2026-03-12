---
slug: too-many-open-files-kubernetes-inotify-debugging
title: "Debugging 'Too many open files' in Kubernetes: nofile vs inotify/fsnotify"
authors: [kbbgl]
tags: [kubernetes, linux, debugging, golang]
date: 2026-03-12
---

When you see **`too many open files`** in a containerized app, it’s tempting to jump straight to `ulimit -n`. Sometimes that’s correct. But on Linux (especially with Go apps using `fsnotify`), the error can also be caused by **inotify limits**—even if your process has a huge file-descriptor limit.

This post is a practical, copy/paste-friendly checklist to debug the problem on a real Kubernetes cluster.

## Step 0: Decide which “limit” you’re hitting

There are (at least) three common failure modes that can all look like “too many open files”:

- **Per-process file descriptor limit** (classic `EMFILE`)
  - Think: `ulimit -n`, `/proc/<pid>/limits`, or systemd `LimitNOFILE`.
- **System-wide file table exhaustion** (`ENFILE`)
  - Think: `/proc/sys/fs/file-nr` approaching `/proc/sys/fs/file-max`.
- **inotify instance/watch limits** (common with `fsnotify`)
  - Think: `fs.inotify.max_user_instances` and `fs.inotify.max_user_watches`.

The rest of the tutorial helps you quickly identify which one applies.

## Step 1: Check the real limits of the failing process (not your shell)

Inside the pod:

```bash
# Find your process (adjust the pattern for your app)
ps -eo pid,comm,args | grep -E 'myapp|server|gateway' | grep -v grep

# Replace <pid> with the real PID
cat /proc/<pid>/limits | sed -n '/Max open files/p'

# Count currently open FDs
ls /proc/<pid>/fd | wc -l
```

Why this matters:

- `ulimit -n` shows the limit for *your current shell*, which may be totally different from a process started by an init system (systemd, supervisord, Kubernetes runtime, etc.).

## Step 2: Check node-wide file table pressure (`ENFILE`)

On the Kubernetes node:

```bash
cat /proc/sys/fs/file-nr
cat /proc/sys/fs/file-max
```

`file-nr` is usually three numbers: allocated, unused, max. If allocated is near max, you have a node-level exhaustion problem that can break unrelated workloads.

## Step 3: Check inotify limits (the usual `fsnotify` culprit)

On the node (or inside the container—these are node kernel settings):

```bash
cat /proc/sys/fs/inotify/max_user_instances
cat /proc/sys/fs/inotify/max_user_watches
cat /proc/sys/fs/inotify/max_queued_events
```

If your app uses Go file watching, `fsnotify`’s Linux notes are worth skimming: [fsnotify README (Linux)](https://github.com/fsnotify/fsnotify?tab=readme-ov-file#linux).

Key concept:

- `max_user_instances` is **per-UID**. In Kubernetes, multiple containers/processes can share the same numeric UID (e.g., a non-root “app user”), which means they share the same inotify instance budget on that node.

## Step 4: Count inotify “instances” currently in use (who’s consuming them?)

On Linux, inotify instances show up as file descriptors named `anon_inode:inotify`.

Run this on the node to see which processes (and UIDs) are holding inotify instances:

```bash
for pid in /proc/[0-9]*; do
  p=${pid#/proc/}
  [ -r "$pid/fd" ] || continue
  c=0
  for fd in "$pid"/fd/*; do
    [ "$(readlink "$fd" 2>/dev/null)" = "anon_inode:inotify" ] && c=$((c+1))
  done
  [ "$c" -gt 0 ] || continue
  uid=$(awk '/^Uid:/{print $2}' "$pid/status" 2>/dev/null)
  cmd=$(tr '\0' ' ' < "$pid/cmdline" 2>/dev/null)
  echo "uid=$uid pid=$p inotify_instances=$c $cmd"
done | sort -k3 -nr | head -n 50
```

What to look for:

- A single UID with `inotify_instances` totals near `max_user_instances`
- Your app process holding many `anon_inode:inotify` FDs
- A node “agent” (log collector, metrics sidecar, etc.) consuming a lot

## Step 5: (Optional) Count how many watches each inotify FD holds

If you suspect “too many watches” (not instances), you can inspect `/proc/<pid>/fdinfo/*`:

```bash
pid=<pid>
for fd in /proc/$pid/fd/*; do
  [ "$(readlink "$fd" 2>/dev/null)" = "anon_inode:inotify" ] || continue
  n=$(grep -c 'inotify wd' /proc/$pid/fdinfo/${fd##*/} 2>/dev/null || true)
  echo "pid=$pid fd=${fd##*/} watches=$n"
done | sort -t= -k3 -nr | head -n 20
```

## Step 6: “How do I kill the file watcher?”

You generally can’t “kill a watch” directly. The kernel releases it when the owning process closes the FD.

Practical actions:

- **Restart the pod / process** that owns the inotify instances.
- If a node is saturated, **cordon/drain** and reschedule to another node as a temporary workaround.
- Fix the root cause by raising limits and/or reducing watcher usage.

## Step 7: Mitigation—raise inotify instance limits (and make it persistent)

If you’ve confirmed `max_user_instances` is the bottleneck, increasing it is often the quickest fix.

Temporary (until reboot):

```bash
sysctl -w fs.inotify.max_user_instances=1024
```

Persistent:

- Add a sysctl config file on the node (exact location varies by distro), for example:
  - `/etc/sysctl.d/99-inotify.conf`
- Then reload sysctls (varies by environment), commonly:
  - `sysctl --system`

In managed Kubernetes, you may prefer:

- baking `sysctl` settings into your node image/bootstrap
- setting them via a privileged `DaemonSet` (policy-dependent)

## Step 8: Don’t forget `nofile` (it’s still real)

Even if inotify was the cause this time, it’s worth capturing `nofile` facts for your runbook:

- **Systemd defaults**:

```bash
systemctl show --property DefaultLimitNOFILE
systemctl show --property DefaultLimitNOFILESoft
```

- **Unit-specific limits** (examples):

```bash
systemctl show kubelet --property LimitNOFILE
systemctl show containerd --property LimitNOFILE
```

And always verify the actual process limit via `/proc/<pid>/limits` (Step 1).

## Summary: the fastest “root cause” loop

1. Check the failing process’s **actual** `Max open files` and FD count.
2. Check node `file-nr`/`file-max` for **system-wide FD exhaustion**.
3. Check `inotify` `sysctls` and enumerate `anon_inode:inotify` to find the **UID/process** consuming instances.
4. Apply the smallest safe mitigation (restart offender, reschedule, raise `sysctl`) and confirm the error disappears.

