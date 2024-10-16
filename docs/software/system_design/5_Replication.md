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

## Synchronous vs Asynchronous Replication

The advantage of using synchronous replication is that the followers are guaranteed to have up-to-date copies of the data. The disadvantage is that the leader must block all writes until the synchronous replica is available again. Therefore, most systems have one of the followers as synchronous where the rest are async and once one of the sync followers becomes unavailable, another async becomes snyc. 

Leader-based replication is configured to be completely async. This means that if the leader fails, any writes that haven't been replicated are lost.

## Handling Node Outages

If a follower fails, it can read the log to recover and then request the leader for any changes since the last log entry.

If a leader fails, a new leader needs to be promoted, followers need to be updated of the new leader. This is called a **failover**. In a situation where there's a **split brain**, two or more nodes believe they are the new leader.

## Replication Logs

There are several replicaton methods that are popular.

**Statement-based replication** has the leader log every write request (statement) that it executes and sends that to the followers.

**Write-ahead log (WAL)** used by PostgresSQL and Oracle. The leader sends the WAL to the followers which use it to reconstruct the replica.

**Logical (row-based) log replication** allows decoupling the replication log from the storage engine which is the main disadvantage using WAL. The replication log in this case is called the **logical log** whereas the storage engine data representation is called the **phyiscal log**.

**Trigger-based replication** allows applications, stored procedures or triggers to read the database log and decide what/when to replicate. A trigger allows you to register custom application code that is automatically triggered when there's a write to the database.
