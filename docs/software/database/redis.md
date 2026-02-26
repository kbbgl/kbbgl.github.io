---
slug: redis-cheatsheet
title: Redis Cheatsheet
description: Practical Redis CLI commands (access, inspect keys, monitoring, admin).
authors: [kgal-akl]
tags: [redis, cache, infrastructure, services, cheatsheet, cli, database]
---


## Access (`redis-cli`)

### TLS disabled

```bash
redis-cli --pass "$REDIS_PASSWORD" --no-auth-warning
```

### TLS enabled

```bash
redis-cli \
  --tls \
  --cert   /opt/app/cache/certs/tls.crt \
  --key    /opt/app/cache/certs/tls.key \
  --cacert /opt/app/cache/certs/ca.crt \
  --pass "$REDIS_PASSWORD" \
  --no-auth-warning
```

## List keys

### List all keys (use carefully)

```
KEYS *
```

### Iterate keys (preferred for production)

```
SCAN 0

SCAN 0 MATCH user:* COUNT 100
```

## Inspect key details

Get the key type:
```
TYPE <key>
```

Then retrieve its value based on its type.

### Get Value of String-type Key

```
GET <key>
```

### Get Value of List-type Key

```
LRANGE <key> 0 -1
```

### Get Value of Set-type Key

```
SMEMBERS <key>
```

### Get Value of Sorted Set-type Key

```
ZRANGE <key> 0 -1 WITHSCORES
```

### Get Value of Hash-type Key

```
HGETALL <key>
```

## Count keys

### Total number of keys (by scanning)

```bash
redis-cli --pass "$REDIS_PASSWORD" --no-auth-warning --scan | wc -l
```

## Database info

Check the number of keys in each Redis database:
```
INFO keyspace

# Keyspace
db0:keys=450,expires=12,avg_ttl=500000
db1:keys=200,expires=5,avg_ttl=300000
```

`db0` has 450 keys, `db1` has 200 keys.
Some keys have expiration times.

## Switch Databases

If Redis has multiple databases (default: `db0`), to work with `db0`
```
SELECT 0
```

To switch to `db1`:
```
SELECT 1
```

## Delete Keys

```
DEL <key>
```

To delete all keys in the currently-selected database:

```
FLUSHDB
```

To delete all keys in Redis:

```
FLUSHALL
```

## Monitor

Observe every operation (e.g. key set/get, `HGETALL`, deletes):

```bash
redis-cli --pass "$REDIS_PASSWORD" --no-auth-warning MONITOR
```

## Configuration

### Set Log Level
```
CONFIG SET loglevel debug
```