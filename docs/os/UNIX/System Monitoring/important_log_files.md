# System and Kernel Log Files

## System - `/var/log/sys`

* non-kernel boot errors
* application-related service errors and the messages that are logged during system startup.

## Kernel - `/var/log/kern.log`

* Perfect for troubleshooting kernel related errors and warnings.

* Can also come handy in debugging hardware and connectivity issues.

## `journalctl`

check logs one hour ago:

```bash
journalctl --since "1 hour ago"
```

Linux uses a daemon named `syslogd` to log events on server. There are several other variations but the most popular Debian-based is called `rsyslog`.

The configuration file is in `/etc/rsyslog.conf`.

The configuration file includes a section called _Rules_ where we can see the following:

```text
#### RULES ####

# auth,authpriv.* /var/log/auth.log 
*.*;auth,authpriv.none -/var/log/syslog 
#cron.* /var/log/cron.log 
daemon.* -/var/log/daemon.log 
kern.*  -/var/log/kern.log 
1pr.* -/var/log/lpr.log 
mail.* -/var/log/mail.log 
user.* -/var/log/user.log
```

Basic rule format is:

```text
facility.priority   action
```

`facility` is the program, `priority` is type of messages to log, `action` indicates where to store the log.

Log rotations is configurable in `/etc/logrotate.conf`.

### Shredding

We can shred a file to make it hard to recover:

```bash
# -f to force change permissions to be able to shred file
# -n N to specify how many times to shred.
shred -f -n 10 /var/log/auth.log.*
```

### Stop logging

```bash
service rsyslog stop
```

### Send Message to `syslog`

```bash
nc -w0 -u 0.0.0.0 5142 <<< "<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8"
```

See [SysLog Message Formats](https://blog.datalust.co/seq-input-syslog/#syslogmessageformats)
See [How to Send a messae to a syslog server](https://www.techiecorner.com/1496/how-to-send-message-to-syslog-server/)
