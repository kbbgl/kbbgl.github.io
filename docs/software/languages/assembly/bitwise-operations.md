---
title: "Bitwise Operations"
---

# Bitwise Operations

The values inside the registers can also be seen as true (`1`) or false (`0`).

We can use this to perform meaningful operations between values of `true` and `false`.

### `NOT`

A unary operation that flips the bit.

```
mov ch,1001b
not ch
; 0110
```

### `AND`

A binary operation (produces one bit from two bits). Results in `1` only if both bits are `1`:

```
mov  al,11110000b
mov  dh,11001100b
and  al,dh
; al == 11000000
```

### `OR`

A binary operation. Results in `1` if either bit is `1`:

```
mov  al,11110000b
mov  dh,11001100b
or   al,dh
; al == 11111100b
```

### `XOR`

A binary operation. Results in `1` if either bit is `1` except when both are `1`:

```
mov  al,11110000b
mov  dh,11001100b
xor  al,dh
; al == 00111100b
```

We can use it to zero a register:

```
xor eax,eax
; eax == 0
```

## Bit Shifting

We can perform operations between bits in different positions by using bit shifting. For example, we can `AND` bit in position 3 with bit in position 5.

### `SHL` and `SHR`

Are used to shift bits left and right. Shifting to the left will insert 0s on the right end, shifting to the right will insert 0s on the left end.

```
mov  al,01001011b
shl  al,1
; al == 10010110b
; CF == 0
```

The argument passed to `shl/r` can only be a small number (1 byte, e.g `0x0` to `0xff`) or a register (e.g. `cl`)

## Arithmetic Shifting

When dealing with unsigned numbers, left shift is multiplication by 2 and right shift is a division by 2. But `SHR` division doesn't work correctly so use `SAR` instead.

## Rotating

Rotating means that the bits in a register are pushed left (`ROL`) or right (`ROR`).

So if we have the following instruction:

```
mov al,10011001
ror al,1

; last bit on the right is pushed into most significant bit
; al == 11001100
; last bit on the right is pushed into CF
; CF == 1
```

## Common Operations

If we wanted to extract a bit number `k` from a number `x`, we `AND` the value with a special `mask` that will cause all bits to equal `0` while the bit at position `k` will be equal to `1`.
For example, if we have the following number:

```
0110 1110 1101 1110 0111 1100 1110 0001
```

And we want the 7th bit (where first bit from the right is `k=0`, we will `AND` the number with:

```
0000 0000 0000 0000 0000 0000 1000 0000
```

```assembly

```
