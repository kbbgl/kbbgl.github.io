# System Init

`/sbin/init` is the first user process (pid=1) run on the system and runs until shutdown. All processes (aside from kernel-related once) are children to `init`.

`init` has a few jobs:

- Coordinates later stages of the boot process.
- Configures all aspects of the environment.
- Starts processes needed for logging into the system.
- Works closely with the kernel in cleaning up processes when they are terminated.

The popular schemes for system startup are:

- `Upstart` (phased out by Ubuntu in favor of `systemd`)
- `systemd` (most modern)
- `SysVinit` (for old, single processor computers)

## `systemd`

`systemd` is a software suite that provides an array of system components for Linux operating systems.

It provides utilities such as `systemctl` to manage services. The commands only control whether the service will be stopped/started during boot time.

```bash
# show all available services
systemctl list-units -t service --all

# show all active services
systemctl list-units -t service

# start/stop service
sudo systemctl start foo
sudo systemctl start foo.service
sudo systemctl start /path/to/foo.service
sudo systemctl stop foo

# enable/disable service
sudo systemctl enable sshd.service
sudo systemctl disable sshd.service
```

All services which run on start-up are listed in:

```bash
ls /etc/init.d
```

To check for a particular service status:

```bash
sudo service network-manager status
sudo service --status-all
```
