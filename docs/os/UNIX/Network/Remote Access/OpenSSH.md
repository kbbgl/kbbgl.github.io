# Architecture Layers

## Transport Layer

Deals with the initial key-exchange and setting up a symmetric-key session.

## User Authentication Layer

Deals with authenticating and authorizing the user accounts.

## Connection Layer

Deals with the communication once the session is set up.

### SSH Session Overview

An SSH session starts with the Transport Layer, which sets up the Connection Layer. SSH communications then happen over the Connection Layer.

### OpenSSH Client

The OpenSSH host-wide client configuration is `/etc/ssh/ssh_config`. The per-user client configuration is`$HOME/.ssh/config`. SSH uses a key-based authentication.

The syntax for the client configuration can be found using the command man 5 ssh_config.

Other protocols can be tunneled over SSH. The X11 protocol support is part of the OpenSSH client. The VNC protocol support is a part of many different VNC clients.

You can also manually open a connection for any other protocol using the `LocalForward` and `RemoteForward` tokens in the OpenSSH client configuration.

### OpenSSH Server

The OpenSSH server is configured by editing the `/etc/ssh/sshd_config` file.

Some of the commonly changed server configurations are:

-Disable root access, or allow only key-based root access, using the PermitRootLogin token.
PermitRootLogin no (No root access)
PermitRootLogin without-password (Key-only root access)

- Disable or enable X11 Forwarding using the `X11Forwarding` token.

```bash
X11Forwarding no
X11Forwarding yes
```

- Disable or enable authentication forwarding using the `AllowAgentForwarding token`.

```bash
AllowAgentForwarding yes
AllowAgentForwarding no
```

### Per-User OpenSSH Configuration

`$HOME/.ssh/config` can be set up with shortcuts to servers you frequently access. Advanced, user-specific options are also available. The following is an example of how you can use ssh web:

```bash
Host web

     HostName www.example.com
     User webusr
```

This is an example of a more advanced configuration:

```bash
$ cat ~/.ssh/config

Host web

     KeepAlive yes
     IdentityFile ~/.ssh/web_id_rsa
     HostName www.example.com
     Port 2222
     User webusr
     ForwardX11 no
```

You can find all of the possible options in the `ssh_config` man page:

```bash
man 5 ssh_config
```

### OpenSSH Key-Based Authentication

OpenSSH client key-based authentication provides a passwordless authentication for users. Private keys can be encrypted and password protected.

The `ssh-agent` program can cache decrypted private keys.

The `ssh-copy-id` program can copy your public key to a remote host.

To generate a user key for SSH authentication, use:

```bash
ssh-keygen -f $HOME/.ssh/id_rsa -N 'supersecret' -t rsa
```

To start `ssh-agent` and use it to cache your private key, use:

```bash
eval $(ssh-agent)
ssh-add $HOME/.ssh/id_rsa
```

To copy your public key to the remote system overthere for remote user joe, use:

```bash
ssh-copy-id joe@overthere
```

Consult the man pages `ssh-keygen`, `ssh-add` and `ssh-copy-id` for details.

### Tunnel

#### Local Tunnel (-L)

The local tunnel indicates which port is to be opened on the local host (4242) and the final destination to be, charlie:2200. The connection to the final destination is going to be made by the machine bob.

#### Remote Tunnel (-R)

The remote tunnel requests machine bob to open a listening port 2424 to which any connection will be transferred to the destination, charlie:2200.

#### Dynamic Port Forwarding (-B)

The third type of tunneling, dynamic port forwarding, can be found in the ssh man page.

Option `-N` sets the option to not execute a command on connection to the remote system, and option `-f` informs `ssh` to go into background just before command execution.

### Parallel SSH Commands

Often, it is required to execute the same command on many systems to help facilitate this. The pssh package is available for most distributions. The pssh package typically includes:

`pssh`: parallel ssh
`pnuke`: parallel process kill
`prsync`: parallel copy program using rsync
`pscp`: parallel copy using scp
`pslurp`: parallel copy from hosts.
The program names may be slightly different on the different distributions.

The `pssh` command and friends use the existing `ssh` configuration. It is best to configure aliases, keys, known hosts and authorized keys prior to attempting to use pssh.

If there is a password or fingerprint prompt, the pssh command will fail.

When using `pssh`, it is convenient to create a file with a list of the hosts you wish to access. The list can contain IP addresses or hostnames. An example is provided below:

```bash
$ cat ~/ips.txt

127.0.0.1
192.168.42.1

$ pssh -i -h ~/ips.txt date

[1] 10:07:35 [SUCCESS] 120.0.0.1
Thu Sep 28 10:07:35 CDT 2017
[2] 10:07:35 [SUCCESS] 192.168.42.1
Thu Sep 28 10:07:35 CDT 2017
```

### Virtual Network Computing (VNC) Server

The Virtual Network Computing (VNC) server allows for cross-platform, graphical remote access. There are several implementations available; the current common version is tigervnc client and server. The server component has `Xvnc` (the main server for VNC and X), `vncserver` (Perl script to control Xvnc), `vncpasswd` (set and change vnc-only password), and vncconfig (configure and control a running `Xvnc`).

To assist in setting up a VNC server session, the `vncserver` script is recommended.

Startup script:

```bash
HOME/.vnc/xstartup
```

Kill option:

```bash
$ vncserver -kill <DISPLAYNUM>

$ vncserver

You will require a password to access your desktops.
Password:
Verify:
Would you like to enter a view-only password (y/n)? n
xauth: file /root/.Xauthority does not exist
xauth: (argv):1: bad display name "host:1" in "add" command
xauth: file /root/.Xauthority does not exist
New 'X' desktop is host:1
Creating default startup script /root/.vnc/xstartup
Starting applications specified in /root/.vnc/xstartup
Log file is /root/.vnc/host:1.log
Kill option is vncserver -kill <DISPLAYNUM>
```

When the server starts, it uses the `xstartup` configuration from the users `//.vnc` directory. If the `xstartup` file is altered, the `vncserver` needs to be restarted.

### VNC Client

VNC is a display-based protocol, which makes it cross-platform. It also means that it is a relatively heavy protocol, as pixel updates have to be sent over-the-wire.

The `vncserver` opens ports starting from 5901 and up. The display number :1 may also be used.

The client, `vncviewer`, is usually packaged separately. It connects to the VNC server on the specified port or display number. Passwords are not sent in clear text.

On its own, VNC is not secure after the authentication step. However, the protocol can be tunneled through SSH or VPN connections.
