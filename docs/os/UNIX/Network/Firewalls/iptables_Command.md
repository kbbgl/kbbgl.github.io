# `iptables` Command

The `iptables` command can be broken into multiple pieces:

- The table using the `-t` switch. If no table is specified, the default is `filter`.
- The command, which is one of the following:
`'-I' (insert)`
- Create a new rule at the top of the chain.
`'-I #' (insert)`
- Create a new rule at position # of the chain.
`'-A' (append)`
- Create a new rule at the bottom of the chain.
`'-P' (policy)`
- Change the chain's policy.
`'-D' (delete)`
- Delete a rule.
`'-D #' (delete)`
- Delete rule number #.
  - The chain name.
  - The match criteria.
  - The target.

An example rule is as follows:

```bash
iptables -t filter -A INPUT -m tcp -p tcp --dport 22 -j ACCEPT
```

This rule appends the rule to the bottom of the `INPUT` chain, loads the `tcp` module, matches the TCP protocol destination port 22 and jumps to the `ACCEPT` target.

Here is another example:

```bash
iptables -t <TABLE> <command> <CHAIN> <match criteria> -j <TARGET>
```

Insert a new rule at the top of a chain:

```bash
iptables -I INPUT -m udp -p udp --dport 53 -j ACCEPT
```

Set the INPUT chain policy to DROP:

```bash
iptables -P INPUT DROP
```

Delete the third rule from the INPUT chain:

```bash
iptables -D INPUT 3
```

## Match Criteria

The match criteria is the essence of the iptables system, and allows for a lot of flexibility. By default, if you use the `-p` or `--protocol` switch, the corresponding module is loaded.

- `state`
Allows for matching for stateful firewalls.
- `tcp`
Allows for matches on the TCP information (source port, destination port, tcp flags).
- `udp`
Allows for matches on UDP information (source port, destination port).
- `icmp`
Allows for matches on ICMP query types.

## Targets

- `ACCEPT`: Pass the packet along to the next stage.
- `DROP`: Send no response to this packet and ignore it.
- `RETURN`: Go back to the calling CHAIN and start processing on the next rule.
- `REJECT`: Send a message back explaining why the packet is not allowed.

## Distribution Defaults

Each distribution has its own method for storing the iptables state. Some make it easy to manage it with generic tools, while some provide easy-to-use tools for firewall management.

CentOS:
CentOS 7 has a package called `iptables-services` which provides the historical program service and `systemd` with scripts to assist in managing `iptables`.
Firewall rules are saved in `/etc/sysconfig/iptables`.
Easiest management is done with `iptables` and service `iptables` save.

OpenSUSE:
OpenSUSE has a standalone firewall management tool called `SuSEfirewall2` that also works in conjunction with `systemd`. This solution requires an administrator to save the configuration, but automates restarting the firewall at boot time.
Firewall rules are saved in `/etc/sysconfig/scripts/SuSEfirewall2-*`.
Easiest management is done with `yast` and `SuSEfirewall2`.

Ubuntu:
Ubuntu has two packages, `iptables-persistent` and `netfilter-persistent`, which are available to make the `iptables` rules persistent by using `systemd`. The default location is `/etc/iptables/rules.v4`.
The default firewall is managed with `ufw`.
Custom rules are stored in `/etc/ufw/*.rules`.
