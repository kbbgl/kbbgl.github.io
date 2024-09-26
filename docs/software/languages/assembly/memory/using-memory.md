---
title: "Using Memory"
---

# Using Memory

To assign memory:

```
# store the value of esi in eax inside dword
mov dword [esi],eax
```

## Constants and Data Initialization

The syntax for storage allocation statement for initialized data is:

```
[variable-name]    define-directive    initial-value   [,initial-value]...
```

We can use constants in the following way:

```
entry start

...
SOME_CONSTANT = 6

...

; dd  == define doubleword (4 bytes)
; dup == duplicate
; (0) == initialize all 6 values to 0
section '.data' data readable writeable
    some_nums     dd SOME_CONSTANT dup (0)


```

`sum_nums` is a label stored in RAM where the `data` section starts. `dd` defines 6, duplicated double `dword`s which are all initialized to 0. So if the `data` section starts at memory address `400000`, the double `dwords` will spread over 24 bytes starting from `400000` to `400024`.

## Initialized vs Uninitialized

We use the `.data` section for initialized memory. For example, we're initializing 4001 dwords to 0:

```

; 4 because every dword is 4 bytes
; 100000h/ 4 + 1 = 4001
AMOUNT_NUMS = (10000h / 4) + 1


section `.data` data readable writable
 some_nums      dd AMOUNT_NUMS dup (0)
```

This would cause the program size to be large and contain many 0s. We can use the `.bss` section to save space:

```
AMOUNT_NUMS = (10000h / 4) + 1


section `.bss` readable writable
 some_nums      dd AMOUNT_NUMS dup (?) ;? don't want to initialize

```
