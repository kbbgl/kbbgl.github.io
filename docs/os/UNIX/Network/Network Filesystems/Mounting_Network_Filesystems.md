# Mounting Network Filesystems

The three methods of mounting network filesystems (immediate mount -command line-; always mounted; and mounted on-demand) use a common configuration file `/etc/fstab`. Over the years, additional features and services have updated the options in `/etc/fstab`. Functionality that required external packages, services and configuration are now combined under `systemd` services.

The `/etc/fstab` has slightly changed its purpose with `systemd` from configuration file to drive the `mount` command during system initialization to an input configuration file for the `systemd-fstab-generator` that creates native unit files used by `systemd.mount` during system startup. The regular mount command still looks in the `/etc/fstab` file.

Since the preferred method to configure file system mounts is still `/etc/fstab`, nothing has changed for most `mount` options. See `man systemd.mount` for additional details.

## `mount` Command

The universal mounting command can mount many types of filesystems, including NFS and cifs filesystems.

- If the device name is in the format `SERVERNAME:/share`, then `NFS` is assumed.
- If the device name is in the format `SERVERNAME//share`, then `cifs` is assumed.
- In most cases, the `-t fstype` is optional.
- Supports options specific to the filesystem, like the `cifs` username and `do-` main options.

## Persistent Mounting Network File Systems

Network filesystems can be persistently mounted through `/etc/fstab`:

- the `fs_vfstype` field must be set, usually to one of the following: `nfs`, `cifs` or `nfs4`
- the `fs_mntopts` field can contain options specific the network filesystem
- the `fs_freq` field is generally set to `0`
- the `fs_passno` is generally set to `0`; setting this field to a non-zero value may cause a delay during system startup time waiting for the mount to complete.

## Automount Network File Systems

The automount will monitor a configured mount point and, if the information in the filesystem is accessed, the mount will be executed.

There may be an option to disconnect an idle connection.

- `autofs`
  - Requires additional packages to be installed.
  - Multi-part configuration files.
  - May require scripts as helpers for mounting some filesystem.

- `systemd.automount`
  - Integrated into `systemd`.
  - Only available on `systemd` distributions.
  - Unified configuration files, uses `/etc/fstab`.
  - Uses common `systemd.mount` facilities, no scripts, but allows unit file overrides.

With the `systemd.automount` implementation, the actual unit configuration files are generated from the entries in `/etc/fstab`. There are options to add to the `/etc/fstab` to indicate to the `systemd-fstab-generator` the entries are to be automounted. The `systemd` unit files that are generated are located in the `/run/systemd/generator` directory. The file names end in `mount` and `automount`. Like other `systemd` files, they can be overridden by entries in the `/etc/systemd/system` directory, but it is easier to just change the options in `/etc/fstab`. The `man` page for `systemd.mount` has information on the options. Some examples from `/etc/fstab` are provided below:

```bash
127.0.0.1:/home/export/nfs /home/share/nfs nfs x-systemd.automount,x-systemd.idle-timeout=10,noauto,_netdev 0 0 //localhost/cifs-share /home/share/cifs cifs creds=/root/smbfile,x-systemd.automount,x-systemd.idle-timeout=10,noauto,_netdev 0 0
```
