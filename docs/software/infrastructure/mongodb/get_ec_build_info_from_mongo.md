---
slug: mongodb-query-sort-limit-regex
title: Query with Regex, Sort, Limit
authors: [kbbgl]
tags: [mongodb, query, sort, json]
---

```javascript
db.collection_name
    .find({
        type: "buildEnded",
        message: /Build failed/
    }, 
    {
        cubeId: 1,
        _id: 0,
        timestamp: 1,
        message: 1
    })
    .sort({
            timestamp: -1
        })
    .limit(10)
    .pretty()
```
