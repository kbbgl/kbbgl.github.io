# Cron Job Scheduling

`crond` is a daemon that runs in the bg. it checks the `crontab` (table) for jobs that need to be scheduled according to jobs configured in `/etc/crontab`.

![crontab](https://ahmadawais.com/wp-content/uploads/2017/06/crontab.png)

To edit the table:

```bash
crontab -e
```

There are also useful shortcuts:

```text
@reboot Run once, at startup.
@yearly Run once a year (replaces: 0 0 1 1 *)
@monthly Run once a month (replaces: 0 0 1 * *)
@weekly Run once a week (replaces: 0 0 * * 0)
@daily Run once a day (replaces: 0 0 * * *)
@hourly Run once an hour (replaces: 0 * * * *)
```

## Startup Scripts

When the system is started, after the kernel loads all modules, the kernel starts a daemon called `init` or `initd` in `pid=1`. The daemon runs scripts found in `/etc/init.d/rc` which are necessary for system service startup.

`/etc/init` is where the upstart init configs live. While they are not scripts themselves, they essentially execute whatever is required to replace sysvinit scripts.

`/etc/init.d` is where all the traditional sysvinit scripts and the backward compatible scripts for upstart live. The backward compatible scripts basically run service myservice start instead of doing anything themselves. Some just show a notice to use the "service" command.

`/etc/init/rc-sysinit.conf` controls execution of traditional scripts added manually or with `update-rc.d` to traditional runlevels in `/etc/rc*`

`/etc/default` has configuration files allowing you to control the behaviour of both traditional `sysvinit` scripts and new upstart configs.

We can add services to run at startup by using:

```bash
update-rc.d $NAME_OF_SCRIPT_SVC remove|defaults|enable|disable

# To add a script
sudo mv script.sh /etc/init.d
update-rc.d script.sh defaults
```

## Runlevels

A `runlevel` is the mode of operation which defines the state of the machine after boot.
There are 7 levels:
![nn](https://www.linuxvasanth.com/wp-content/uploads/2019/01/NN.jpg)

When loading a server without GUI, the `runlevel` is usually 3 and with GUI it's 5. When the server is rebooted, it enters `runlevel` 6.

To check the `runlevel`:

```bash
runlevel
# N 5
```
