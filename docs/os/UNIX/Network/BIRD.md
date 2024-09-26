# BIRD

[docs](https://bird.network.cz/)

Install:

```bash
sudo apt install bird
```

```bash
sudo birdc show status
```

Show router status, that is BIRD version, uptime and time from last reconfiguration.

```bash
sudo birdc show interfaces [summary]
```

Show the list of interfaces. For each interface, print its type, state, MTU and addresses assigned.

```bash
sudo birdc show protocols [all]
```

Show list of protocol instances along with tables they are connected to and protocol status, possibly giving verbose information, if all is specified.

```bash
sudo birdc show ospf interface [name] ["interface"]
```

Show detailed information about OSPF interfaces.

```bash
sudo birdc show ospf neighbors [name] ["interface"]
```

Show a list of OSPF neighbors and a state of adjacency to them.

```bash
sudo birdc show ospf state [all] [name]
```

Show detailed information about OSPF areas based on a content of the link-state database. It shows network topology, stub networks, aggregated networks and routers from other areas and external routes. The command shows information about reachable network nodes, use option all to show information about all network nodes in the link-state database.

```bash
sudo birdc show ospf topology [all] [name]
```

Show a topology of OSPF areas based on a content of the link-state database. It is just a stripped-down version of 'show ospf state'.

show ospf lsadb [global | area id | link] [type num] [lsid id] [self | router id] [name]
Show contents of an OSPF LSA database. Options could be used to filter entries.

show rip interfaces [name] ["interface"]
Show detailed information about RIP interfaces.

show rip neighbors [name] ["interface"]
Show a list of RIP neighbors and associated state.

show static [name]
Show detailed information about static routes.

show bfd sessions [name]
Show information about BFD sessions.

show symbols [table|filter|function|protocol|template|roa|symbol]
Show the list of symbols defined in the configuration (names of protocols, routing tables etc.).

show route [[for] prefix|IP] [table (t | all)] [filter f|where c] [(export|preexport|noexport) p] [protocol p] [(stats|count)] [options]
Show contents of specified routing tables, that is routes, their metrics and (in case the all switch is given) all their attributes.

You can specify a prefix if you want to print routes for a specific network. If you use for prefix or IP, you'll get the entry which will be used for forwarding of packets to the given destination. By default, all routes for each network are printed with the selected one at the top, unless primary is given in which case only the selected route is shown.

The show route command can process one or multiple routing tables. The set of selected tables is determined on three levels: First, tables can be explicitly selected by table switch, which could be used multiple times, all tables are specified by table all. Second, tables can be implicitly selected by channels or protocols that are arguments of several other switches (e.g., export, protocol). Last, the set of default tables is used: master4, master6 and each first table of any other network type.

You can also ask for printing only routes processed and accepted by a given filter (filter name or filter { filter } or matching a given condition (where condition).

The export, preexport and noexport switches ask for printing of routes that are exported to the specified protocol or channel. With preexport, the export filter of the channel is skipped. With noexport, routes rejected by the export filter are printed instead. Note that routes not exported for other reasons (e.g. secondary routes or routes imported from that protocol) are not printed even with noexport. These switches also imply that associated routing tables are selected instead of default ones.

You can also select just routes added by a specific protocol. protocol p. This switch also implies that associated routing tables are selected instead of default ones.

If BIRD is configured to keep filtered routes (see import keep filtered option), you can show them instead of routes by using filtered switch.

The stats switch requests showing of route statistics (the number of networks, number of routes before and after filtering). If you use count instead, only the statistics will be printed.

mrt dump table name|"pattern" to "filename" [filter f|where c]
Dump content of a routing table to a specified file in MRT table dump format. See MRT protocol for details.

configure [soft] ["config file"] [timeout [num]]
Reload configuration from a given file. BIRD will smoothly switch itself to the new configuration, protocols are reconfigured if possible, restarted otherwise. Changes in filters usually lead to restart of affected protocols.

If soft option is used, changes in filters does not cause BIRD to restart affected protocols, therefore already accepted routes (according to old filters) would be still propagated, but new routes would be processed according to the new filters.

If timeout option is used, config timer is activated. The new configuration could be either confirmed using configure confirm command, or it will be reverted to the old one when the config timer expires. This is useful for cases when reconfiguration breaks current routing and a router becomes inaccessible for an administrator. The config timeout expiration is equivalent to configure undo command. The timeout duration could be specified, default is 300 s.

configure confirm
Deactivate the config undo timer and therefore confirm the current configuration.

configure undo
Undo the last configuration change and smoothly switch back to the previous (stored) configuration. If the last configuration change was soft, the undo change is also soft. There is only one level of undo, but in some specific cases when several reconfiguration requests are given immediately in a row and the intermediate ones are skipped then the undo also skips them back.

configure check ["config file"]
Read and parse given config file, but do not use it. useful for checking syntactic and some semantic validity of an config file.

enable|disable|restart name|"pattern"|all
Enable, disable or restart a given protocol instance, instances matching the pattern or all instances.

reload [in|out] name|"pattern"|all
Reload a given protocol instance, that means re-import routes from the protocol instance and re-export preferred routes to the instance. If in or out options are used, the command is restricted to one direction (re-import or re-export).

This command is useful if appropriate filters have changed but the protocol instance was not restarted (or reloaded), therefore it still propagates the old set of routes. For example when configure soft command was used to change filters.

Re-export always succeeds, but re-import is protocol-dependent and might fail (for example, if BGP neighbor does not support route-refresh extension). In that case, re-export is also skipped. Note that for the pipe protocol, both directions are always reloaded together (in or out options are ignored in that case).

down
Shut BIRD down.

graceful restart
Shut BIRD down for graceful restart. See graceful restart section for details.

debug protocol|pattern|all all|off|{ states|routes|filters|events|packets [, ...] }
Control protocol debugging.

dump resources|sockets|interfaces|neighbors|attributes|routes|protocols
Dump contents of internal data structures to the debugging output.

echo all|off|{ list of log classes } [ buffer-size ]
Control echoing of log messages to the command-line output. See log option for a list of log classes.

eval expr
Evaluate given expression.
