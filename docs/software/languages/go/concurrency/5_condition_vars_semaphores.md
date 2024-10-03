---
slug: condition-vars-semaphores
title: Condition Variables and Semaphores
author: [kbbgl]
tags: [go,concurrency,threads,mutex,semaphore]
---

**Condition variables** give us controls to complement exclusive locking provided by mutexes. They allow us to unblock once a certain condition is met.

**Semaphores** allow us to control how many concurrent goroutines can execute a certain section simultaneously. They also allow us store an occurring event signal that can be accessed by an execution later on.

## Conditional Variables

Conditional variables work together with mutexes and give us the ability to suspend the current execution until we have a signal that a particular condition has changed.

The steps that usually involve using conditional variables are:

1. goroutines A holds a mutex and checks for a particular condition on some shared state.
2. If the condition is not met, goroutine A calls the `Wait()` function on the condition variable.
3. The `Wait()` function performs two operations **atomically**. This means that no other execution can come in between these two operations. It first releases the mutex and then blocks the current execution putting the goroutine to sleep.
4. goroutine B acquires the available mutex and updates the shared state. By updating the shared state, it calls either `Signal()` or `Broadcast()` on the condition variable and unlocks the mutex. We must make sure that there's another goroutine waiting for the `Signal()` or `Broadcast()` otherwise the signal/broadcast will be missed.
5. goroutine A receives the `Signal()` or `Broadcast()` and wakes up to reacquire the mutex. It can then recheck if the condition in the shared state was met and continue execution.

The `sync.Cond` type is the Go implementation of a conditional variable:

```go
type Cond
    func NewCond(l Locker) *Cond
    func (c *Cond) Broadcast()
    func (c *Cond) Signal()
    func (c *Cond) Wait()
```

To create a new condition variable using the `NewCond` function, we can use a `Mutex`:

```go

func doWork(int *state, condition *sync.Cod) {
    condition.L.Lock()
    *state++
    condition.Signal()
    condition.L.Unlock()
}


mutex := sync.Mutex{}
condition := sync.NewCond(&mutex)
sharedState := 100

go doWork(&sharedState, &condition)
```

[`race_condition_variable.go`](./code/race_condition_variable.go) has an example with comments on how we use a condition variable.

### Synchronizing multiple goroutines

In cases wen multiple goroutines are suspended on `Wait()`, `Signal()` will arbitrarily wake up one of the goroutines. `Broadcast()` will wake up all goroutines that are suspended on a `Wait()`.

## Semaphores

Semaphores allow us to specify the number of concurrent executions that are permitted. This means that we can control how many goroutines can access our shared resources. A semaphore is basically a mutex but with a variable amount of goroutines that are allowed to interact with the shared resource. However, in case of a mutex, the execution should hold and release it. Using semaphores, this is not always the case.

Semaphores provide 3 functions:

* New semaphore function which creates a new semaphore with variable amount of permits.
* Acquire permit function which allows a goroutine to take one permit from the semaphore. If none are available, the goroutine will suspend and wait until one of them is available.
* Release permit function which releases one permit so a goroutine can use it again with the acquire function.

### Building a semaphore

Go comes with an [extension to the `sync` package](golang.org/x/sync/semaphore).

A sample implementation of a semaphore:

```go
package main

import (
    "sync"
)

type Semaphore struct {
    // Permits remaining on the semaphore
    permits int

    // Condition variable used for waiting when
    // there are not enough permits
    cond *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
    return &Semaphore{
        // Initial number of permits 
        permits: n,

        // Condition variable used for waiting when
        // there are not enough permits
        cond: *sync.NewCond(&sync.Mutex{})
    }
}

func (rw *Sempahore) Acquire() {
    // Acquires mutex to protect permits variable
    rw.cond.L.Lock()
    for rw.permits <= 0 {
        // Waits until there is an available permit
        rw.cond.Wait()
    }
    // Decrease the number of available permits
    rw.permits--
    // Release mutex
    rw.cond.L.Unlock()

}

func (rw *Semaphore) Release() {
    // Acquires mutex to protect permits variable
    rw.cond.L.Lock()

    // Increases the number of available permits
    rw.permits++

    // Signals condition variable that one more permit 
    // is available
    rw.cond.Signal()

    // Releases mutex
    rw.cond.L.Unlock()
}
```
