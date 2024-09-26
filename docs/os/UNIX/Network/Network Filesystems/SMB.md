# Server Message Block (SMB)

The Server Message Block (SMB) protocol was originally designed at IBM and later incorporated as the de facto networking file/print sharing system for Microsoft Windows.

In 1996, the latest version of the SMB protocol was renamed to Common Internet File System (CIFS), as many new features were added.

The Samba Project started as a reverse-engineered implementation of the SMB protocol for Solaris servers. To learn more about this, visit [Samba's website](https://www.samba.org/).

Samba is built to be an SMB/CIFS server which will run on any UNIX-like system.

## Samba

### Features

Samba features include the following:

- Samba can create file or printer shares.
- Samba version 3.x can act as a WindowsNT domain controller.
- Samba version 4.x can act as an Active Directory domain controller.
- Samba version 4.x is available on most distributions.

### Configuration

The default location for the Samba configuration file is `/etc/samba/smb.conf`, which uses an INI-file-like syntax, with section headers enclosed in square brackets.

```ini
[global]
    workgroup = MYGROUP
    server string = Samba Server Version %v
    log file = /var/log/samba/log.%m
    max log size = 50
    cups options = raw
```

Each individual share then goes into its own section:

```ini
[mainshare]
    path = /srv/exports/
    read only = yes
    comment = Main exports share
```

### User Accounts

Due to the difference in system password hashing mechanisms, Samba cannot verify some user passwords without additional help.

The `smbpasswd` command allows you to manage your passwords in both the Samba password file and in directory services.

To create a new password entry for Samba, do:

```bash
smbpasswd -a geoff
```

To change the password for `geoff`, do:

```bash
smbpasswd geoff
```

In the event that the UNIX username does not match the Samba username, the `/etc/samba/smbusers` file allows for the translation.  

```bash
UNIXNAME = SMBNAME SMBNAME2
```

### `testparm`

`testparm` is the syntax checker for `smb.conf`.

Once you have created your `smb.conf` file, test the syntax with the testparm command:

```bash
testparm

Load smb configuration files from /etc/samba/smb.conf
rlimit_max: increasing rlimit_max (1024) to minimum Windows limit
(16384)
Processing section "[mainshare]"
Loaded services file OK.
Server role: ROLE_STANDALONE
Press enter to see a dump of your service definitions
```

When you press Enter, you will get a parsed copy of the configuration file.

### Clients

There are several ways to interface with a Samba server:

Query the shares on a server:

```bash
smbclient -L 172.16.104.131 -U student


Enter SAMBA\student's password:
Domain=[UBUNTU] OS=[Windows 6.1] Server=[Samba 4.5.8-Ubuntu]
      Sharename     Type          Comment
      -------       ----          ---------
      cifs-share    Disk          Example share for testing home-export-cifs
      IPC$          IPC           IPC Service (ubuntu server (Samba, Ubuntu))

Domain=[UBUNTU] OS=[Windows 6.1] Server=[Samba 4.5.8-Ubuntu]
      Server                 Comment
      ---------              -------
      Workgroup              Master
      -------                -------
      LFLAB                  UBUNTU
```

Using the FTP-like interface:

```bash
smbclient //SERVER/mainshare
smb: \> get /foo
```

Mount the SMB/CIFS share into your name space:

```bash
mount -t cifs -o username,password //SERVER/share/ /mnt/point/
```

To avoid putting usernames and passwords into a world-readable file (`/etc/fstab`), the CIFS mount command takes the credentials=filename option.

```bash
mount -t cifs -o credentials=filename //SERVER/share/ /mnt/point/
```

The credentials file has the following syntax:

```bash
username=value
password=value
domain=value
```
