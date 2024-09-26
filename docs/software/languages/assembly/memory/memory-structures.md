---
title: "Memory Structures"
---

# Memory Structures

We need ways to organize data that makes to us into the machine code. To do this, we use memory structures.

For example, we could define some point such as a coordinate `(x, y)` using 2 consecutive `dwords`:

```
section '.bss' readable writable
        ; Declare point
  point:        dd ?
                dd ?
       
section '.txt' code readable executable
start:
     ; setting x = 3, y = 4
 mov       dword [point]    , 3
 mov       dword [point + 4], 4
```

But this very messy!

## `struct`

We could instead use the `struct` directive to define data structures:

```
struct PNT
 x dd ?
 y dd ?
ends

section '.bss' readable writable
 my_pnt PNT ?
 
section '.text' code readable executable
start:
 mov    dword [my_pnt.x],3
 mov    dword [my_pnt.y],4
```

The assembler converts the 2nd way into the 1st way.

- Structs can be defined anywhere in the source file but usually is found before sections.
- The definitions creates a set of labels which point to memory addresses:
  - `my_pnt = 0x402000` - base address in memory.
  - `my_pnt.x = 0x402000`
  - `my_pnt.y = 0x402004`
  - `sizeof.my_pnt = 8`
  - `PNT` = 0
  - `PNT.x` = 0
  - `PNT.y` = 4
- You can define default values for fields of the `struct`:

```
struct PNT
 x dd 3
 y dd 4
ends
```

So if we wanted to access the `y` field and the size of a structure:

```
struct PNT
 x dd ?
 y dd ?
ends

section '.data' data readable writeable
 ; Declare a point
 my_pnt    PNT   3,4
 end_pnt:
 
section '.text' code readable executable
start:
 mov eax, dword [my_pnt + PNT.y]
 ; OR
 mov eax, dword [my_pnt.y]
 
 ; getting the size
 mov eax, sizeof.PNT
 
 ; OR
 mov eax, end_pnt - my_pnt
```

## Nesting Structures

```
struct PNT
 x dd ?
 y dd ?
ends

struct CLINE
 color   dd  ?
 p_start PNT ?
 p_end   PNT ?
ends
```

`sizeof.CLINE` in this case is 4 + 8 + 8 = 20 bytes.

We use the `<x,y>` brackets to define the nested structure.

```
struct PNT
 x dd ?
 y dd ?
ends

struct CLINE
 color   dd  ?
 p_start PNT ?
 p_end   PNT ?
ends

section '.data' data readable writeable
 my_line  CLINE 0, <3,4>, <1,5>
 
 
section '.text' code readable executable
start:
 ; eax == 4
 mov    eax, dword [my_line.color] 
 
 ; eax == 3
 mov    eax, dword [my_line.p_start.x]
```

### Anonymous Structures

```
struct DLINE
 struct ;anonymous
  red     db ?
  green   db ?
  blue    db ?
          db ? ; placeholder
 ends
 p_start PNT ?
 p_end   PNT ?
ends
```

In this case, `sizeof.DLINE` is 20 bytes == 0x14.

- `DLINE.red = 0x0`
- `DLINE.green = 0x1`
- `DLINE.p_end = 0xC`

## Unions

Sometimes we need to think about the same chunk of data in more than one way.

For example, we might want to store an IPv4 address as a `dword` but also be able to access each byte separately.

```
struct IPv4
 union
  struct
   a db ?
   b db ?
   c db ?
   d db ?
  ends
  addr dd ?
 ends
ends
```

Unions are anonymous. Inside unions, the memory location offset doesn't increase. So `IPv4.addr` and `IPv4.a` both point to the start memory address.

```
section '.data' data readable writeable
 lhost     IPv4    <127,0,0,1>
 
section '.text' code readable executable
start:
 mov     eax, dword [lhost.addr]
 ; eax == 0x0100007f
 
 mov     eax, dword [lhost]
 ; eax == 0x0100007f
 
 mov     bl, byte [lhost.d]
 ; bl == 1
 
 mov     bl, byte [lhost + 3]
 ; bl == 1
```
