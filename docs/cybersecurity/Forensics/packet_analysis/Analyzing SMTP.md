---
slug: analyzing-smtp
title: Analyzing SMTP
description: Fill me up!
authors: [kbbgl]
tags: [pcap,smtp,tshark,wireshark]
---

Taken from _Network Forensics, Tracking Hackers through Cyberspace_, Case Study 4.6.

## Description

After being released on bail, Ann Dercover disappears! Fortunately, investigators were carefully monitoring her network activity before she skipped town. “We believe Ann may have communicated with her secret lover, Mr. X, before she left,” says the police chief. “The packet capture may contain clues to her whereabouts.”

## Mission
Your mission is to analyze the packet capture and gather information about Ann’s activities and plans.

The following questions will help guide your investigation:
• Provide any online aliases or addresses and corresponding account credentials that may be used by the suspect under investigation.

• Who did Ann communicate with? Provide a list of email addresses and any other identifying information.

• Extract any transcripts of Ann’s conversations and present them to investigators.

• If Ann transferred or received any files of interest, recover them.

• Are there any indications of Ann’s physical whereabouts? If so, provide supporting evidence.

## Network


• Internal network: `192.168.30.0/24`

• DMZ: `10.30.30.0/24`

• The “Internet”: `172.30.1.0/24` [Note that for the purposes of this case study, we are treating the `172.30.1.0/24` subnet as “the Internet.” In real life, this is a reserved nonroutable IP address space.]

## Evidence

Investigators provide you with a packet capture from Ann’s home network, [`evidence-packet-analysis.pcap`](./evidence-packet-analysis.pcap) They also inform you that in the course of their monitoring, they have found that Ann’s laptop has the MAC address `00:21:70:4D:4F:AE`.


## Analysis

### Protocol Summary

Let's first review the protocols used in this packet capture.

```bash
tshark -n  -r evidence-packet-analysis.pcap -q -z io,ph
```

```
===================================================================
Protocol Hierarchy Statistics
Filter: 

eth                                      frames:2487 bytes:1508942
  ip                                     frames:2487 bytes:1508942
    udp                                  frames:481 bytes:53801
      ntp                                frames:17 bytes:1530
      dhcp                               frames:10 bytes:3509
      nbdgm                              frames:18 bytes:3888
        smb                              frames:18 bytes:3888
          mailslot                       frames:18 bytes:3888
            browser                      frames:18 bytes:3888
      nbns                               frames:108 bytes:10584
      dns                                frames:328 bytes:34290
    tcp                                  frames:2000 bytes:1454781
      http                               frames:111 bytes:71429
        data-text-lines                  frames:18 bytes:15217
          tcp.segments                   frames:11 bytes:10088
        media                            frames:2 bytes:1490
          tcp.segments                   frames:1 bytes:694
        image-gif                        frames:4 bytes:2205
          tcp.segments                   frames:1 bytes:995
        png                              frames:23 bytes:14320
          tcp.segments                   frames:18 bytes:8640
          _ws.malformed                  frames:2 bytes:2251
        image-jfif                       frames:6 bytes:4737
          tcp.segments                   frames:6 bytes:4737
      smtp                               frames:269 bytes:303146
        imf                              frames:3 bytes:180
      imap                               frames:291 bytes:307021
        imf                              frames:2 bytes:1189
        tcp.segments                     frames:202 bytes:284153
    igmp                                 frames:6 bytes:360
===================================================================
```

All ethernet (layer 2) and IP (layer 3) communication.

We can see that DHCP communication occurred. It can help us link the provided MAC address to the assigned IP address. 


As a reminder, this is how the DHCP request/response works:

![dhcp-req-res](https://www.researchgate.net/profile/Chi-Fu-Huang/publication/4109437/figure/fig8/AS:670523923103777@1536876769687/llustrates-the-related-DHCP-message-flows-The-DHCP-Discover-is-forwarded-by-the-DHCP.ppm)


Let's list all DHCP communication in the packet capture:

```bash
❯ tshark -r evidence-packet-analysis.pcap -Y "eth.addr == 00:21:70:4d:4f:ae and bootp" -V
```

We can see in frame 2 that the DHCP message is sent to request a local IP address `192.168.30.108`:
```
Frame 2: 352 bytes on wire (2816 bits), 352 bytes captured (2816 bits)
    ...
Dynamic Host Configuration Protocol (Request)
    Message type: Boot Request (1)
    Option: (53) DHCP Message Type (Request)
        Length: 1
        DHCP: Request (3)
    Option: (61) Client identifier
        Length: 7
        Hardware type: Ethernet (0x01)
        Client MAC address: Dell_4d:4f:ae (00:21:70:4d:4f:ae)
    Option: (50) Requested IP Address (192.168.30.108)
        Length: 4
        Requested IP Address: 192.168.30.108
    Option: (12) Host Name
        Length: 10
        Host Name: ann-laptop
```

Then, in frame 5, we see the reply/acknowledgement for the requeted IP address, the DNS server (`10.30.30.20`) and the DHCP server (`192.168.30.108`)

```
Frame 5: 358 bytes on wire (2864 bits), 358 bytes captured (2864 bits)

Dynamic Host Configuration Protocol (ACK)
    Message type: Boot Reply (2)
    Option: (53) DHCP Message Type (ACK)
        Length: 1
        DHCP: ACK (5)
    Option: (54) DHCP Server Identifier (192.168.30.10)
        Length: 4
        DHCP Server Identifier: 192.168.30.10
    Option: (51) IP Address Lease Time
        Length: 4
        IP Address Lease Time: (3600s) 1 hour
    Option: (58) Renewal Time Value
        Length: 4
        Renewal Time Value: (1800s) 30 minutes
    Option: (59) Rebinding Time Value
        Length: 4
        Rebinding Time Value: (3150s) 52 minutes, 30 seconds
    Option: (81) Client Fully Qualified Domain Name
        Length: 25
        Flags: 0x00
            0000 .... = Reserved flags: 0x0
            .... 0... = Server DDNS: Some server updates
            .... .0.. = Encoding: ASCII encoding
            .... ..0. = Server overrides: No override
            .... ...0 = Server: Client
        A-RR result: 0
        PTR-RR result: 0
        Client name: ann-laptop.example.com
    Option: (1) Subnet Mask (255.255.255.0)
        Length: 4
        Subnet Mask: 255.255.255.0
    Option: (6) Domain Name Server
        Length: 4
        Domain Name Server: 10.30.30.20
    Option: (46) NetBIOS over TCP/IP Node Type
        Length: 1
        NetBIOS over TCP/IP Node Type: H-node (8)
    Option: (3) Router
        Length: 4
        Router: 192.168.30.10
    Option: (255) End
        Option End: 255
```

### Keyword Search

A common way to find any information is to perform some search for keywords. Since we know the targets name is 'Ann Dercover', we can search for it within the packet capture.

```bash
❯ ngrep "Ann" -N -t -q -I evidence-packet-analysis.pcap
```

We can see that Ann sent message to get a fake passport which was split into a couple of messages (output was cleaned up):

```
T(6) 2011/05/17 22:33:08.648555 192.168.30.108:1685 -> 205.188.58.10:143 [AP] #1615
  From: "Ann Dercover" <sneakyg33ky@aol.com>
  To: <inter0pt1c@aol.com>
  Subject: need a favor
  Date: Tue, 17 May 2011 13:32:17 -0600
  Hey, can you hook me up quick with that fake passport you were taking 
  about? - Ann

  T(6) 2011/05/17 22:34:16.481132 192.168.30.108:1687 -> 64.12.168.40:587 [A] #1749
  From: "Ann Dercover" <sneakyg33ky@aol.com>
  To: <d4rktangent@gmail.com>
  Subject: lunch next week
  Date: Tue, 17 May 2011 13:33:26 -0600
  Sorry-- I can't do lunch next week after all. Heading out of town. 
  Another time!
  
  -Ann

T(6) 2011/05/17 22:35:16.962873 192.168.30.108:1689 -> 64.12.168.40:587 [A] #1825
  From: "Ann Dercover" <sneakyg33ky@aol.com>
  To: <mistersekritx@aol.com>
  Subject: rendezvous
  Date: Tue, 17 May 2011 13:34:26 -0600
  
  Hi sweetheart! Bring your fake passport and a bathing suit. Address 
  attached. love, Ann
```

And we see there's an email attachment called `secretrendezvous.docx` and the base64 representation of this file:
```
T(6) 2011/05/17 22:35:16.963123 192.168.30.108:1689 -> 64.12.168.40:587 [A] #1826
  it. Address attached. love, Ann</FONT></DIV></BODY></HTML>
  Content-Type: application/octet-stream;
  name="secretrendezvous.docx"
  Content-Transfer-Encoding: base64
  Content-Disposition: attachment;
  filename="secretrendezvous.docx"

  UEsDBBQABgAIAAAAIQCht/xGcgEAAFIFAAATAAgCW0NvbnRlbnRfVHlwZXNdLnhtbCCiBAIooAAC
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
....
```


### Carving Out Attachment

Putting this base64 payload into a file `evidence-packet-analysis-smtp3-attachment.b64.enc`

We can take the base64 payload and decrypt it into a Microsoft Word document format:

```bash
❯ base64 -d evidence-packet-analysis-smtp3-attachment.b64.enc > secretrendezvous.docx

❯ file secretrendezvous.docx
secretrendezvous.docx: Microsoft Word 2007+
```

Opening the document, we can see that it has an image in png format. We can also carve it out:

```bash
❯ unzip secretrendezvous.docx
Archive:  secretrendezvous.docx
  inflating: [Content_Types].xml
  inflating: _rels/.rels
  inflating: word/_rels/document.xml.rels
  inflating: word/document.xml
 extracting: word/media/image1.png
  inflating: word/theme/theme1.xml
  inflating: word/settings.xml
  inflating: word/webSettings.xml
  inflating: docProps/core.xml
  inflating: word/styles.xml
  inflating: word/fontTable.xml
  inflating: docProps/app.xml
❯ open word/media/image1.png
```

