# I/O Scheduling

The I/O scheduler provides an interface with the **Generic Block Layer**, which is a kernel component that handles the requests for all block devices in the system.

A **block device** reads and writes 'blocks' of data to and from the hardware (usually a hard disk).

Both the **Virtual Memory (VM)** and **Virtual File System(VFS)** layers submit requests to block devices.

The scheduling scheme is loaded at kernel boot.

To view the available schedulers:

```bash
# Selected is cfq scheduler
cat /sys/block/sda/queue/scheduler
noop deadline [cfq]
```

To change the value:

```bash
echo deadline > /sys/block/sda/queue/scheduler
```

## Types of Schedulers

1) Completely Fair Queue (`cfq`)

    - spreads i/o bandwidth equally across all processes submitting requests.
    - there are 64 dispatch queues which receive requests.
    - requests are read from all queues in round-robin and in FIFO per queue.

2) `deadline`

    - aggressively reorders requests to prevent large latencies and improve overall performance.
    - the kernel associates a deadline for every request.
    - has 5 queues:
    - 2 sorted lists (1 for reading, 1 for writing) and are arranged by the starting block.
    - 2 FIFO lists (1 for `r`, 1 for `w`) and are sorted by submission time.
    - dispatch queue which contains requests to be sent into the device driver.

    ![schedulers](http://books.gigatux.nl/mirror/kerneldevelopment/0672327201/images/0672327201/graphics/13fig03.gif)

We can tune the specifics of the scheduler by modifying different configurations in `/sys/block/sda/queue/iosched`.

To check whether the device is SSD (`0` is SSD):

```bash
cat /sys/block/$device/queue/rotational
```
