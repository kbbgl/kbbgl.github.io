---
slug: net-inspection-k8s-container
title: Net Inspection in Kubernetes Container
authors: [kgal-akl]
tags: [container, network, debug, tcpdump, k8s]
---

Let me break down the options for your specific use case of capturing **localhost HTTP traffic** inside a Kubernetes `Pod` named `pod-78db6d7fc-gtdk9` with a container named `container`.

### `kubectl debug` (Recommended)

This is the least invasive and fastest approach. An ephemeral container shares the `Pod`'s network namespace (including the `lo` loopback interface), so it can see all localhost traffic without any `Pod` restart or chart changes.

```bash
kubectl debug -it pod-78db6d7fc-gtdk9 \
--image=nicolaka/netshoot \
--target=container \
-- bash
```

Once inside, you can capture localhost HTTP traffic on the loopback interface:

```bash
# Raw tcpdump on loopback - ports the pod listens on (8000, 8080, 8081, 8200)
tcpdump -i lo -A -s 0 'tcp port 8443 or tcp port 8081 or tcp port 8000 or tcp port 8200'

# Or for more human-readable HTTP output:
ngrep -d lo -W byline 'HTTP' 'port 8080 or port 8081'

# Or use tshark for structured output:
tshark -i lo -f 'tcp port 8080 or tcp port 8081' -Y http
```

**Why this works well:** When your pod already runs as root (`runAsUser: 0`), and `netshoot` has `tcpdump`, `ngrep`, `tshark`, and other network tools pre-installed. The ephemeral container shares the network namespace via `--target`, giving full visibility into `lo`.

---

### Sidecar container

This is the right approach if you want a _persistent_ traffic inspection setup (e.g., always-on logging), but overkill for ad-hoc debugging. The sidecar shares the same network namespace so it can also see localhost, but the restart makes it less ideal for a quick investigation.

---

### `kubectl gadget` / Inspektor Gadget

This uses eBPF at the node level -- very powerful and zero-modification to the workload. However:
- It requires Inspektor Gadget to be **deployed as a `DaemonSet`** on your cluster first.
- It's better suited for cluster-wide observability rather than one-off HTTP inspection of a single pod.
- The `tcpdump` gadget works, but the setup overhead isn't worth it for a quick localhost capture.

---

### Practical recommendation

Start with the ephemeral debug container. If you want to save a pcap for later analysis in Wireshark:

```bash
# Inside the ephemeral container:
tcpdump -i lo -s 0 -w /tmp/pod-localhost.pcap 'tcp port 8080 or tcp port 8081'

# Then from your machine, copy it out:
kubectl cp <namespace>/pod-78db6d7fc-gtdk9:/tmp/pod-localhost.pcap ./pod-localhost.pcap -c debugger-xxxxx
```

If the ephemeral container approach hits permission issues (some cluster policies restrict ephemeral containers or `SYS_PTRACE`/`NET_RAW` capabilities), then the sidecar approach becomes the fallback.

