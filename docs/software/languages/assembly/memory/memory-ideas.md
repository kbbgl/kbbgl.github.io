---
title: "Memory Ideas"
---

# Memory Ideas

## Array of `structs`

```
NUM_DOGS = 12

section '.bss' readable writable
 ; will allocate 96 bytes
 my_dogs   db    NUM_DOGS*sizeof.DOG dup (?)
 
section '.txt' code readable executable
start:
 ...
 ; Access dog in register ecx
 mov      esi,my_dogs
 lea      esi, [esi + ecx*sizeof.DOG]
 
 ; Access elements of struct
 mov      eax, dword [esi + DOG.color]
 mov      edx, dword [esi + DOG.age]

```

## Multiplication Table

We can think of a multiplication table as:

```
WIDTH = 4
HEIGHT = 4

section '.bss' readable writeable
 mul_tbl    dd WIDTH*HEIGHT dup (?)
 
section '.text' code readable writeable
 mov esi, mul_tbl ; cell ptr
 mov esx, 0       ; row counter
 
next_row:
 mov ebx, 0       ; column counter
next_column:
 mul eax, ecx
 mul ebx
 
```
