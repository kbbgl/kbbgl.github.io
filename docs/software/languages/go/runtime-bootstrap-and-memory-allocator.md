---
title: Go runtime bootstrap and memory allocator
slug: go-runtime-bootstrap-memory-allocator
authors: [kbbgl]
tags: [go, runtime, memory, allocator, gc, scheduler, performance, profiling]
---

This note condenses two excellent Internals for Interns articles into a mental model you can use while debugging performance, memory, and concurrency issues in Go.

## Related docs in this repo

- **Profiling (heap/allocs/goroutines)**: [`pprof.md`](./pprof.md)
- **Concurrency basics (goroutines/channels)**: [`channels-and-routines.md`](./channels-and-routines.md)
- **Debugging with Delve**: [`debugger.md`](./debugger.md)

## Runtime bootstrap: what runs before `main.main`

The entry point of a Go binary is *not* `main.main`. It starts in an architecture-specific runtime entry stub like `_rt0_amd64_linux`, which jumps into `rt0_go` to bootstrap the runtime. During this phase the runtime (roughly):

- **Creates `g0` and `m0`**: the initial goroutine (`g0`, used for runtime work) and initial OS thread (`m0`).  
- **Sets up TLS**: thread-local storage is used so the runtime can quickly find “the current goroutine” for a given OS thread.
- **Detects CPU capabilities**: checks CPU/vendor features (and may refuse to run if required instructions aren’t available).
- **Transitions into Go runtime init**: saves args, initializes OS/CPU info (including a default `GOMAXPROCS`), then calls `schedinit()` which wires up the major subsystems.

Inside `schedinit()` the runtime initializes key pieces such as:

- **Stack pools**: goroutines start small (e.g. 2KB) and grow; stacks are pooled to make goroutine creation cheap.
- **Allocator** (`mallocinit()`): sets up heap management; later, per-`P` caches make the common allocation path lock-free.
- **Map hashing selection** (`alginit()`): chooses faster hashing (e.g. AES-based) when hardware support is present.
- **Type/interface runtime tables**: module/type/linkage structures plus interface dispatch tables (`itabs`).
- **Environment and security checks**: parses env vars (including `GODEBUG=*`) and validates stdio fds are open.
- **GC is prepared but not enabled yet**: `gcinit()` sets up GC machinery, but `gcenable()` happens later once the runtime can safely start goroutines/channels.

Then the runtime creates the first “real” goroutine to run **`runtime.main`** (not your `main` yet). In `runtime.main`, it:

- **Starts the system monitor** (`sysmon`): watchdog thread that helps with scheduling fairness, syscalls, GC pacing nudges, etc.
- **Enables GC** (`gcenable()`), then runs package `init()`s in dependency order.
- **Calls your `main.main`**.

One practical implication: when **`main` returns**, the process is exiting; goroutines that are still running are not “gracefully completed” by default. Use explicit synchronization (e.g. `sync.WaitGroup`, context cancellation) if you need work to finish.

### Runtime debug knobs worth remembering

`GODEBUG` can emit high-signal traces during startup and runtime:

- **`GODEBUG=inittrace=1`**: timings for package init.
- **`GODEBUG=schedtrace=1000`**: scheduler state every second.
- **`GODEBUG=gctrace=1`**: GC cycle logs.

## Memory allocator: how heap memory is organized and why it’s fast

### Stack vs heap (and why the allocator matters)

Not all variables go through the allocator. The compiler uses **escape analysis** to decide what can stay on the stack vs what must live on the heap (e.g. if it outlives the current function call). The allocator manages heap memory and also provides the backing memory for goroutine stacks (even though stack variables themselves are not individually “allocated” via the heap allocator).

### Arenas → pages → spans → slots

The allocator avoids frequent OS syscalls by requesting memory in large chunks and managing it internally:

- **Arenas**: typically **64MB** chunks of *address space* reserved from the OS (with OS-level commitment happening incrementally as memory is used).
- **Pages**: arenas are subdivided into **8KB** runtime pages (distinct from typical 4KB OS pages).
- **Spans**: one or more contiguous pages dedicated to objects of a single size (and scan vs noscan category).
- **Slots**: a span is subdivided into fixed-size slots; allocation is often “find next free slot”.

Spans track allocation state with bitmaps such as:

- **`allocBits`**: which slots are currently allocated.
- **`gcmarkBits`**: which slots were marked live by the GC in the last mark phase.

After a GC mark phase, the runtime can effectively “free” dead objects by swapping/interpreting these bitmaps (live remains allocated; unmarked becomes available again).

### Size classes (8B → 32KB) and the edges

For **small objects (up to ~32KB)**, allocations are rounded up to one of **68 size classes**. This keeps allocation fast and limits fragmentation inside spans.

Two important edge behaviors:

- **Large objects (> 32KB)**: bypass size classes and get their own “class 0” span sized to the object.
- **Tiny allocator (< 16B, no pointers)**: packs multiple tiny, pointer-free objects into a single 16-byte block to reduce waste (so “1-byte things” don’t each consume an 8-byte slot).

### Span classes (scan vs noscan)

The allocator separates objects that **contain pointers** (must be scanned by GC) from those that **do not** (can be skipped). Combining:

- **size class** × **(scan|noscan)**  

…yields **136 span classes** in total.

### Concurrency: `mcache` → `mcentral` → `mheap`

To avoid allocator lock contention under heavy concurrency, Go uses a cache hierarchy:

- **`mcache` (per-`P`, fast path, lock-free)**: most allocations are served from a per-processor cache because only one goroutine runs on a `P` at a time.
- **`mcentral` (per span class, briefly locked)**: refills `mcache` with spans when local spans are exhausted.
- **`mheap` (global, expensive/rare path)**: supplies new pages/spans (and may request new arenas from the OS).

### Sweeping and scavenging: reclaiming memory

- **Sweeping**: spans may be swept lazily “on demand” when `mcentral` needs to reuse spans; sweeping interprets GC mark results to find free slots.
- **Scavenging**: even after objects are freed, the runtime often keeps pages mapped for reuse. A background scavenger can advise the OS to reclaim physical memory for unused pages (e.g. via `MADV_DONTNEED` on Linux) while keeping address space reserved.

## How to apply this when profiling/debugging

- If you’re looking at heap growth, allocations, or goroutine state, jump to [`pprof.md`](./pprof.md) for capture/compare commands (including `heap?gc=1`, goroutine dumps, and allocs sampling).
- If you suspect allocator/GC pressure:
  - consider whether objects are **tiny** (packed), **small** (size classes), or **large** (>32KB direct spans),
  - and whether your objects contain pointers (scan cost) vs not (noscan).
- If you see “work still running” at exit, remember that returning from `main` ends the process; coordinate goroutines explicitly.

## Sources

- Internals for Interns — *Understanding the Go Runtime: The Bootstrap*: `https://internals-for-interns.com/posts/understanding-go-runtime/`
- Internals for Interns — *Understanding the Go Runtime: The Memory Allocator*: `https://internals-for-interns.com/posts/go-memory-allocator/`
