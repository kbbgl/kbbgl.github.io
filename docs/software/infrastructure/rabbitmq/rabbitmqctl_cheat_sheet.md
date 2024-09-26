---
slug: rabbitmq-ctl
title: rabbitmqctl Cheat Sheet
authors: [kbbgl]
tags: [rabbitmq, rabbitmqctl, cheatsheet]
---


## List Queues

```bash
rabbitmqctl list_queues name consumers messages
```

## Get queues with no consumers

```bash
rabbitmqctl list_queues name consumers | grep -E $'\t0'

my_app-7cd44df77d-xcqwr 0
my_app-7cd44df77d-qf4qw 0
```

## Delete Queue

```bash
rabbitmqctl delete_queue my_app-7cd44df77d-qf4qw
Deleting queue 'my_app-7cd44df77d-qf4qw' on vhost '/' ...
Queue was successfully deleted with 234 messages
```
