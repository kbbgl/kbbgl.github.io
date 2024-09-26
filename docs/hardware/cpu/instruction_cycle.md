---
slug: cpu-instruction-cycle
title: CPU Instruction Cycle
tags: [hardware, cpu, resources]
last_update:
  date: 12/31/2022
  author: kgal-pan
---

From when the computer is boot-up until shutdown, it follows the cycle to process instructions.

## Components

The program counter (`PC`) is the register that holds the memory address of the next instruction to be executed.

The memory address register (`MAR`) holds the address of the instruction to be executed.

The memory data register (`MDR`) acts as a two-way register that holds data fetched from memory or data waiting to be stored in memory (can also be known as `MBR`).

The current instruction register (`CIR`) acts as a temporary storage for the instruction fetched from memory.

The control unit (`CU`) decodes the instruction in the CIR and sends signals to the arithmetic logic unit (`ALU`) and the floating point unit (`FPU`)

### Fetch

- The address of the `PC` is copied into the `MAR`.
- The `PC` is incremented to point to the next instruction.
- The instruction in address at `MAR` is copied to the `MDR`.
- The instruction in `MDR` is copied to the `CIR`.

### Decode

- The encoded instruction held in the `CIR` is decoded.

### Execute

- The `CU` of the CPU passes the decoded information as signals to the CPU (`FPU` or `ALU`) to perform the instructions and storing the result back into memory, register or an output device.
- Example operations: Add, Subtract, AND, OR, Branches of execution, XOR.

https://youtu.be/vgPFzblBh7w
https://youtu.be/o_WXTRS2qTY
