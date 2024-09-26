# SELinux

Developed by the NSA.

It's a set of security rules that are used to determine which processes can access which files, directories, ports and other items in the system.

It works with 3 conceptual quantities:

- **Contexts**: labels to files, processes and ports.
- **Rules**: describe access control in terms of context, process, files, ports, users, etc.
- **Policies**: Set of rules that describe what system-wide access control decisions should be made by SELinux.

It can run in 3 modes:

- **Enforcing**: SELinux is operating, access is denied according to policy. All violations are audited and logged.
- **Permissive**: SELinux is enabled but only audits and warns about operations that would be denied in enforcing mode.
- **Disabled**: SELinux is disabled, no policies enforced.

Modes are written in `/etc/selinux/config` or `/etc/sysconfig/selinux`/`/etc/default/selinux`.

To check status of current mode and policy:

```bash
sestatus
```

To set or examine current mode:

```bash
getenforce

sudo setenforce Permissive
```

To disable SELinux:

```bash
# /etc/selinux/config
# set to disabled
SELINUX=disabled
```

## Policies

`/etc/sysconfig/selinux` or `/etc/default/selinux` sets the policy.

Multiple policies are allowed but only can be active.

Each policy has it's own files under:

```bash
/etc/selinux/[SELINUXTYPE]
```

Common policies:

- `targeted`: The default policy in which SELinux is more restricted to targeted processes. User processes and init processes are not targeted. SELinux enforces memory restrictions for all processes, which reduces the vulnerability to buffer overflow attacks.
- `minimum`: A modification of the targeted policy where only selected processes are protected.
- `MLS`: The Multi-Level Security policy is much more restrictive; all processes are placed in fine-grained security domains with particular policies.

## Contexts

- User
- Role
- Type (most common) - convention for name is `_t`, e.g `kernel_t`
- Level

To see context:

```bash
ls -Z
ps auZ
```

To change context:

```bash
chcon -t etc_t $somefile
chcon --reference somefile so
```

### Context Inheritence

Newly created files inherit the context from their parent directory, but when moving files, it is the context of the source directory which may be preserved, which can cause problems.

To reset file contexts based on parent:

```bash
restorecon
```

To configure the default context for new directories:

```bash
# make the change
semanage fcontext -a -t httpd_sys_content_t /virtualHosts

# apply it
restorecon -Rfv /virtualHosts
```

## Monitoring Access

SELinux comes with a set of tools that collect issues at run time, log these issues and propose solutions to prevent same issues from happening again. These utilities are provided by the setroubleshoot-server package. Here is an example of their use:

```bash
[root@rhel7 ~]# echo 'File created at /root' > rootfile
[root@rhel7 ~]# mv rootfile /var/www/html/
[root@rhel7 ~]# wget -O - localhost/rootfile
--2014-11-21 13:42:04-- http://localhost/rootfile
Resolving localhost (localhost)... ::1, 127.0.0.1
Connecting to localhost (localhost)|::1|:80... connected.
HTTP request sent, awaiting response... 403 Forbidden
2014-11-21 13:42:04 ERROR 403: Forbidden.

[root@rhel7 ~]# tail /var/log/messages
Nov 21 13:42:04 rhel7 setroubleshoot: Plugin Exception restorecon
Nov 21 13:42:04 rhel7 setroubleshoot: SELinux is preventing /usr/sbin/httpd from getattr access on the file .
....
Nov 21 13:42:04 rhel7 python: SELinux is preventing /usr/sbin/httpd from getattr access on the file .
....
Do allow this access for now by executing
# grep httpd /var/log/audit/audit.log | audit2allow -M mypol
# semodule -i mypol.pp

Additional Information:
Source Context system_u:system_r:httpd_t:s0
Target Context unconfined_u:object_r:admin_home_t:s0
Target Objects [ file ]
Source httpd
Source Path /usr/sbin/httpd
....
```
