---
title: "Computer Memory"
---

# Computer's Memory

A computer's memory is usually modeled as a flat list of cells, i.e. flat memory model. Every cell has an address and the contents of every cell could be written to or read from.

There are different hardware implementations of computer memory:

- RAM, HD, USB, CD/DVD.
- The x86 processor has many instructions to communicate with the RAM.

## RAM

Random means that every memory cell could be accessed directly (in any random order)

The processor communicates with the RAM. The processor and the RAM are connected together through electricity in the motherboard.

The processor may send read/write requests to the RAM.

The electric lines which transfer the data are called **buses**.

## Motherboard

On the motherboard, the CPU and the RAM are connected through the Northbridge:

![](https://upload.wikimedia.org/wikipedia/commons/thumb/b/bd/Motherboard_diagram.svg/1280px-Motherboard_diagram.svg.png)

## Memory Abstraction

The programmer doesn't have to worry about memory management.

The program will run under the illusion of owning lots of flat memory. In reality, the program shares the total memory of the system with other programs.

The operating system and the processor work together using **Paging** and **Swapping** (using HD as memory) to create this illusion. The memory address your program sees are not real, they are **virtual**.
