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
