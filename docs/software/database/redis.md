---
slug: redis-cheatsheet
title: Redis Database Cheatsheet
description: A cheatsheet for Redis database.
authors: [kgal-akl]
tags: [cheatsheet, cli, redis, database, cache]
---


## Access

### TLS Disabled
When When TLS is not enabled:

```bash
redis-cli --pass $REDIS_PASSWORD --no-auth-warning
```

### TLS Enabled

```bash
redis-cli --tls --cert   /opt/app/cache/certs/tls.crt --key    /opt/app/cache/certs/tls.key --cacert /opt/app/cache/certs/ca.crt --pass $REDIS_PASSWORD --no-auth-warning
```

## Get Total Number of Keys

```bash
redis-cli --pass $REDIS_PASSWORD --no-auth-warning keys "*" | wc -l
```

## Get Hash Key

```bash
redis-cli --pass $REDIS_PASSWORD --no-auth-warning hgetall $KEY_NAME
```

## Monitor

With monitor we can observe every single operation done (e.g. key set, get, `hgetall`, delete):

```bash
redis-cli --pass $REDIS_PASSWORD --no-auth-warning MONITOR
```