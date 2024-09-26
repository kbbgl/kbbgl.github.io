# AppArmor

AppArmor is an LSM alternative to SELinux. Support for it has been incorporated in the Linux kernel since 2006. It has been used by SUSE, Ubuntu and other distributions.

AppArmor:

- Provides Mandatory Access Control (MAC)
- Allows administrators to associate a security profile to a program which restricts its capabilities
- Is considered easier (by some but not all) to use than SELinux
- Is considered filesystem-neutral (no security labels required).

AppArmor supplements the traditional UNIX Discretionary Access Control (DAC) model by providing Mandatory Access Control (MAC).

In addition to manually specifying profiles, AppArmor includes a learning mode, in which violations of the profile are logged, but not prevented. This log can then be turned into a profile, based on the program's typical behavior.

## Interacting with `apparmor`

Disabling/Enabling service

```bash
sudo systemctl [start|stop|restart|status] apparmor
sudo systemctl [enable|disable] apparmor
```

To see status:

```bash
sudo apparmor_status
```

## Modes and Profiles

Profiles restrict how executable programs, which have pathnames on your system, such as /usr/bin/evince, can be used.

Processes can be run in either of the two modes:

- **Enforce Mode**: Applications are prevented from acting in ways which are restricted. Attempted violations are reported to the system logging files. This is the default mode. A profile can be set to this mode with aa-enforce.
- **Complain Mode**: Policies are not enforced, but attempted policy violations are reported. This is also called the learning mode. A profile can be set to this mode with aa-complain.

Linux distributions come with pre-packaged profiles, typically installed either when a given package is installed, or with an AppArmor package, such as apparmor-profiles. These profiles are stored in `/etc/apparmor.d`.

For full docs:

```bash
man apparmor.d
```

## Utilities

```bash
rpm -qil apparmor-utils | grep bin​
/usr/bin/aa-easyprof
/usr/sbin/aa-audit
/usr/sbin/aa-autodep
/usr/​sbin/aa-cleanprof
/usr/sbin/aa-complain
/usr/sbin/aa-decode
/usr/sbin/aa-disable
/usr/sbin/aa-enforce
/usr/sbin/aa-exec
```

```bash
# shows status of all profiles and processes with profiles
apparmor_status

# show summary for AppArmor log messages
apparmor_notify

# set a specified profile to complain/enforce mode
complain
enforce

# unload specified profile from current kernel and prevent being loaded on system startup
disable

# scan log files
logprof

# help set up a basic apparmor profile for a program
easyprof
```
