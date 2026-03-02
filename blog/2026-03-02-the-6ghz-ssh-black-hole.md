---
slug: 6ghz-ssh-isolation-troubleshooting
title: "The 6GHz Black Hole: Troubleshooting SSH Failures Across WiFi Bands"
authors: [kbbgl]
tags: [networking, macos, ssh, wifi6e]
date: 2026-03-02
---

Have you ever had a service that works perfectly on `localhost` but acts like it doesn't exist to the rest of your network? We recently spent an afternoon debugging a Mac Mini M4 that refused to accept SSH connections, despite every local check saying "All Systems Go."

The culprit wasn't a firewall or a wrong config, it was the **invisible wall between 5GHz and 6GHz WiFi bands**.

## The Setup
* **Target:** Mac Mini M4 (`192.168.x.10`) connected via a WiFi Extender on the **5GHz** band.
* **Host:** Laptop (`192.168.x.20`) connected directly to the main router on the **6GHz** (Wi-Fi 6E) band.
* **Goal:** Simple SSH access from Host to Target.

## Step 1: Is the service actually running?
First, we checked if the target was actually listening on Port 22.
```bash
sudo lsof -i :22
```

Result: `sshd` was listening on all interfaces (z). Testing `ssh localhost 22` worked perfectly. The "phone was off the hook," but the call wasn't getting through.

## Step 2: Checking the "Shields"
We checked the standard macOS suspects:
- Application Firewall: Turned OFF.
- Stealth Mode: Disabled.
- Third-Party Tools: None (clean install).

Despite this, a simple `nc -vz 192.168.x.10 22` from the host timed out, and ping failed entirely.

## Step 3: Following the Breadcrumbs (The Topology)

By looking at the router’s device list, we noticed a discrepancy:

- The Host was on the 6GHz band.
- The Target was on the 5GHz band (linked via an extender).

## The Theory: Why it Failed

When we disabled the 6GHz radio on the router, forcing the Host onto the 5GHz band, everything started working immediately. Why? There are three concrete theories for this behavior:

1. The ARP Bridge Failure (Layer 2): For the laptop to talk to the Mac Mini, it needs to know its MAC address via an ARP (Address Resolution Protocol) request. Many routers and extenders fail to "bridge" these broadcast requests across different physical radios (6GHz to 5GHz). If the ARP request never crosses the bridge, the Host has no "physical address" to send the SSH packets to.

2. Extender "Station Separation": WiFi extenders often use a technique called MAC Translation to manage clients. This can sometimes create a "one-way mirror" effect where the client on the extender can reach the internet, but unsolicited incoming traffic (like a New SSH Connection) from the main router's side is dropped because the extender doesn't recognize the route.

3. mDNS/Bonjour Isolation: macOS relies heavily on Bonjour (mDNS) for discovery. 6GHz is often treated as a "high-security" or "high-speed" isolated segment by consumer routers. If the router isn't explicitly configured to repeat Multicast traffic between the 6GHz and 5GHz radios, the devices remain invisible to each other.

## The Solution

If you encounter "Port 22 Timeout" on a modern Wi-Fi 6E/7 network:

1. Check your bands: Ensure both devices are on the same frequency (2.4, 5, or 6GHz).

2. Disable "Smart Connect" / SON: Give your 5GHz and 6GHz networks different names to control exactly where your "Server" devices land.

3. Wired Backhaul: If using an extender, connect it to the router via Ethernet or MoCA to avoid wireless bridging bugs.

The lesson? Just because they have the same IP subnet doesn't mean they're on the same "road."