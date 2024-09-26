---
slug: remote-command-login-exec-win
title: Remote Command Login Execution in Windows
description: Remote Command Login Execution in Windows
authors: [kbbgl]
tags: [cybersecurity,offensive,windows,remote,login]
---

https://docs.microsoft.com/en-us/sysinternals/downloads/psexec

```bash
psexec \\mylap cmd
```

```bash
psexec pentest:'PentestPassword'@10.10.10.10
```

Optional, can be run without `metasploit` by installing [impacket](https://github.com/SecureAuthCorp/impacket). Alternatives to `psexec` are `smbexec` or `wmiexec`.