---
slug: match-uuid
title: Match UUID4 
authors: [kbbgl]
tags: [regex, uuid, grep]
---

```bash
grep -E "\w{8}-\w{4}-\w{4}-\w{4}-\w{12}" /path/to/file
```