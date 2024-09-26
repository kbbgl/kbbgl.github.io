---
title: objdump
---

# `objdump`

Displays information from object files.

## Disassemble

```bash
objdump -d $FILE
```

### See File Headers and Start Address

```bash
❯ objdump -f exit

exit:     file format elf64-x86-64
architecture: i386:x86-64, flags 0x00000112:
EXEC_P, HAS_SYMS, D_PAGED
start address 0x0000000000401000
```

### See Section Headers

```bash
❯ objdump -h exit

exit:     file format elf64-x86-64

Sections:
Idx Name          Size      VMA               LMA               File off  Algn
  0 .text         0000000c  0000000000401000  0000000000401000  00001000  2**4
                  CONTENTS, ALLOC, LOAD, READONLY, CODE

```

We have information about:

- Size of section.
- virtual memory address
- logical memory address
- offset from the beginning of the file
- alignment
- flags specific to the section
