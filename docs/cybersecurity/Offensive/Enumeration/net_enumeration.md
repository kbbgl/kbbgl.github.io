---
slug: net-enumeration-scripts
title: Network Enumeration Scripts
description: Network Enumeration Scripts
authors: [kbbgl]
tags: [cybersecurity,offensive,enumeration,network,bluetooth,wifi,dns,arp]
---

## ARP Scan

Scan whole subnet:

```bash
sudo arp-scan 192.168.1.0/24 --interface ens5
```

## DNS Server Enumeration

`dig` offers a way to gather DNS information about a target domain.

Get nameserver (which translates FQDN to IP addres

```bash
dig $HOST ns
```

Get email exchange server:

```bash
dig $HOST mx
```

## Host Scanning using `dmitry`

```bash
dmitry -winfsepo 10.100.102.1
```

## Inspecting Wireless Networks (WiFi and Bluetooth)

### WiFi

**AP** - Access Point. Physical device to where users connect to access internet.

**SSID** - Service Set Identifier. The name of the network.

**BSSID** - Base Service Set Identifier. Same as the MAC address of the device.

**ESSID** - Extended Service Set Identifier. Same as SSID but can be used for multiple APs.

**Frequency** - WiFi Bands are 2.4Ghz or 5GHz. 2.4Ghz gets farther but 5Ghz is faster.

**Channels** - Within the 2.4/5GHz bands there are smaller bands called channels. The channels are the medium by which data is sent and received. WiFi operates in 14 channels (in US 11) in 2.4GHz while 45 in 5GHz.

**Power** - measured in db or W. The closer physically to the AP, the greater the power, the easier it is to crack the connection. The range should not exceed 0.5W (around 100m)

**Security** - Security protocol used by the AP. The most common protocols are:

* Wired Equivalent Privacy (WEP) - Badly flawed and easily cracked.

* WiFi Protected Access (WPA) - Replaced WEP.

* WPA2-PSK - Most secure, uses preshared key (PSK) that all users share.

**Modes** - WiFi can operate in 3 modes:

* Managed - Ready to join or already joined AP.

* Master - Is an AP.

* Monitor - Scanning

#### Basic Wireless Commands

* Print information about wireless interface:

```bash
iwconfig
```

Wireless interfaces in Linux are usually called `wlanX0/1/2...`.

* Print all wireless interface:

```bash
 iwlist $interface $action
 iwlist wlan0 scan
```

This command supplies the following important information for hacking network:

* MAC address of target AP (BSSID).
* MAC address of client.
* Channel the AP is operating on.
  
* Daemon that provides high-level interface for net interfaces is the **Network Manager**. It has a cli which supplies more information than `iwlist`:

```bash
  # dev = devices
  nmcli $dev $networktype
  
  nmcli dev wifi
  
  # connect to AP
  nmcli dev wifi connect AP-SSID password AP-password
  
  nmcli dev wifi connect kgdev password 123456789
```

### Exploitation

Prerequesites:

* MAC address of target AP (BSSID).
* MAC address of client.
* Channel AP transmitting on.
* Password list.

`aircrack-ng` suite can supply this information.

To use `aircrack-ng`, we need to change the network interface into **monitor mode** to check all traffic going through device. We can do that using:

```bash
airmon-ng start wlan0
```

This will rename the wireless interface to `wlan0mon`.

Then run:

```bash
airodump-ng wlan0mon
```

This will supply:

* `BSSID` - MAC address of AP and client.
* `PWR` - signal strength (dbm)
* `ENC` - encryption used to secure transmission.
* `#Data` - data throughput rate.
* `CH` - channel.
* `ESSID` - name of the AP.

Need 3 terminals open running following commands to crack AP:

1. Capture all packets on channel 10:

  ```bash
  # -c $channel
  # -w SSID
  airodump-ng -c 10 --bssid 01:01:AA:BB:CC:22 -w kgdev wlan0mon
  ```

1. Deauthenticate anyone connected to AP so we can capture hash of the users that reauthenticate:

  ```bash
  aireplay-ng --deauth 100 -a 01:01:AA:BB:CC:22  wlan0mon
  ```

The password hash should show up in upper-right corner of `airodump-ng`.

1. Use password list to find password in captured hash:

  ```bash
  aircrack-ng -w wordlist.dic -b 01:01:AA:BB:CC:22 kgdev.cap
  ```

-------

## Bluetooth

NFC communication operating between 2.4 - 2.485GHz with 1.6k hops/sec. Range is on average ~10m but can get to 100m max.

Two Bluetooth devices can communicate if they are paired. They can be paired if they are both discoverable. When in discoverable mode, the device transmits:

* Name
* Class
* List of services
* Technical information

When two devices pair, they exchange a secret or link key. Each device stores this link key so it can identify the other in future pairing.

Each device has a unique 48-bit identifier (resembles MAC address) and usually a manufacturer-assigned name.

### Scanning and Reconnaissance

Linux has a Bluetooth stack implementation called `bluez`:

```bash
apt install bluez
```

`bluez` supplies a number of tools to manage and scan for Bluetooth devices:

* `hciconfig` - Similar to `ifconfig` but for Bluetooth devices.

* `hcitool` - Inquiry tool that can provide device name, device ID, device class and device clock information (which enables the devices to work synchronously)

* `hcidump` - Enables sniffing for the Bluetooth communicaton so data can be captured.

```bash
hciconfig 

# hci0: Type: BR/EDR Bus: USB BD Address: 10:AE:60:58:F1:37 ACL MTU: 310:10 SCO MTU: 64:8 
# UP RUNNING PSCAN INQUIRY 
# RX bytes:131433 acl:45 sco:0 events:10519 errors:0 
# TX bytes:42881 acl:45 sco:0 commands:5081 errors:0
```

Adapter name is `hci0` with MAC address `10:AE:60:58:F1:37`.

We can check whether the adapter is enabled by running:

```bash
hciconfig hci0 up
```

To scan for Bluetooth devices in discovery mode:

```bash
hcitool scan

# Scanning... 
# 72:6E:46:65:72:66 ANDROID BT 
# 22:C5:96:08:5D:32 SCH-I535
```

To gather more information about detected devices:

```bash
hcitool inq
# Inquiring... 
# 24:C6:96:08:5D:33 clock offset:0x4e8b class:0x5a020c 
# 76:6F:46:65:72:67 clock offset:0x21c0 class:0x5a020c
```

The `class` gives us information about what type of Bluetooth device it is (use [this link](http://domoticx.com/bluetooth-class-of-device-lijst-cod/)).

To scan for services:

```bash
sdptool browse $MAC_ADDRESS
```

**The device does not need to be in discovery mode** to be scanned.

To ping the devices:

```bash
l2ping $MAC_ADDRESS
```

## Get IP from Hostname

```bash
nslookup $HOSTNAME
```

## IP Ping Scan

```bash
#!/bin/bash

# If iplist.txt contains a list of IPs in each line:
# 192.168.1.1
# 192.168.1.2
# ...


# & runs each ping in different thread
# -c 1 -> run only once
for ip in $(cat iplist.txt): 
 do ping $ip -c 1 & 
done
```

## `massscan`

```bash
massscan -p1-65535 192.168.1.X
```

## `netdiscover`

```bash
sudo netdiscover 192.168.1.0/24 
```


## `nmap` Scan

### Scan TCP

```bash
nmap 10.10.10.4 -A -T4 -p-
```

### Scan UDP

```bash
nmap -sU 10.10.10.4 -A -T4 -p-
```

## Routing Tables (Linux)

```bash
show ip route
```

![routing table](https://geek-university.com/wp-content/images/ccna/ip_routing_example.jpg)

a network of two computers and a router. Host A wants to communicate with Host B. Because hosts are on different subnets, Host A sends its packet to the default gateway (the router). The router receives the packet, examines the destination IP address, and looks up into its routing table to figure out which interface the packet will be sent out

## Routing Tables (Windows)

```bash
route print
```

## SMB Server

Ports 445, 1433

We see:
```
| smb-security-mode: 
|   account_used: guest
|   authentication_level: user
|   challenge_response: supported
|_  message_signing: disabled (dangerous, but default)
```

Attempting connection to SMB server
```bash
# -L: list resources on the server
# -N: No password
smbclient -N -L \\\\10.10.10.27\\

smbclient -N \\\\10.10.10.27\\backups

# List directory contents

smb: \> dir

# Download file
smb: \> get prod.dtsConfig
```