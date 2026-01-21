---
slug: profiling-golang-using-pprof
title: How to Profile Golang application using pprof
description: to Profile Golang using pprof
authors: [kgal-akl]
tags: [testing, development, go, golang, profiling, performance, memory]
---

Once you have the Golang application running and exposing the profiler endpoint we can take a snapshot:

```bash
PROFILING_PORT=3888
curl -sS -o "/tmp/app_heap_t0.pb.gz" "http://127.0.0.1:${PROFILING_PORT}/debug/pprof/heap?gc=1"
```

Then we do the actions that cause a memory spike and we take another snapshot:

```bash
curl -sS -o "/tmp/app_heap_t1.pb.gz" "http://127.0.0.1:${PROFILING_PORT}/debug/pprof/heap?gc=1"
```

And we compare them:

```bash
go tool pprof --base /tmp/app_heap_t0.pb.gz /tmp/app_heap_t1.pb.gz
```

Useful commands inside the comparison:

```
(pprof) top
(pprof) top -cum
```

We can use to show Go objects instead of sizes:
```
(pprof) sample_index=inuse_objects  
```

If we'd like to also get a snapshot of the goroutines and memory allocations, we can use:

```bash
curl -sS -o "/tmp/pprof_app_goroutines" "http://127.0.0.1:${PROFILING_PORT}/debug/pprof/goroutine?debug=2"
curl -sS -o "/tmp/pprof_app_allocs" "http://127.0.0.1:${PROFILING_PORT}/debug/pprof/allocs?seconds=30"
```