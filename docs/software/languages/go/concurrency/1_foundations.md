---
slug: foundations
title: Foundations
description: Fill me up!
authors: [kbbgl]
tags: [go,concurrency]
---

Multicore processors allow us to use concurrency. In addition, cloud computing services give developers to more computational power. They allow us to scale horizontally (distribute the load over multiple processing resources) instead of vertically (adding more processing power to improve existing resources).

Concurrency allows us to improve responsiveness when accepting user input since we don't have to wait for a task to finish.

Go uses a lightweight construct called a **goroutine** to model the basic unit of concurrent execution. The principle is that if you need multiple things to be done concurrently, create as many goroutines as needed without thinking about resource allocation. Go provides us with many abstractions that allow us to coordinate the concurrent executions on a common task. One of them is a **channel** which allow two or more goroutines to pass messages to each other. This enables the exchange of information and synchronization of the multiple executions in an intuitive manner.
