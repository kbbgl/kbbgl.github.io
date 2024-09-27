# Wireshark Useful Filters

Sets a filter for any packet with `10.0.0.1`, as either the source or dest]

```text
ip.addr == 10.0.0.1 
```

Shows packets to and from any address in the `10.0.0.0/24` space

```text
ip.addr == 10.0.0.0/24
```

displays all packets that contain the word ‘traffic’. Excellent when searching on a specific string or user ID

```text
frame contains traffic
```

masks out arp, icmp, stp, or whatever other protocols may be background noise. Allowing you to focus on the traffic of interest

```text
!(arp or icmp or stp)
```

Sets a filter to display all http and dns protocols. It lets you narrow down to the exact protocol you need. So, if you need to track down an odd FTP traffic, then you just have to set it for ‘ftp’. Want to find out why some websites don’t appear? You just have to set it to ‘dns’.

```text
http or dns
```

Sets filters to display all TCP resets. All packets have a TCP, if this is set to 1, it tells the receiving computer that it should at once stop using that connection. So, this filter is a powerful one, being that a TCP reset kills a TCP connection immediately.

```text
tcp.flags.reset==1
```

## AMQP Filter

```text
amqp and !(frame.len==92)
```

### To search for text

Edit > Find Packet. Under "Find By:" select "string" and enter your search string in the text entry box

filter out heartbeats
