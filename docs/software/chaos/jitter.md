---
slug: add-flaky-network
title: How to Add Network Flakiness
description: 
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

The first step is to create a root queueing discipline (A.K.A. `qdisc`). Whenever the kernel needs to send a packet to an interface, it is enqueued to the `qdisc` configured for that interface. Immediately afterwards, the kernel tries to get as  many  packets  as possible from the `qdisc`, for giving them to the network adaptor driver.

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

**Note:** Adding this will destabilize the VM where this is run from.

To remove all rules:

```bash
sudo tc qdisc del dev cni root
```