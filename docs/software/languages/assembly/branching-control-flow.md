---
title: "Branching/Control Flow"
---

# Branching and Control Flow

If we want to control the flow of the program (such as adding conditional statements), we would need to use the `eip` register.

## `eip` Register

The `eip` register is the Extended Instruction Pointer.

It is **32 bits** in size (64 bits in size in long-mode). In 32 bit protected mode, `eip` cannot be changed.

**It contains/points to the address to the current instruction**. If we want to jump to a different instruction, we need to change `eip`.

## `JMP` Instruction

The `JMP` instruction sets the value of `eip`.

```armasm
JMP ecx ; changes eip to contents of ecx. it will continue from the address ecx
JMP 777d1044h ; changes eip to 0x777d1044. The program will continue execution on that address
```

There are two types of jumps:

- **Absolute jump**: Jump to a specified location in memory.
- **Relative jump**: Jump to a location which is X bytes from this location.

The assembler picks the suitable jump type.w

## Labels

When a program is launched, we usually can't predict their loading location.

Labels allow us to refer to a location in our program without knowing the exact address of that location at runtime. It is the job of the assembler to translate labels into actual addresses.

Example of an infinite loop:

```armasm
    mov ecx, 0
my_label:        ; label
    inc ecx
    jmp my_label ; eip will jump back to address of instruction 'inc ecx'
```

## Flags Register

Flags values are used to make branches and control the flow in code.

- It's a 32-bit register inside the x86 processor.

- There is no direct access to this register.

- Every bit in this register is a flag which represents `True` or `False`.

- It contains bits with a value of the result of the latest calculation and other system-related bits.

![image.png](https://boostnote.io/api/teams/OdksQwLn2/files/2d05e0b92856731702c2f3b2fc293a42b94135cf8f1a8a5d957c4efe5a1d2b68-image.png)

Every instruction can have certain effects on some bits of the flags register. It is the 'mood' of the process.

The most important are `CF`, `ZF`, `SF` and `OF`.

### The Zero Flag (`ZF`)

The most basic and fundamental flag.

- Is set (to 1) whenever the last calculation had a result of 0.
- Is cleared (set to 0) when the last calculation had a non-zero result.

```armasm
mov   eax,3h ;no effect on zero flag
mov   ecx,3h
sub   eax,ecx ; 3-3 = 0, zf is set (1)
```

### The Sign Flag (`SF`)

The flag equals the most significant bit of the last result.

```armasm
mov edx,0
dec edx ; edx == 0xffffffff therefore sf == 1

mov eax,0
dec eax ; edx == 1 therefore sf == 0
```

### The Carry Flag (`CF`)

The carry flag understands unsigned addition and subtraction.

cf is set to 1 when the addition of two numbers causes a carry out of the most significant bits. For example:

```armasm
mov eax,ffffffffh
add eax,1 ; eax == 0, cf == 1
```

```armasm
mov eax,f0h
mov ecx,35h
sub cl,al

; cl == 0x45
; carry flag == 1
```

It's also set when the subtraction of two numbers requires a borrow into the most significant bits.

The carry flag is a good indicator when calculations are not correct.

### The Overflow Flag (`OF`)

The overflow flag "understands" signed addition and subtraction according to the two's complement representation.

In addition, `of` is set if the addition of two positive numbers has a negative result or when two negative numbers have a positive result.

In subtraction, `of = 1` if `x - y > 0` where `x > 0` and `y < 0` or `of = 1` if `x - y > 0` where `x < 0` and `y > 0`.

It's an indicator that a signed arithmetic has a wrong result.

The processor looks on the most significant bit of the two operands and the msb of the result. The msb of the result represents the sign of the number. If the result of the operation has a 'reasonable' sign, the flag is cleared. If not, it's set.
If `x > 0` and `y < 0`, `of` is never set because the result of the operation can be of any sign.

Example:

```armasm
mov al,7fh
mov cl,1h
add al,cl

; al == 0x80 (1 0 0 0 | 0 0 0 0) so msb is set, therefore of = 1 
```

### Overflow and Carry Comparison

Every subtraction or addition operation will set either flag, no matter if it's signed/unsigned or there's a carry or not.

| Code                                                                 | CF  | OF  |
|----------------------------------------------------------------------|-----|-----|
| `mov eax,0x0` <br/> `sub eax,1` <br/> `;eax == 0xffffffff`            | `1` | `0` |
| `mov dl,0x7f` <br/> `add dl,0x1` <br/> `;dl == 0x80`                  | `0` | `1` |
| `mov ax,0x5` <br/> `mov si,0x4` <br/> `add si,ax` <br/> `;si == 0x9`   | `0` | `0` |
| `mov cl,0x80` <br/> `mov dl,0x80` <br/> `add cl,dl` <br/> `;cl == 0x0` | `1` | `1` |

## Basic Conditional Branching

The `JMP` instruction unconditionally changes the value of `eip`.

To be able to jump when a different condition is met, we can use the `Jcc` instruction where `cc` represents some condition. The condition is usually dependent upon the value of one of the flags register.

These are the possible conditional jumps available:

- jump depending on carry flag is set or not, `JC` and `JNC`
- jump depending on overflow flag is set or not, `JO` and `JNO`
- `JS` and `JNS`

### `JZ` and `JNZ`

Jump Zero and Jump Not Zero will occur only when the `ZF` is set or not set, respectively.

For example:

```
  mov ax,1
  dec ax
  jz my_label 
  add ax,5
my_label:
  add ax,2
  
; ax == 2
```

```
  mov ax,1
  inc ax
  jnz my_label 
  add ax,5
my_label:
  add ax,2
  
; ax == 9
```

So we can construct a simple loop:

```
mov eax,0
mov ecx,3

again:
  add eax,ecx
  dec ecx
  jz outside
  jmp again
outside:
  ...
  
```

Or simpler using `JNZ`:

```
mov eax,0
mov ecx,3

again:
  add eax,ecx
  dec ecx
  jnz again
  ...
  
```

### `CMP`

This instruction is a simulation of the `SUB` instruction and used for comparison of numbers.

```
cmp eax,edx
```

Using the instruction doesn't modify the `eax` register (unlike `sub`) but it does modify the values of the flags register.

### Unsigned Comparison Instructions

We can use the following instructions to perform unsigned number comparison:

- `JB`: Jump if below, checks if `CF` is set. This instruction is just another name for the `JC` instruction.
- `JBE`: Jump if below or equal, checks if `CF` or `ZF` is set.
- `JA`: Jump if above, checks if `CF` and `ZF` is not set.
- `JBA`: Jump if above or equal, checks if `CF` is not set. This instruction is just another name for the `JNC` instruction.

```
  cmp   eax,ecx
  jb    my_label ; can use JB or JC since they are synonymous.
  ; we are here if eax >= ecx
  jmp   outside
my_label:
  ; we are here if eax < ecx
outside:
  ; ...
```

### Signed Comparison Instructions

We can use the following instructions to perform signed number comparison:

- `JL`: Jump if less, checks if `SF != OF`
- `JLE`: Jump if less or equal, checks if `SF != OF && ZF == 1`.
- `JG`: Jump if greater, checks if `SF == OF && ZF == 0`.
- `JGE`: Jump if greater or equal, checks if `SF != OF`

## Structured Branching

High-level languages use conditionals (`if`, `switch`) and loops (`for`, `while`). We can perform the same using Assembly code.

Branching/jumping should be used only for the following operations:

- Part of `break` statement.
- Part of `if` statement.
- Part of `while` statement.
- Part of `for` statement.

For `break` and `if` we jump forward.

For `while` and `for` we jump backwards.

### Conditionals

For example, if we have pseudo-code:

```
if eax < edx:
 eax++
else:
 eax--

end if
```

Can be translated to assembly:

```
 cmp eax,edx
 jae else
 inc eax
 jmp end_if
else:
 dec eax
end_if:
 ...
```

### Loops

If we have the following pseudo-code for a `for` loop:

```
for ecx from 0 to 99 do:
 eax = eax + ecx

end for
```

We can translate it into Assembly code:

```
 mov ecx,0
for_loop:
 add eax,ecx
 inc ecx
 cmp ecx,100d
 jnz for_loop
```

If we have the following pseudo-code for a `while` loop to sum 1+2+3... until we reach 1000:

```
eax = 0
ecx = 0
while ecx > 1000:
 eax = eax + ecx
 ecx++
end while
```

We can translate it into Assembly code:

```
 mov eax,0
 mov ecx,0
while_loop:
 cmp eax,1000d
 jae end_while
 
 add eax,ecx
 inc ecx
 jmp while_loop
end_while:
 ...
```

Sometimes we need to break a loop. In high-level languages we use the `break` statement.

Let's say we have a program that would sum numbers until it reaches 1000 but it can't sum more than 300 numbers.

In pseudo-code, it would look like this:

```python
eax = 0

for ecx from 0 to 299 do:
 eax = eax + ecx
 
 if eax >= 1000:
  break
```

In assembly:

```
 mov eax,0
 mov ecx,0
for_loop:
 add eax,ecx
 cmp eax,1000d
 jae end_for
 inc ecx
 cmp ecx,300d
 jb  for_loop
end_for:
 ...
```
