---
slug: add-flaky-network
title: How to Add Network Flakiness
description: How to add network flakiness for testing applications.
authors: [kgal-akl]
tags: [network, testing, development, vm, linux, kernel, traffic, latency, performance, kubernetes, k8s, chaos]
---

On a VM, we can use the `tc` tool to manipulate outbound traffic control.

On an Azure VM that has a `k3s` cluster deployed, the traffic path from an application running in a Pod to SaaS service would go like this:

1. Gateway Pod -> `cni0` (bridge)
2. `cni0` -> host routing
3. Host -> physical interface (`eth0`) -> internet

So we will apply `tc` rules to the `eth0` interface. We can display information about the interface:

```bash
kgal@kgal-azure-dev:~$ ip a show eth0

6: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc prio state UP group default qlen 1000
    link/ether da:01:a4:37:22:4e brd ff:ff:ff:ff:ff:ff
    inet 10.42.0.1/24 brd 10.42.0.255 scope global cni0
       valid_lft forever preferred_lft forever
    inet6 fe80::5884:dfff:fe48:9a1/64 scope link
       valid_lft forever preferred_lft forever
```

The first step is to create a root queueing discipline (A.K.A. `qdisc`). Whenever the kernel needs to send a packet to an interface, it is enqueued to the `qdisc` configured for that interface. Immediately afterwards, the kernel tries to get as many packets as possible from the `qdisc`, for giving them to the network adaptor driver.

To create a root `qdisc` on `en0` interface with 500ms delay, 50ms jitter and 10% packet loss:

```bash
SAAS_IPS=("123.123.123.123" "122.122.122.122")
PODS_IPS=($(kubectl get pod -n  -l "app.kubernetes.io/name=myapp" -o jsonpath='{.items[*].status.podIP}'))
echo $POD_IPS
("10.42.0.73" "10.42.0.74" "10.42.0.75")

DELAY_MS=500
JITTER_MS=50
CONNECTION_LOSS_PERCENTAGE=10

# 1. Clear any existing config on the physical interface
sudo tc qdisc del dev eth0 root 2>/dev/null

# 2. Add the root priority qdisc
sudo tc qdisc add dev eth0 root handle 1: prio

# 3. Create the netem "Chaos Lane" (Handle 10: attached to Band 1:1)
sudo tc qdisc add dev eth0 parent 1:1 handle 10: netem delay "${DELAY_MS}ms" "${JITTER_MS}ms" loss "${CONNECTION_LOSS_PERCENTAGE}"

for s_ip in "${SAAS_IPS[@]}"; do
  for p_ip in "${POD_IPS[@]}"; do
    sudo iptables -t mangle -A POSTROUTING -s "$p_ip" -d "$s_ip" -j MARK --set-mark 1
  done
done

# 5. Direct marked traffic into the Chaos Lane
sudo tc filter add dev eth0 protocol ip parent 1:0 prio 1 handle 1 fw flowid 1:1
```

## Example: Squid Proxy running on an Ubuntu VM

```bash
# Tune these
DELAY_MS=500
JITTER_MS=50
LOSS_PERCENT=10
SQUID_USER=proxy
IFACE=eth0

# 1) clean slate
sudo tc qdisc del dev $IFACE root 2>/dev/null || true

# 2) root prio
sudo tc qdisc add dev $IFACE root handle 1: prio

# 3) chaos lane on band 1:1
sudo tc qdisc add dev $IFACE parent 1:1 handle 10: netem \
delay ${DELAY_MS}ms ${JITTER_MS}ms loss ${LOSS_PERCENT}%

# 4) normal lane on band 1:2
sudo tc qdisc add dev $IFACE parent 1:2 handle 20: fq_codel

# 5) marked packets -> chaos lane
sudo tc filter add dev $IFACE protocol ip parent 1:0 prio 1 handle 1 fw flowid 1:1

# 6) everything else -> normal lane (important to avoid surprises)
sudo tc filter add dev $IFACE protocol all parent 1:0 prio 10 matchall flowid 1:2

# 7) mark ONLY squid-owned outbound packets (so SSH etc stays clean)
# (Squid user is usually 'proxy')
sudo iptables -t mangle -A OUTPUT -o $IFACE -m owner --uid-owner $SQUID_USER -j MARK --set-mark 1
```

We can also reject requests completely:

```bash
sudo iptables -I OUTPUT 1 -o $IFACE -m owner --uid-owner $SQUID_USER -p tcp --dport 443 -j REJECT
```

## Example: Degraded + Blackhole

```bash
#!/usr/bin/env bash

# This script is used to simulate an unstable network environment and is meant to be run on a VM with a 
# Squid proxy running. It was run on EC2 instance therefore the interface is enX0.

#!/usr/bin/env bash
set -euo pipefail

IFACE="${IFACE:-enX0}"
SQUID_USER="${SQUID_USER:-proxy}"

GOOD_SECONDS="${GOOD_SECONDS:-20}"
BLACKHOLE_SECONDS="${BLACKHOLE_SECONDS:-10}"

GOOD_DELAY_MEAN_MS="${GOOD_DELAY_MEAN_MS:-300}"
GOOD_DELAY_JITTER_MS="${GOOD_DELAY_JITTER_MS:-200}"
GOOD_LOSS_PCT="${GOOD_LOSS_PCT:-1}"
GOOD_LOSS_CORR_PCT="${GOOD_LOSS_CORR_PCT:-25}"

cleanup() {
  echo "[INFO] cleaning up..."
  sudo iptables -t mangle -D OUTPUT -o "$IFACE" -m owner --uid-owner "$SQUID_USER" -j MARK --set-mark 1 2>/dev/null || true
  sudo tc qdisc del dev "$IFACE" root 2>/dev/null || true
}
trap cleanup EXIT

echo "[INFO] using IFACE=$IFACE SQUID_USER=$SQUID_USER"

# Root + lanes
sudo tc qdisc del dev "$IFACE" root 2>/dev/null || true
sudo tc qdisc add dev "$IFACE" root handle 1: prio
sudo tc qdisc add dev "$IFACE" parent 1:1 handle 10: netem delay 1ms 1ms loss 0%
sudo tc qdisc add dev "$IFACE" parent 1:2 handle 20: fq_codel

# Filters
sudo tc filter add dev "$IFACE" protocol ip parent 1:0 prio 1 handle 1 fw flowid 1:1
sudo tc filter add dev "$IFACE" protocol all parent 1:0 prio 10 matchall flowid 1:2

# Mark only Squid-owned egress
sudo iptables -t mangle -A OUTPUT -o "$IFACE" -m owner --uid-owner "$SQUID_USER" -j MARK --set-mark 1

echo "[INFO] starting volatility loop: ${GOOD_SECONDS}s degraded + ${BLACKHOLE_SECONDS}s blackhole..."
while true; do
  echo "[INFO] degraded window (${GOOD_SECONDS}s)"
  sudo tc qdisc change dev "$IFACE" parent 1:1 handle 10: netem \
    delay "${GOOD_DELAY_MEAN_MS}ms" "${GOOD_DELAY_JITTER_MS}ms" \
    loss "${GOOD_LOSS_PCT}%" "${GOOD_LOSS_CORR_PCT}%"
  sleep "$GOOD_SECONDS"

  echo "[INFO] blackhole window (${BLACKHOLE_SECONDS}s)"
  sudo tc qdisc change dev "$IFACE" parent 1:1 handle 10: netem loss 100%
  sleep "$BLACKHOLE_SECONDS"
done
```

## Modify Rules

```bash
sudo tc qdisc change dev enX0 parent 1:1 handle 10: netem delay 500ms 50ms loss 10%
```


**Note:** Adding this will destabilize the VM where this is run from.

## Remove Rules

```bash
sudo tc qdisc del dev $IFACE root
sudo iptables -t mangle -F POSTROUTING
sudo iptables -t mangle -D OUTPUT -o $IFACE -m owner --uid-owner proxy -j MARK --set-mark 1
sudo iptables -D OUTPUT -o $IFACE -m owner --uid-owner proxy -p tcp --dport 443 -j REJECT
sudo tc qdisc del dev $IFACE root 2>/dev/null || true
```

## See Applied Rules

```bash
sudo tc -s qdisc show dev $IFACE
sudo tc -s class show dev $IFACE 2>/dev/null || true
sudo iptables -t mangle -S OUTPUT | grep MARK
```

