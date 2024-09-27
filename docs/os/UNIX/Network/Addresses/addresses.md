# Network Addresses

IP addresses uniquely identify nodes across the internet. They are registered using an ISP.

There are two versions of IP address:

- IPv4: 32-bit address composed of 4 octets, `148.114.252.10`
- IPv6: 128-bit address composed of 8 16-bit octet pairs, `2003:0db5:6123:0000:1f4f:0000:5529:fe23`

## IPv4 Address Types

- **Unicast**: An address associated with a specific host.
- **Network**: An address whos host portion is set to all binary zeroes, e.g. `192.168.1.0`.
- **Broadcast**: An address to which each member of a particular network will listen. Will have the host portion set to all 1 bits, e.g. `172.16.255.255`.
- **Multicast**: An address to which appropriately configured nodes will listen, e.g. `224.0.0.2`. Only nodes specifically configured to pay attention to a specific multicast address will interpret packets for that multicast group.

## Reserved Address

- `127.x.x.x`: loopback address.
- `0.0.0.0`: address when attempting to communicate with server.
- `255.255.255.255`: General broadcast private address.
- `10.x.x.x`, `172.16.y.y`, `192.168.z.z`.

## IPv4 Address Classes

![v3](http://35.245.49.226/wp-content/uploads/2014/12/IPV4addressclass.jpg)

Netmasks are used to determine how much of the address is used for the network portion and host portion.

![subnet](https://www.tutorialspoint.com/ipv4/images/class_a_subnets.jpg)

## Managing Hostname

The hostname is a label to identify a networked device. For DNS, hostnames are appended with periods to create the FQDN.

To set the hostname for the session only:

```bash
sudo hostname myhost
```

To modify the hostname permanently:

```bash
sudo hostnamectl set-hostname myhost
```

The value is stored in:

```text
/etc/hostname
```
