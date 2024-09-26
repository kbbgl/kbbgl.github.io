---
title: "Defining Memory"
---

# Defining Memory

We can define memory in its own file which will assemble into a `.bin` file.

The syntax to define memory is as follows:

```
[label]    definition type     value    
```

Here are some examples:

```
; define one byte using Data Byte definition type of value 'ab' in hex
a    db   0abh

; define two bytes using Data Word definition with value 'cc99' in hex
b    dw   0cc99h

; Define a Double Word (4 bytes)
c    dd   12345678h

; Define a Quad Word (8 bytes)
d    dq   0aabbccdd11223344h

```

To define arrays, we use the `X dup [i, ..]` instruction where `X` is the number of duplications to generate and `i` is the initialized value.

For example:

```
array_bytes     db 4 dup (12h)

array_words     dw 5 dup (5678h)

array_dwords    dw 6 dup (aabbccddh)

array_qwords    dq 7 dup (1111222233334444h)

; will create 20 byte array.
repeat_byte     db 4 dup (1,2,3,4,5)

; repeat sequence of words (0006, 0007) 3 times
repeat_word     db 3 dup (6,7)

; will repeat the dword 6 times.
repeat_dword2   dd 3*2 dup (0abcd1234h)
```
