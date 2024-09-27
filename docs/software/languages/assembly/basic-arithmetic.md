---
title: "Basic Arithmetic"
---

# Basic Arithmetic

## `ADD`

## `SUB`

## `INC`

## `DEC`

## `MUL`

Multiplies `eax` by a value and stores the result 64-bit value across `edx:eax`:

```
mul value

mul 0x10
```

```
```

## `DIV`

Divides one number by another

```
DIV arg
```

Some forms:

- arg of size 8 bits:

```
al <- ax / arg; quotient
ah <- ax % arg; remainder
```

- arg of size 16 bits:

```
ax <- dx:ax / arg
dx <- dx:ax % arg
```

- arg of size 32 bits:

```
eax <- edx:eax / arg
edx <- edx:eax % arg
```

Examples:

- `div ch`: divides `ch` and stores the division result inside `al`. The remainder is stored inside `ah`.
- `div esi`: divides `edx:eax` by `esi`. Stores the quotient inside `eax`. Stores the remainder inside `edx`.
- `div di`: divides `dx:ax` by `di`. Stores the quotient inside `eax`. Stores the remainder inside `edx`.

`edx:eax` is the extension of `edx:eax` so:

```armasm
mv edx, '20' ; edx:eax is now 0000002000000000
mv ebx, '3A' 

div ebx      ; divide edx:eax by ebx ==> eax= 8D3DCB0, edx=30 (remainder)

div bx       
; divide ax:dx by bx ==> 
bx == 3A
dx:ax== 30CB08

ax == D75C (quotient)
dx == 0 (remainder)

```
