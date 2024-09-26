---
title: "Signed Operations"
---

# Signed Operations

## `NEG` Instruction

`NEG` negates the sign of the number using two's compliment. It flips all the bits (0 becomes 1, 1 becomes 0) and add 1.

Example in 8 bits

```
mov al,3d
neg al

; 3 in binary is [ 0 0 0 0 | 0 0 1 1]
; neg 3 ==       [ 1 1 1 1 | 1 1 0 0] == 0xfd
```

## `MOVSX` and `MOVZX`

In some cases we want to extend an 8-bit to a 16-bit number.

When we want to extend a positive number, we add leading `0`s.
So `3`  == `[ 0 0 0 0 0 0 1 1]` becomes `[ 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1 ]`.

When we want to extend a negative number, we add leading `1`s.
So `-3` == `[ 1 1 1 1 1 1 0 1]` becomes `[ 1 1 1 1 1 1 1 1 1 1 1 1 1 1 0 1]`.

Because this is cumbersome, we have `MOVSX` and `MOVZX` which extend while moving.

To extend unsigned numbers, use `MOVZX`:

```
movzx eax,bl ; extends bl using leading zeroes and stores the result into eax.
```

To extend signed numbers, use `MOVSX`. If the leading bit is `1`, `MOVSX` extends using leading zeroes. If the leading bit is `0`, `MOVSX` extends using leading `1`s.

## `CBW` and `CWDE`

`CBW` - convert byte to word (word is 2-bytes). It is used to sign-extend `al` to `ax`.
`CWDE` - convert word to double word. It is used to sign-extend `ax` to `eax`.

For example:

```
mov al,10001010b
cbw
; ax == 1111111110001010

cwde
; eax == 11111111111111111111111110001010
```

There are also `CWD` and `CDQ` specifically for the `edx` register.

`CWD` - convert word to double-word. It is used to sign-extend `ax` to `dx:ax`.
`CDQ` - convert double-word to quadword (32-bit to 64-bit). It is used to sign-extend `eax` to `edx:eax`.

## `IMUL` and `IDIV`

`IMUL` and `IDIV` are sign-aware multiplication and division instructions.

`CDQ` and `IDIV` are usually used together:

```
; program to divide eax by 3

mov  esi,3
cdq
idiv esi 
```
