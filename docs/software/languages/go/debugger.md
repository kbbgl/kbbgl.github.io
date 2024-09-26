---
title: Using Delve Debugger
slug: how-to-debug-go-code-delve
authors: [kbgll]
tags: [go,debugging,debug,gdb]
---

#### Sources:

- https://chetan177.medium.com/runtime-debugging-in-golang-b8a065d0fb5e

## Building and Running

After making code changes in `/path/to/server`, build the binary:
```bash
go build -gcflags="all=-N -l" -v
```

The flags all


Then run server:
```bash
./server -stdout -loglevel debug
```

## Attach to Process

On another session, attach debugger:
```bash
$GOPATH/bin/dlv attach $(pidof server)
```

## Debugging Code

### Breakpoints

#### Set Breakpoint
```bash
break /path/to/module.go:120
```

#### List/Clear Breakpoints
```bash
breakpoints
clear 1
```

### Control Flow

#### See what code line we're in
```bash
(dlv) l
```

#### Continue until breakpoint/exit
```bash
(dlv) c
```
#### Go To Next Line
```
(dlv) n
```

#### Step Into Function
```bash
step
```

### Examine Variables

#### List local variables

```bash
locals
```

#### Print variable 
```bash 
print x
```

