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
