---
title: "The Assembler"
---

The processor understand the instructions as numbers.

Every opcode was given a mnemonic to make it easier to remember.

For example:

| Encoding         | Instruction  |
|------------------|--------------|
| `B8 03 00 00 00` | `mov eax, 3` |
| `F7 F1`          | `div ecx`    |

The assembler translates the text based instructions into numeric encoding.

## Common x86 Assemblers

- MASM (Microsoft Macro Assembler)
- NASM (Netwide Assembler): cross platform.
- GAS (GNU Assembler): cross platform.
- FASM (Flat Assembler): cross platform written in assembly.

## Getting FASM

http://flatassembler.net
