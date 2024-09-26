---
slug: rabbitmq-cookie-location-windows
title: Erlang Cookie Location in Windows
authors: [kbbgl]
tags: [rabbitmq, windows, cookie]
---

```batch
type C:\Windows\system32\config\systemprofile\.erlang.cookie
```

If there are any authentication issues when running `rabbitmqctl`, copy cookie from above to `%USERPROFILE%`.
