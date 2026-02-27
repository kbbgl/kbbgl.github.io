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

We can also delete them in batches if we don't access/permission/ to `FLUSHALL` or it's disabled:

```bash
BATCH_SIZE=100
redis-cli --tls \
--cert /opt/app/cache/certs/tls.crt \
--key /opt/app/cache/certs/tls.key \
--cacert /opt/app/cache/certs/ca.crt \
-a "$REDIS_PASSWORD" --no-auth-warning \
--scan | xargs -d '\n' -L $BATCH_SIZE redis-cli --tls \
--cert /opt/app/cache/certs/tls.crt \
--key /opt/app/cache/certs/tls.key \
--cacert /opt/app/cache/certs/ca.crt \
-a "$REDIS_PASSWORD" --no-auth-warning DEL
```

## Persistence

```bash
redis-cli -a "$REDIS_PASSWORD" INFO persistence
redis-cli -a "$REDIS_PASSWORD" CONFIG GET save
```

## Configuration

### Set Log Level
```
CONFIG SET loglevel debug
```

## Debugging

### Monitor

Observe every operation (e.g. key set/get, `HGETALL`, deletes):

```bash
redis-cli --pass "$REDIS_PASSWORD" --no-auth-warning MONITOR
```

### Ping a Remote Redis Server

```bash
redis-cli --tls \
--cacert /opt/app/cache/certs/ca.crt \
--cert /opt/app/cache/certs/tls.crt \
--key /opt/app/cache/certs/tls.key \
-h $REMOTE_REDIS_HOST -p 6379 \
-a "$REDIS_PASS" PING
```

### Check Latency

```bash
redis-cli -a "$REDIS_PASSWORD" LATENCY DOCTOR
redis-cli -a "$REDIS_PASSWORD" SLOWLOG GET 128
```

### List Clients

The command provides all the clients that communicated with the Redis server:

```bash
redis-cli CLIENT LIST
```

It provides a file descriptor (`fd`) and an IP address/port of the client. It is very useful when there are connection/TLS issues and we see prints such as in the Redis server logs:

```
Error accepting a client connection: error:0A000126:SSL routines::unexpected eof while reading (conn: fd=259)
```

We see here that the client got assigned `fd=259` so we can use the command above to see the IP of that specific client belonging to that file descriptor.