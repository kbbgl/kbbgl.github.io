# Routing

We can get the routing list of IP to device/interface:

```bash
route -n
ip route
```

The default route is the route taken when packets destinations are not specified.

The default route can be set:

```bash
sudo route add default gw 192.168.1.1 ens192
```

## Static Routes

Static routes are used to control packet flow when there is more than one route and are defined for each interface. They can be persistent or non-persistent.

To create a non-persistent route:

```bash
sudo ip route add 10.50.0.0/16 via 192.168.1.100
```

To create a persistent route, we need to add lines to `/etc/network/interfaces`:

```bash
iface eth1 inet dhcp

post-up route add -host 10.1.2.51 eth1
post-up route add -host 10.1.2.52 eth1
```
