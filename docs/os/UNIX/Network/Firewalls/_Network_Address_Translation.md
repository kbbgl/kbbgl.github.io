# Network Address Translation

Network Address Translation (NAT) allows for multiple network hosts to share the same external IP address. There are two types of outbound NAT or source NAT:

- `MASQUERADE`: Works with a dynamic source IP address. It is useful for servers with dynamic IP addresses.
- `SNAT`: Works with a static source IP address. It is less complex than `MASQUERADE`.

There is also a form of inbound or destination `NAT` (`DNAT`). `DNAT` allows for services to be behind a bastion host and to be easily load-balanced to different hosts.

To enable any of these types of NAT, the ip_forward kernel option must be set to 1.

```bash
echo 1 > /proc/sys/net/ipv4/ip_forward
```

It is also a good idea to make this change in the `sysctl.conf` file.

An example of a masquerade rule is the following:

```bash
iptables -t nat -A POSTROUTING -o eth1 ! -d 192.168.12.0/24 -j MASQUERADE
```
