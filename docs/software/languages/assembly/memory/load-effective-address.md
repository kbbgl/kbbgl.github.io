---
title: "Load Effective Address (LEA)"
---

# Load Effective Address `LEA`

Syntax:

```
l dest, [expr]
```

- `lea` calculates `expr` and stores the result inside `dest`.
- **It doesn't actually access any memory**. It only calculates the resulting address.
- Can be used to calculate addresses.
- Doesn't change the flags register.
- `dest` has to be a register.

## Examples

```
; will add 1 to eax and store the result in eax
lea  eax, [eax+1]

lea esi, [eax + 2 * edx]

; di is 16 bits (2 bytes), we might need to
; wrap around if expression so (eax + 2 * edx + 5) is larger than 2 bytes 
; ==> (eax + 2 * edx + 5) % 2^16
lea di, [eax + 2 * edx + 5]
```

### Calculate Addresses

```
section '.data' data readable writeable
 nums      dd 100h dup (12345678h)
 snums     dw 100h dup (0ababh)
 
section '.text' code readable executable
start:
 mov       esi,nums
 mov       edi,snums
 call      read_eax
 
 ; Get address of dword number eax
 lea       edx, [esi + 4*eax]
 
 ; Get dword number eax
 mov       edx, [esi + 4*eax]
 
 ; Get address of word number eax
 lea       ebx, [edi + 2*eax]
 
 ; Get word number eax
 mov       ebx, [esi + 2*eax]
```

### Adding Numbers from 0 to 100 using `lea` and `test`

```
 mov       ecx, 100
 xor       ecx, ecx
 
add_num
 add       edx, edx
 lea       ecx, [ecx - 1]
 test      ecx, ecx
 jnz       add_num
```

### Combining Instructions

```
mov      esi, ecx
shl      edx, 2
add      esi, edx
add      esi, 5
```

Can be converted to:

```
lea      esi, [ecx + edx * 4 + 5)
```
