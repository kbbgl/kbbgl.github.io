---
slug: gdb-cheatsheet
title: GDB Cheatsheet
tags: [gdb, debug, reverse_engineer, cheatsheet]
last_update:
  date: 12/31/2022
  author: kbbgl
---


# GDB Cheatsheet

### Show source code

This option will show the source code when we compile using the `-g` flag.

```
(gdb) list [function_name]
(gdb) list [start_line,end_line]
```

```c
(gdb) list main
1       #include <stdio.h>
2       #include <string.h>
3
4       void return_input (void)
5       { 
6          char array[30]; 
7
8          gets (array); 
9          printf("%s\n", array); 
10      }
```

### Show Functions and their memory addresses

```
(gdb) info functions 
All defined functions:

Non-debugging symbols:
0x00000000004003a8  _init
0x00000000004003e0  __libc_start_main@plt
0x00000000004003f0  __gmon_start__@plt
0x0000000000400400  _start
0x0000000000400430  deregister_tm_clones
0x0000000000400460  register_tm_clones
0x00000000004004a0  __do_global_dtors_aux
0x00000000004004c0  frame_dummy
0x00000000004004f0  fce
0x00000000004004fb  main
0x0000000000400510  __libc_csu_init
0x0000000000400580  __libc_csu_fini
0x0000000000400584  _fini
```

### Show Memory Address of Function

```
info address readflag
Symbol "readflag" is a function at address 0x555555555289.
```

### Dissassemble first line of function `win`

```
(gdb) disas /s win,+1
Dump of assembler code from 0x401de9 to 0x401dea:
pwd.c:
13      void win() {
   0x0000000000401de9 <win+0>:  endbr64
```

This option also allows us to see the Assembly instructions per line of source code:

```
(gdb) disas /s copier +1
Dump of assembler code for function copier:
copier.c:
16      int copier(char *str) {
   0x08049dc0 <+0>:     endbr32 
   0x08049dc4 <+4>:     push   %ebp
   0x08049dc5 <+5>:     mov    %esp,%ebp
   0x08049dc7 <+7>:     push   %ebx
   0x08049dc8 <+8>:     sub    $0x3e8,%esp
   0x08049dce <+14>:    call   0x8049df3 <__x86.get_pc_thunk.ax>
   0x08049dd3 <+19>:    add    $0x9b22d,%eax

17              char buffer[1000];
18              register int i asm("esp");
19              strcpy(buffer, str);
   0x08049dd8 <+24>:    pushl  0x8(%ebp)
   0x08049ddb <+27>:    lea    -0x3ec(%ebp),%edx
   0x08049de1 <+33>:    push   %edx
   0x08049de2 <+34>:    mov    %eax,%ebx
   0x08049de4 <+36>:    call   0x8049030
   0x08049de9 <+41>:    add    $0x8,%esp

20              return i;
   0x08049dec <+44>:    mov    %esp,%eax

21      }
   0x08049dee <+46>:    mov    -0x4(%ebp),%ebx
   0x08049df1 <+49>:    leave  
   0x08049df2 <+50>:    ret
```

We can also disassemble by range of memory addresses:

```
(gdb) info line main
Line 3 of "main.c" starts at address 0x401050 <main> and ends at 0x401075 <main+
(gdb) disas 0x401050 0x401075
Dump of assembler code from 0x401050 to 0x401075:
0x00401050 <main+0>:    push   %ebp
0x00401051 <main+1>:    mov    %esp,%ebp
0x00401053 <main+3>:    sub    $0x18,%esp
0x00401056 <main+6>:    and    $0xfffffff0,%esp
0x00401059 <main+9>:    mov    $0x0,%eax
0x0040105e <main+14>:   add    $0xf,%eax
0x00401061 <main+17>:   add    $0xf,%eax
0x00401064 <main+20>:   shr    $0x4,%eax
0x00401067 <main+23>:   shl    $0x4,%eax
0x0040106a <main+26>:   mov    %eax,-0xc(%ebp)
0x0040106d <main+29>:   mov    -0xc(%ebp),%eax
0x00401070 <main+32>:   call   0x4010c4 <_alloca>
```

### Create Breakpoint

```
(gdb) b 10

(gdb) disas return_input
...
   0x080491d0 <+26>:    call   0x8049070 <gets@plt>
...

# This will put a breakpoint at 99 bytes of `main` function
(gdb) break *(main+99)
(gdb) run
(gdb) jump *(main+104)
```

### Show and Delete Breakpoints

```
(gdb) info break
Num     Type           Disp Enb Address    What
1       breakpoint     keep y   0x080491e9 in return_input 
                                           at wh/stack_overflows/overflow.c:10
        breakpoint already hit 1 time
2       breakpoint     keep y   0x080491d0 in return_input 
                                           at wh/stack_overflows/overflow.c:8
        breakpoint already hit 1 time
3       breakpoint     keep y   0x080491e9 in return_input 
                                           at wh/stack_overflows/overflow.c:10
        breakpoint already hit 1 time

(gdb) del 1

```

### List sections and memory addresses

```
(gdb) info files

.bss
.data
.text
```

### Show Local Variables and Values

```
(gdb) list
31
32      enum HttpMethod {
33          GET,
34          POST,
35          PUT,
36          PATCH,
37          DELETE,
38          OPTIONS 
39      }
40

(gdb) info locals
put = server::Method::PUT
post = server::Method::POST

(gdb) p/d put
$1 = 2

(gdb) p/d post
$2 = 1

```

### Run and continue program execution

```
(gdb) r

# Run with arguments and input from file
(gdb) run --args < file_input

(gdb) c
```

### Execute one line of source code

```
(gdb) step
```

### Execute one line of Assembly code

```
(gdb) stepi
```

### Examine the 50 dwords in memory storing `array`

```
(gdb) x/50x $array
```

### Show Address Sections

We can use this command to output a table of address ranges

```
(gdb) info proc mappings
process 62432
Mapped address spaces:

        Start Addr   End Addr       Size     Offset objfile
         0x8048000  0x8049000     0x1000        0x0 /tmp/pwd
         0x8049000  0x80b4000    0x6b000     0x1000 /tmp/pwd
         0x80b4000  0x80e3000    0x2f000    0x6c000 /tmp/pwd
         0x80e3000  0x80e5000     0x2000    0x9a000 /tmp/pwd
         0x80e5000  0x80e7000     0x2000    0x9c000 /tmp/pwd
         0x80e7000  0x8109000    0x22000        0x0 [heap]
        0xf7ff8000 0xf7ffc000     0x4000        0x0 [vvar]
        0xf7ffc000 0xf7ffe000     0x2000        0x0 [vdso]
        0xfffdd000 0xffffe000    0x21000        0x0 [stack]
```

We can see the stack in the end.

### Examine 10 instructions 10 bytes before EIP

```
(gdb) x/10i $eip-10
```

### Examine registers

```
(gdb) info registers

eax            0xffffcbce          -13362
ecx            0x875a803d          -2024112067
edx            0xffffcc24          -13276
...
```

### Logging to File

```bash
#enable logging:

(gdb) set logging on

# Now GDB will log to ./gdb.txt. You can tell it which file to use:

(gdb) set logging file my_god_object.log

# And you can examine the current logging configuration:

(gdb) show logging
```

### Showing Assembly Layout

```
(gdb) layout asm
(gdb) break *(main+99)
(gdb) run
(gdb) jump *(main+104)
```
