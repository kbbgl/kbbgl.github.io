# Troubleshooting Boot Failures

## No boot loader screen

Check for GRUB misconfiguration, or a corrupt boot sector. You may have to re-install the boot loader.

## Kernel fails to load

If the kernel panics during the boot process, it is most likely a misconfigured or corrupt kernel, or incorrect parameters specified on the kernel command line in the GRUB configuration file. If the kernel has booted successfully in the past, it has either been corrupted, or the kernel command line in the GRUB configuration file has been altered in an unproductive way. Depending on which, you can reinstall the kernel, or enter into the interactive GRUB menu at boot and use very minimal command line parameters and try to fix that way. Or, you can try booting into a rescue image as described in the next chapter.

## Kernel loads, but fails to mount the root filesystem

The main causes here are:

- Misconfigured GRUB configuration file
- Misconfigured `/etc/fstab`.
- No support for the root filesystem type either built into the kernel or as a module in the initramfs initial ram disk or filesystem.

## Failure during the init process

There are many things that can go wrong once init starts; look closely at the messages that are displayed before things stop. If things were working before, with a different kernel, this is a big clue. Look out for corrupted filesystems, errors in startup scripts, etc. Try booting into a lower runlevel, such as 3 (no graphics) or 1 (single user mode).
