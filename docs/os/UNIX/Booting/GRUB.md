# Grand Unified Bootloader (GRUB)

Bootloader that handles the early phases of system startup.

GRUB can:

- boot different operating systems.
- boot different kernel and initial ramdisks.
- modify boot parameters.

## Configuration

At boot, one the following configuration files are read:

```bash
/boot/grub/grub.cfg

/boot/grub2.grub.cfg

/boot/efi/EFI/redhat/grub.cfg
```
