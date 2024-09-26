---
slug: mongodb-update-all-docs-in-collection
title: Update All Documents in Collection
authors: [kbbgl]
tags: [mongodb, collection, update]
---


# Update All in Collection

```javascript
db.getCollection('users').updateMany(
	{},
	{ $set: { hash: {...}}
)
```
