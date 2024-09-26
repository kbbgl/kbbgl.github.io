---
slug: how-send-messages-between-computers-in-terminal
title: How To Send Messages in the Terminal between Computers
authors: [kbbgl]
tags: [linux,terminal,send,messages]
---

## On Receiving Computer

```bash
nc -l $PORT
```

## On Sending Computer

```bash
nc $RECEIVING_HOST_IP

# type text...
```
