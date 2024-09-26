---
slug: mongodb-get-primary-member-rs
title: Get Primary Node from Replica Set
authors: [kbbgl]
tags: [mongodb, replica_set, high_availability, node]
---

```javascript
rs.status().members.filter(member => member.stateStr === "PRIMARY")[0].name
```
