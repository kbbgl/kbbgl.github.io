---
slug: python-threading-basics
title: Threading Basics
authors: [kbbgl]
tags: [python,threading,multithread]
---

## Threading Basics

CPUs have X amount of cores. This means that at any given time, the machine can perform **X amount of operations in parallel**.

Threads refer to the highest level of code executed by a processor, so with many threads, your CPU can handle several tasks at the same time. All CPUs have active threads, and every process performed on your computer has at least a single thread.

The number of threads you have depends on the number of cores in your CPU. **Each CPU core can have two threads**. So a processor with two cores will have four threads. A processor with eight cores will have 16 threads.

Running the following program will create 2 threads:

```python
import threading
# for python2 use 'import _thread'


import time

def func(y):
    print('ran')
    time.sleep(y)
    print("done")
    
x = threading.Thread(target=func, args=(4,))
x.start()

# Get number of active threads, 2 in this case
# 1 thread for running the actual script.
# 1 thread for running `x`.1
print(threading.activeCount())

# Thread syncronization
# Don't move to next command until x thread finishes execution
x.join()

print("Done")
```

## Thread Pool

See example [here](https://docs.python.org/3/library/concurrent.futures.html#threadpoolexecutor-example)

```python
from concurrent.futures import ThreadPoolExecutor

with ThreadPoolExecutor as executor:
    results = [executor.submit(some_function, some_function_arg) for _ in range(NUMBER_OF_THREADS)]

    for f in concurrent.futures.as_completed(results):
        print(f.result())
```
