---
slug: os-linux-network-troubleshooting-get-interface-data-throughput-tx-rx
title: "Get Network Interface Throughput"
authors: [kbbgl]
tags: [os, linux, network, troubleshooting, get_interface_data_throughput_tx_rx]
---

# Get Network Interface Throughput

```bash
sar -n DEV 1 3

Linux 4.19.118-Re4son-v7l+ (kali)  08/25/20  _armv7l_ (4 CPU)

16:38:41        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
16:38:42      docker0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:42           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:42         eth0      1.00      1.00      0.06      0.16      0.00      0.00      0.00      0.00
16:38:42        wlan0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

16:38:42        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
16:38:43      docker0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:43           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:43         eth0     11.00     11.00      0.71      2.20      0.00      0.00      0.00      0.00
16:38:43        wlan0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

16:38:43        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
16:38:44      docker0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:44           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
16:38:44         eth0     11.00     11.00      0.71      2.24      0.00      0.00      0.00      0.00
16:38:44        wlan0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00

Average:        IFACE   rxpck/s   txpck/s    rxkB/s    txkB/s   rxcmp/s   txcmp/s  rxmcst/s   %ifutil
Average:      docker0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:           lo      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
Average:         eth0      7.67      7.67      0.49      1.53      0.00      0.00      0.00      0.00
Average:        wlan0      0.00      0.00      0.00      0.00      0.00      0.00      0.00      0.00
```
