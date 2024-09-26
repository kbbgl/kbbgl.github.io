# NFS

Network File System is a filesystem protocol built upon the Open Network Computing Remote Procedure Call system (ONC RPC). RPCs are managed by the portmap service.

## Server Configuration

The main configuration is done in `/etc/exports` file.
The file has:

- A directory to share.
- Host with mount options.

The server can be reloaded using:

```bash
exportfs -ra
```

The syntax of the `exports` file is:

```bash
# <DIR> <HOST OR NETWORK> [<OPTIONS>]

/srv/nfs 192.168.122.0/24(rw,sync,root_squash)
```

## Client Configuration

The NFS client mounts the remote filesystem onto the local system. There a few important commands:

```bash
# queries mount daemon on the remote server for information including shares that are available for mounting
showmount -e server

# daemon is a dynamic port mapping daemon designed to reduce usage of well-known port
portmap


# mount command has the filesystem type NFS, which links to the mount.nfs command. There are two formats of the mount command for NFS shares.
# - One option for mounting NFS shares is mount HOST:/export /mount-point where the host:/export portion causes the mount command to process this mount as NFS.
# - The other form of the mount command is mount -t NFS HOST:/export /mountpoint which specifies which NFS is being used.
# mount SERVER:/share /mnt/share
```

## Security Considerations

The NFS default security is to use the UNIX `UID` and `GID`. The challenge of using `UID/GID` on different systems is the values must match. User `Bib` with `UID 1000` on system A must have the same `UID` on system B, or the wrong information may be accessed. Having a single sign on system with network available information like `NIS`, `LDAP`, or `Kerberos` will remove the `UID` confusion issues. NFS4 is Kerberos-aware and an excellent option to eliminate the `UID/GID` mapping issues.

The `root_squash` option `/etc/exports` file translates the root user's `UID/GID` (0) to an anonymous `UID/GID`. This is on by default and prevents root level access to the shared files. You should not disable `root_squash` without a good reason.

## Performance Considerations

Many factors contribute to the speed and performance of an NFS server or client.

Properly setting the values of `rsize` and `wsize` will allow for greater speed in a file transfer. However, you can only reasonably increase the block size to the `MTU` of your network between client and server. Increasing the frame size (Jumbo Frames) is one option.

Moving from a 1G to a 10G Ethernet network would vastly speed up an NFS setup.

The asynchronous mode trades speed for lack of robustness. An unclean shutdown of a server or client operating in an asynchronous mode has the potential to corrupt the data.
