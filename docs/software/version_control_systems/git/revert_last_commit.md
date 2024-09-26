---
slug: revert-last-commit
title: Revert Last Commit
authors: [kbbgl]
tags: [git,revert,commit]
---

## Revert Last Commit
```bash
git revert $(git --no-pager log | head -1 | cut -d" " -f2 | tr -d "[[:space:]]")
```
