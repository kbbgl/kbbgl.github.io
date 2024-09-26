---
slug: mongodb-update-doc
title: Update Queried Document
authors: [kbbgl]
tags: [mongodb, query, update]
---

```javascript
db.users.update(
    { email: 'user@domain.com' },
    { $set: { lastActivity: '' } }
);
```
