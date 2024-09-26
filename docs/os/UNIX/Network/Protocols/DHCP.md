# Dynamic Host Configuration Protocol (DHCP) Server

The Dynamic Host Configuration Protocol (DHCP) is used to configure the network-layer addressing. The `dhcpd` daemon used to be configured using both a configuration file (`/etc/dhcp/dhcpd.conf`) and a daemon options file that was distribution-dependent. Recent versions of `dhcp` have moved the daemon options into `systemd`.

The daemon options are configured in a separate file:

## On CentOS

```bash
/etc/sysconfig/dhcpd
```

## On Ubuntu

```bash
/etc/default/isc-dhcp-server
```

## On OpenSUSE

```bash
/etc/sysconfig/dhcpd
```

The `dhcp` server will only serve out addresses on an interface that it finds a subnet block defined in the `/etc/dhcp/dhcpd.conf` file. It is no longer a requirement to explicitly tell dhcp which interfaces to use.

Additional or different daemon command line options may be passed to the daemon at start time by the systems' drop-in files. Please see the `COMMAND LINE` section in the `dhcpd man` page for additional details.

## Configuration

Global options are settings which should apply to all the hosts in a network. You can also define options on a per-network basis.

A sample configuration would be:

```bash
subnet 10.5.5.0 netmask 255.255.255.224 {
  range 10.5.5.26 10.5.5.30;
  option domain-name-servers ns1.internal.example.org;
  option domain-name "internal.example.org";
  option routers 10.5.5.1;
  option broadcast-address 10.5.5.31;
  default-lease-time 600;
  max-lease-time 7200;
  }
```
