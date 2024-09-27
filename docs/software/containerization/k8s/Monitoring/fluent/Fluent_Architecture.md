### Standard I/O

```
stdin => app
stdout <= app
stderr <= app
```

### Logging & Docker Containers

App writes message to stdout

```
app => stdout “hi”
```

message “hi” is encapsulated in a JSON map by docker engine:

“hi” becomes

```json
{ 
    “log”: “Hi”, 
    “stream”: “stdout”, 
    "time": “2020-08-13T16:16:05.53515125”
}
```

so JSON has context (`stdout`, `log message`, `timestamp`)

this message is appended to the container log file:

```
/var/log/containers/containername.log (symlink to /var/lib/docker/containers/CONTAINER_HASH/CONTAINER_HASH-json.log)
```

### Docker Log Streams

```bash
docker run -d busybox echo -n "helloworld"
```

this will return the hash as it's running with a daemon flag

```bash
sudo cat /var/lib/docker/containers/CONTAINER_HASH/CONTAINER_HASH-json.log | jq
```

### Logging & Kubernetes

In addition to Docker context, there's:

* Pod Name and Pod ID
* Namespace
* Node
* Labels
* Annotations

So the fluent-bit has to process:

* Container name and container id (received from `fs`/`journal`)
* Pod Name and Pod ID
* Namespace
* Node
* Labels
* Annotations

```
Logs source (k8s API server + fs) <=> log processor <=> Storage (logzio)
```

so the log processor (`fluentbit`) needs to correlate the fs (container name + id) from docker with the resources created by k8s API server.

### Log Processing in Kubernetes

log process runs as a `DaemonSet`.

every node has N Pods which write to `/var/log/containers/*` which is a `symlink` to docker engine logs (`/var/lib/docker/..`)

every node has a `fluent-bit` Pod running on it

`fluent-bit` pod does 2 things:

1) reads each log in `/var/log/containers/*`

2) looks up the metadata from k8s API server

### `fluent-bit` config

------------------

Sections:

* Service

```yaml
[SERVICE]
     Flush           5 # flush time in seconds. every 5 seconds the engine flushes records to output plugin
     Daemon          off # whether to run as bg
     Log_Level       debug 
     HTTP_Monitoring On.  # enable web service monitoring interface
     HTTP_Port       2020 # TCP port running of monitoring interface
```

* Input: defines a source of data.

```yaml
[INPUT]
     Name cpu.      # name of input plugin
     Tag  my_cpu    # tag for associated records coming from this plugin
```

* Output: defines a destination records should follow after a tag match

```yaml
[OUTPUT]
     Name  stdout # name of output plugin
     Match my*cpu # pattern to match certain record's tag. 
```
