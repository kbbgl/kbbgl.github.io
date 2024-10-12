---
slug: reliable-scalable-maintainable
title: Reliable, Scalable, Maintainable
description: Chapter 1 from Designing Data Intensive Applications
authors: [kbbgl]
tags: [scalability,system,design,reliability,maintenability]
---

Many apps are data-intensive and are built with standard building blocks that provide commonly needed functionality:

* Store data so it can be queried (databases)
* Remember results for expensive operations to speed up reading (caches)
* Allow users to search data by keywords/filter (search indices)
* Send an async message to another process (stream processing)
* Periodically crunch a large amount of accumulated data (batch processing)

Three concerns that are important in most software systems.

## Reliability

The system should continue to work correctly even in the face of adversity (hw, sw faults, human error). Its performance is good enough for the required use case under expected load and data volume. It prevents unauthorized access and abuse.

Faults mean that one component of the system deviated from its spec whereas a failure is when a system as a whole stops pance (e.g. Chaos Monkey developed by Netflix).

### Hardware Faults

Hard disk crash, faulty RAM, unplug a network cable.

The first response is to add redundancy to the individual hardware components. Disks can be set up in a RAID configuration or sharding. Power can be backed up by UPS/generators. Entire machines/nodes can be taken down so we can add redundant nodes.

### Software Faults

Software faults can cause many more system failures than uncorrelated hardware faults. Some examples of software faults are bugs caused by unexpected/bad input, process taking too much of a shared resource, service becomes unresponsive or returns unexpected results, cascading failures.

The solutions are usually to thoroughly end-to-end test, isolate processes, crash and recovery, monitoring.

### Human Errors

To prevent/minimize human errors we can design systems in a way that minimizes opportunity for error with well-designed abstractions, APIs, sandbox environments, thorough tests (unit to integration).

Other approaches can be to easily recover from human errors such as rolling back changes, gradually roll out new releases, set up performance metrics and error rates with telemetry.

## Scalability

As the system grows (data, traffic, complexity), there should be reasonable ways of dealing with the growth.

### Describing Load

We need to be able to describe the current load on the system before discussing growth. Load can be described by **load parameters**. It can be requests-per-second for web servers, read-to-write ration for databases, simultaneous active users in a chat room, hit rate on a cache.

### Describing Performance

After describing load, we can investigate what happens when the load increases. We can increase a load parameter and keep the system resources constant and see how it affects the performance. We can also change a load parameter and see what it would take to keep the performance the same.

Performance metrics vary according to the systems. In online systems the most important factor is **response time**. Response time can vary depending on the load, network latency, page faults, etc.
It's usually best to use **percentiles** to describe response time by sorting the fastest to slowest responses and finding the median. If the median is 200ms, this means that half the requests took less than 200ms and half took more than that.

## Maintainability

Over time, many different people will work on the system. They should all be able to work on it productively.

We should design software in such a way that it will minimize pain during maintainance. There are 3 design principles for software systems.

## Operability

Make it easy for operations teams to keep the system running smoothly. They are vital to the system running smoothly. A good operations team is typically responsible for:

* Monitoring the health of the system and quickly restoring service if it goes into bad state.
* Tracking down cause of problems or degradation.
* Patch platforms with updates and security.

## Simplicity

Make it easy for new engineers to understand the system by removing as much complexity as possible.

Complexity can come in different forms such as tight coupling of modules, tangled dependencies, inconsistent naming and terminology, hacks.

One of the best tools to solve complexity is by **abstraction**.

## Evolvability

Make it easy for engineers to make changes to the system in the future.

Agile working patterns provide a good framework for adapting to change. Also TDD and refactoring.
