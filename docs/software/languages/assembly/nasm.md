---
title: NASM Assembler
---

# NASM Assembler

## Generate 32-bit

```bash
nasm -f elf -o execve.o execve.asm

# -s strips debugging symbols
ld -m elf_i386 -s  -o execve execve.o
```

### With debugging symbols

```bash
nasm -f elf -gdwarf -o execve.o execve.asm
ld -m elf_i386 -o execve execve.o
```
