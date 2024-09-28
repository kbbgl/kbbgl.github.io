---
slug: dealing-with-threads
title: Dealing with Threads
author: [kbbgl]
tags: [go,concurrency,threads]
---

The OS gives us various tools to help manage concurrency. Two of them, **processes** and **threads** are used in our code. A process represents a program that is running on the system. A thread is a construct that executes within the process context to give a lightweight and  efficient approach to concurrency. A process is started with one, main thread. Processes do not share memory and therefore minimize communication with other processes.

When we run a command such as `grep hello large_readme.md`, the following things happen:

1. The user submits input for execution.
2. The OS places the job in the job queue.
3. Once the job is ready to execute, it is moved to the ready/run queue.
4. When the CPU frees up, the OS picks up the job from the ready queue and starts executing it on the CPU. The processor is running the instructions contained in the job.
5. As soon as the job requests an instruction to read from a file, the OS removes the job from the CPU onto an I/O waiting queue. It waits here until the requested I/O op returns data. If there's another job in the ready queue, the OS picks it up and executes it on the CPU.
6. The device will complete the read bytes from text file  I/O operation.
7. Once the I/O operation is complete, the job moves back to the ready queue. It waits for the OS to pick it up so that it can continue execution.
8. When the CPU frees up, the OS picks up the job and continues execution instructions on the CPU (try to find match in file text).
9. The system might raise an interrupt while the job is in execution. An **interrupt** is a mechanism used to stop current execution and notify the system of a particular event. This is handled by the interrupt controller.
10. The OS pauses execution of current job and puts the job back on the ready queue. The job of the OS scheduling algorithm is to determined which job from the ready queue to pick up for execution.
11. Steps 4-9 will repeat until the job execution finishes.

Steps 9 and 10 are an example of **context switch** which occurs whenever the system interrupts a job and the OS steps in to schedule another one. The OS saves the current job state so it can resume where it left off and also load the next job to be executed. This state is referred to as the **process context block** and is a data structure used to store all the details about a job.

### Concurrency with Processes

`CreateProcess()` on Windows or `fork()` on UNIX systems are syscalls that create a copy of an execution. They create the child process, allocate the required resources and loads the program code. The OS also makes a complete copy of the memory space, the process' resource handlers (registers, stack, file handlers, program counter). The new process then takes over this new memory space and continues execution from that point.

Support for creating and forking processes in Go is limited to the `syscall` package and is OS-specific. We find the `CreateProcess` function for Windows and `ForkExec`/`StartProcess` for UNIX.

Processes also do not share memory space.

### Concurrency with Threads

Creating a new process is resource-heavy and provides good isolation. Threads are the answer to some of the problems that come with using processes for concurrency. Creating a thread is much faster (can reach x100) and it consumes much less resources than a process.

Threads share memory space and run within the context of the same process. This means that if a change is made in the main memory by one thread (e.g. changing a global variable's value), this change is visible to all other threads. This is the main advantage of using threads, multiple threads can use this shared memory to work on the same problem together. Threads do not share stack space so local variables are only visible to the thread that created them. They also have their own instruction pointer.

There's no isolation so threads can overstep each other. To prevent this, we use thread communication and synchronization.

In Windows, we can create a thread using `CreateThread()` syscall. In UNIX, we use the `clone()` syscall with the `CLONE_THREAD` option. Java models threads as objects, Python blocks multiple threads from executing in parallel using a global interpreter lock and Go has the concept of the goroutine.

goroutines give us a lightweight construct, consuming far fewer resources than the OS thread.

### Creating goroutines

Suppose we have the [`sequential.go`](./code/sequential.go).

To put perform `doWork()` concurrently we put it in a goroutine as in [`parallel.go`](./code/parallel.go).

### Implementing goroutines in user-space

Instead of implementing threads at the kernel leve, we can have threads running
completely in user space. This means that the memory space that is part of our
application. Using user-level threads is like having different threads of
execution running inside the main kernel-level thread.

A process containing user-level threads will appear to have
just one thread of execution. The process is responsible for
managing, scheduling and context switching its own user-level threads.

Go provides a hybrid system that gives us the great performance of user-level threads
without most of the downsides. It achieves this by using a set of kernel-level threads
each managing a queue of goroutines.

Go's runtime determines how many kernel-level threads
to use based on the number of logical processors. This is set by env var `GOMAXPROCS` which is by default the number of CPUs
the system has.

Go's runtime will assign a **local run queue** (LRQ) for each kernel-level thread.
Each LRQ will contain a subset of goroutines in the program. There's also a **global run queue** (GRQ)
for goroutines that Go hasn't assigned yet to a kernel-level thread.
Each kernel-level thread will execute the goroutines present in its LRQ.
The system of moving goroutines from one queue to another is known in Go as **work stealing**. Its there to overcome blocking I/O calls and balancing work across processors.

We can force a goroutine to lock itself to an OS thread by calling `runtime.LockOSThread()` and `runtime.UnlockOSThread()`.

We can call the `runtime.Gosched()` after `go sayHello()` as seen in [`scheduler.go`](./code/scheduler.go).

### Concurrency vs parallelism

Concurrency is an attribute of the program code while parallelism is a property of the executing program.
Concurrent is about planning how to do many tasks at the same time. Parallelism is about performing many tasks at the same time.
