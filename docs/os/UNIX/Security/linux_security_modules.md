# Linux Security Modules

The LSM framework enhances security to the Linux kernel. The basic idea is to hook system calls and insert code whenever an application requests a transition to kernel (system) mode in orider to accomplish work that requiries enhances abilities.

## LSM Choices

- SELinux
- AppArmor
- Smack
- Tomoyo

Only one choice can be used on a server.
