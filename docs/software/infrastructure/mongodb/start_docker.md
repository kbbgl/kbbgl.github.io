---
slug: start-mongo-container
title: Start MongoDB Docker Container
authors: [kbbgl]
tags: [mongo,container,docker,image,db]
---

## Start MongoDB Container

```bash
export MONGODB_VERSION=6.0-ubi8

docker run --name mongodb -d -p 27017:27017 mongodb/mongodb-community-server:$MONGODB_VERSION
```

With data persistence:

```bash
docker run --name mongodb -d -p 27017:27017 -v $(pwd)/data:/data/db mongodb/mongodb-community-server:$MONGODB_VERSION
```

## Stop MongoDB Container

```bash
docker stop mongodb && docker rm mongodb
```
