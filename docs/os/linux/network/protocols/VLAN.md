---
slug: os-linux-network-protocols-vlan
title: "Virtual Local Area Network"
authors: [kbbgl]
tags: [os, linux, network, protocols, vlan]
---

# Virtual Local Area Network

VLANs use functionality in the switches and routers. The switch or router functionality must exist in the device as it is not usually an option that be added later.

The most common use for VLANs is to bridge switches together. Creating a trunk connection between two switches essentially connects the networks together. VLANs can be created on the trunked together switches to then isolate specific ports on each switch to belong to a Virtual LAN.

VLANs use optional functionality within the packet to identify which VLANs are being used.

## Packet Attributes

A Virtual Local Area Network (VLAN) allows multiple network nodes to end up in the same broadcast domain, even if they are not plugged into the same switch or switches. A VLAN is also a method for securing two or more LANs from each other on the same set of switches.

VLANs can be linked from point to point by doing VLAN trunking (802.1q is one such protocol).

When VLANs are enabled, an additional header is added to the packet. This is the 802.1Q or dot1q header. The 802.1Q header contains:

Tag Protocol ID (TPID), a 16bit identifier to distinguish between a VLAN tagged frame or it indicates the EtherType. The value of x8100 indicates an 802.1Q tagged frame.
The next 16bits are the Tag Control Information (TCI), comprising of:

- Priority Code Point (PCP), a 3bit field that indicates the 802.1q priority class. See the IEEE P802.1Q webpage for more details.
- Drop Eligible Indicator (DEI). This 1bit flag indicates the frame may be dropped when network congestion occurs. The field may be used alone or in conjunction with the PCP. This field used to be the Canonical Format Indicator (CFI), and was used for compatibility between Ethernet and Token Ring frames.
- VLAN Identifier ID (VID), a 12bit field indicating to which VLAN the frame belongs to.
Additional information can be found on the IEEE 802.1Q webpage.
