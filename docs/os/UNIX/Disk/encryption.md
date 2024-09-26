# Disk Encryption

Linux distributions provide block device level encryption primarily through the use of LUKS (Linux Unified Key Setup). LUKS is installed on top of `cryptsetup`.

To encrypt a partition:

```bash
# This command uses the default encrpytion method. Will prompt for a secret.
sudo cryptsetup luksFormat /dev/sdc12

# Tto see To see other encryption types supported on the OS, run `cat /proc/crypto | grep name` and supply it in the arg, e.g. `aes`
sudo cryptsetup luksFormat --cipher aes /dev/sdc12

# To make volume available
sudo cryptsetup --verbose luksOpen /dev/sdc12 SECRET

# Formatting the partition
sudo mkfs.ext4 /dev/mapper/SECRET

# Mount it
sudo mount /dev/mapper/SECRET /mnt

# When done
sudo umount /mnt

# And remove association
sudo cryptsetup --verbose luksClose SECRET
```

We can also mount an encrypted partition at boot by:

1) Create entry in `/etc/fstab`

    ```bash
    /dev/mapper/SECRET /mnt ext4 defaults 0 0
    ```

2) Create entry in `/ect/crypttab`:

    ```bash
    SECRET /dev/sdc12
    ```
