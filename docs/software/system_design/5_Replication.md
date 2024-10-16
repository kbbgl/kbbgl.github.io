---
slug: replication
title: Replication
description: Chapter 5 from Designing Data Intensive Applications
authors: [kbbgl]
tags: [replication,system,design,distributed]
---

**Replication** means keeping a copy of the same data on multiple machines in the same network.

Some reasons for replicating are:

- Reduce latency by keeping data geographically closer to users.
- High-availability in case a component of the system goes down.
- Increase read throughput by horizontally scaling the amount of machines that can serve read queries.

The main challenge in replication is in handling changes to the replicated data. There are 3 main algorithms for replicating changes between nodes, **single-leader**, **multi-leader**, **leaderless**.

## Leaders and Followers

Every write to the database needs to be processed by every replica. To solve this problem,  the **leader-based** (aka **active/passive**, **master-slave**) replication is used.

The designated leader receives all write operations and writes it to its local storage. It then sends the data change to all the followers as part of a **replication log** or change stream. The follower applies the replication log to its own local storage. Clients can read from any replica but writes are only sent to the leader.
