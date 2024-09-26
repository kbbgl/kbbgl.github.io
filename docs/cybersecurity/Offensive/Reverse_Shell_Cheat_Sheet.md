---
slug: reverse-shell-cheat-sheet
title: Reverse Shell Cheat Sheet
description: Reverse Shell Cheat Sheet
authors: [kbbgl]
tags: [cybersecurity,offensive,reverse,shell,netcat]
---

[Source](http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet)

After shell created, run:

```bash
nc -nvlp $PORT
```

If reverse shell does not include `tty`, we can use [this](https://netsec.ws/?p=337) list of commands.
