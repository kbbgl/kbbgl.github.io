---
slug: sync-with-mutexes
title: Synchronization with Mutexes
author: [kbbgl]
tags: [go,concurrency,threads,mutex]
---

**Mutex** is short for mutual exclusion.
We can protect our code with mutexes so that only one goroutine at a time accesses a shared resource. In this way, we eliminate race conditions.
Variations on mutexes, sometimes called locks, are used in every programming language that supports concurrency. An example variation is called **readers-writers** mutexes.

## How to Use a Mutex

A mutex is used by marking the beginning and end of the section of the code that we want to protect.
The first step in using a mutex is to identify where in the code the shared resource is accessed, updated and needs to be protected.

When a goroutine comes to the mutex-marked section of the code, it first locks this mutex via an instruction in the the program code. It then executes the code and when done, it releases the lock so that other goroutines can access it.

In case another goroutine attempts to lock a mutex that is already locked, the goroutine attempting to relock will get suspended until the mutex is released.

In Go, mutex functionality is provided in the `sync` package under the `Mutex` type and provides the `Lock()` and `Unlock()` operations.

An example of how to use these operations can be found in [race_condition_mutex.go](./code/race_condition_mutex.go) when there are 2 goroutines and in [letter_frequency_concurrent_fixed.go](./code/letter_frequency_concurrent_fixed.go) when numerous goroutines are executed.

There's also a `TryLock()` operation but it should not be used in the majority of cases.

## Readers-Writer Mutexes

Mutexes have a performance penalty since it blocks concurrency. In some cases, such as read-heavy operations on shared data, readers-writer mutexes are a better fit. Readers-writer mutexes only block concurrency when we need to update a shared resource, not when reading. This is true because race conditions only happen if we change the shared state without proper synchronization. If we don't modify the shared data, there is no risk of race conditions.

Go has its own implementation of a readers-writer lock in `sync.RWMutex` type. These are the functions that it offers:

```go
type RWMutex
    // Locks writer part of mutex
    func (rw *RWMutex) Lock()

    // Locks read part of mutex
    func (rw *RWMutex) RLock()

    // Returns read part locker of mutex
    func (rw *RWMutex) RLocker() Locker

    // Unlocks read part of mutex
    func (rw *RWMutex) RUnlock()

    // Tries to lock writer part of mutex
    func (rw *RWMutex) TryLock() bool

    // Tries to lock read part of mutex
    func (rw *RWMutex) TryRLock() bool

    // Unlock writer part of mutex
    func (rw *RWMutex) Unlock()
```

A custom readers-writer mutex is implemented in [`custommutex.go`](./code/custommutex.go).
