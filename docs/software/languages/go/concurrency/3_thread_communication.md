---
slug: thread-communication
title: Thread Communcation
author: [kbbgl]
tags: [go,concurrency,threads]
---

Threads working together require some form of communication.
There are **inter-thread communications** (ITC) and **inter-process communications** (IPC).
This type of communication falls under two main classes:

- [Memory sharing](#memory-sharing) is similar to talking to a person through a whiteboard.
- Message passing is similar to talking to a person through sending messages.

## Memory Sharing

All executions share the process' memory where each execution gets to write the results
of its own computation. We can coordinate the executions so that they collaborate
using the process' memory.

The memory sharing involves a few layers:

- main process memory.
- system bus.
- cache
- CPU + threads

In a scenario where a thread wants to read data from a memory location (the process' main memory),
the CPU requests data in the system bus. the system bus forwards the request
to the main memory. The entire memory block is read from the main memory and is placed on the bus.
The bus puts the memory in cache so it's closer to the CPU. The CPU uses this cache every time it needs to
access the same data.

In a scenario where thread 1 updates a variable and thread 2 wants to read this updated value, thread 1
updates the variable in cache memory. A **write-through** occurs where the cache line is flushed to main.
The cache line is placed on the bus and the main memory is updated from the bus.
The cache listens to the bus for updates and invalidates the cache line if it detects
the update. Then the CPU where thread 2 lives reads the updated variable.

### Sharing a Variable between goroutines

To share a variable between goroutines, we pass the variable's memory address (using `&variable`)
to the function that will run in the goroutine. In the function, we dereference and set the variable.
Then we can reference the variable which will have the updated value after the function run in the goroutine
finishes executing.

When we execute the `countdown()` function in a separate (from the main goroutine) goroutine,
the Go compiler realizes that we are sharing memory between goroutines and allocates
memory on the heap (instead of the stack). This is called **escaping**. Heap memory is used by all
parts of the process whereas stack memory is used only by one thread.

See [`countdown.go`](./code/countdown.go) for an example.

We can tell a variable has escaped to heap memory by asking the compiler to show its optimization
decisions:

```bash
go tool compile -m countdown.go
```

### Updating shared variables from multiple goroutines

To share variables between multiple goroutines, all we need to do is pass the variable that we want to manipulate into the function that is being run in the goroutine.

In [`letter_frequency_concurrent.go](./code/letter_frequency_concurrent.go), we can see that the`frequency` slice of integers is passed into the `countLetters` function where the slice is modified.

The program has a race condition though. If we run and compare the results between `letter_frequency_concurrent.go` and `letter_frequency_seq.go`, we can see that the letter 'e' has a different frequency:

```bash
 go run docs/software/languages/go/concurrency/code/letter_frequency_concurrent.go | grep -E "e-"
e-161063

 go run docs/software/languages/go/concurrency/code/letter_frequency_seq.go | grep -E "e-"
e-178215
```

Every time we run `letter_frequency_concurrent.go`, we will get a different result. This is because the program has a race condition. This means that multiple threads/processes share a resource but step over each other.

#### Race Conditions

In [`race_condition_demo.go](./code/race_condition_demo.go), we can see that we have a variable with an initial value (100) that we pass as a reference to 2 functions that either increase or decrease the value by 10. We run this increase and decrease the same amount of times in its own goroutine so we're expecting that after the functions finish, the amount would equal to the initial value. But this doesn't happen.

The cause is that the increase/decrease operation is not an **atomic operation**. An atomic operation is an operation that is indivisible and cannot be interrupted. The `*money -= 10` and `*money += 10` operations have 3 steps:

1. Read the value of the integer, `money`.
1. Modify the value.
1. Write the modified value back to memory.

We can use the `sync/atomic` and `atomic.AddInt32/64` that can safely increment or decrement shared variables. Also, if we set the program to use only one processor using `GOMAXPROCS=1`, all goroutines will run on the same processor and use the same cache so we see the expected behavior:

```bash
GOMAXPROCS=1 go run docs/software/languages/go/concurrency/code/race_condition_demo.go
Spendy Done
Stingy Done
Money in bank account:  100

GOMAXPROCS=1 go run docs/software/languages/go/concurrency/code/race_condition_demo.go
Spendy Done
Stingy Done
Money in bank account:  100
```

Using `GOMAXPROCS` is obviously not a good solution because this would take away the advantage of concurrency.

#### Solving Race Conditions

There's no single technique to solve all race conditions. We need to consider:

1. Is memory sharing really needed for the problem or is there another way to communicate between goroutines? We can use channels as well.
1. Is there a chance that a race condition can occur with the program?

To avoid race conditions, we need good synchronization and communcation with the rest of the goroutines.

#### Detecting Race Conditions

We can use the `-race` command line flag to detect race conditions. The compiler adds special code to all memory accesses to track when goroutines are reading and writing memory and outputs memory locations and source code line numbers where the program read/wrote to, causing a race condition.
