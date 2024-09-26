# netfilter Vocabulary

`netfilter` is a packet-filtering framework built into the Linux kernel. To better understand `netfilter`, we need to start with some vocabulary:

- The netfilter firewall consists of tables.
- Tables consist of chains.
- Chains have a default policy.
- Chains consist of rules.
- Rules consist of a match criteria and a target.

Rules in each chain are read first to last, and the first match wins. If a packet does not match any rule in a chain, the policy of the chain applies.

## Configuration Utilities

Because of the inherent complexity of the `netfilter` management, many tools have been created to help alleviate the burden on systems administrators. The CentOS `system-config-firewall`, the Ubuntu `gufw`, and the OpenSUSE `yast` firewall are examples of configuration utilities.  

Each distribution has its own GUI or TUI mode firewall tool.

There are also generic tools like shorewall, which wrap the complexity of iptables/netfilter in an API.

The latest addition to firewall management is `firewalld`. The `firewalld` tool is available in most of the recent distributions.

## `netfilter` Hooks

`nftables` is an update to packet filtering for the `netfilter` framework. `nftables` replaces:

- `iptables`
- `ip6tables`
-`arptables`
- `ebtables`

Most of the concepts from `netfilter` apply to `nftables`. One major change is the hooks into the network stack. In `netfilter` the hooks are pre-defined and known as `filter`, `raw`, `nat`, etc.

In `nftables` the hooks are not connected by default. The hooks have to be connected to chains by the administrator, as they are not pre-configured.

There are hooks for several types of packets such as:

- `ip` - IPv4 address family
- `ip6` - IPv6 address family
- `inet` - Internet (IPv4/IPv6) address family
- `arp` - ARP family, handling IPv4 ARP
- `bridge` - Handles packets transversing a bridge device.
- `netdev` - netdev address family, handling packets from ingress.

## `nftables` Configuration Structure

Once the requirements are known, the configuration can be added, building the structure as needed. In our example on the next page, the requirements are sparse and intended only as examples. You could add the configuration in a number of ways:

Using `nft` commands:
The commands can be issued one at a time from the command line.

Using `nft` shell:
A non-interactive shell is available that will process through a file. The format of the file is the same as the nft commands with the nft removed.

Using `nft -f filename`:
An atomic method of changing the `nftables` configuration. The command uses an input file formatted like the output of the list command. Here is an example of save, flush, restore, a table configuration:

```bash
nft list table inet filter_it > nft.conf
nft flush ruleset
nft -f nft.conf
```

## `nft` Administration Interfaces

One of the common firewall configuration programs, `firewalld`, is compatible with `nftables`.

To switch from `iptables` to `nftables` set the option `FirewallBackend=nftables` in the `/etc/firewalld/firewalld/conf` file. Most distributions have it compiled by default for `firewalld` set to `iptables` and the option may not appear in the file.

With the support of `nftables` in `firewalld`, the interface options include a GUI and a CLI beyond the basic `nft` command.

## `nft` File Locations

It's very important to know that locations may change, so make sure to check your distribution.

The key file locations for `nft` in the Fedora 30 release are the `systemd` top level configuration file `/usr/lib/systemd/system/nftables.service`, which points to the `nft` configuration file, `/etc/sysconfig/nftables.conf` file, that includes the configuration files in the directory `/etc/nftables/*.nft`.
