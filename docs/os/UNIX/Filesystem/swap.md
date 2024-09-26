# Swap

Linux employs a virtual memory system that allows it to function as if it had more memory than it physically has. Overcommiting functions two ways:

- Programs (such as ones that have child processes that they inherit from parent using COW) don't always use all allocated memory.
- Less active memory regions are moved to disk when memory becomes low on the system.

Swapping is usually done on dedicated partitions or files. **The recommended swap size is the total RAM on the system**. The kernel  memory is never swapped out.

We can see how much swap area the system is using:

```bash
cat /proc/swaps

free -m
```

Other commands:

```bash
# format swap partition/file
mkswap

# activate swap partition/file
swapon

# deactivate swap partition/file
swapoff
```

Full procedure to create add swap:

```bash
# Create fs
dd if=/dev/zero of=swpfile bs=1M count=1024

# change ownership of `swpfile# to root
sudo chown root:root swpfile

# activate swap
sudo swapon swpfile

# check it exists
cat /proc/swapfile
Filename    Type  Size Used Priority
/home/app/lfs201/fs/swpfile         file  1048572 512 -2

# deactivate and remove swap
sudo swapoff swpfile
sudo rm swpfile
```
