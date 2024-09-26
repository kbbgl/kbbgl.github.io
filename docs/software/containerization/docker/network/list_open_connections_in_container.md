---
slug: list-open-connections-container
title: List Open Connectios in Container
authors: [kbbgl]
tags: [docker, network, containers]
---

## Get container ID (leftmost column)

```bash
docker ps -a | grep translation | grep -v init | grep Up | grep -v pause
d72db97ab0f1        9878c6437bb8                               "/etc/bootstrap.sh -d"    4 hours ago         Up 4 hours                                         k8s_translation_translation-77c6c76954-kgs28_app_e67d6658-9fbb-46bc-8582-b3ce510f295f_18
```

## Get container PID

```bash
docker inspect -f '{{.State.Pid}}' d72db97ab0f1
28886
```

## Run `netstat` in container

```bash
nsenter -t 28886 -n netstat -ano
Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       Timer
tcp        0      0 0.0.0.0:8070            0.0.0.0:*               LISTEN      off (0.00/0/0)
tcp        0      0 10.233.102.173:44925    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:44599    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:8070     10.233.102.161:48740    TIME_WAIT   timewait (52.48/0/0)
tcp        0      0 10.233.102.173:41985    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:37003    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:36123    10.233.102.175:5672     ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:42213    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:32965    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:39525    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:35733    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:41297    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
tcp        0      0 10.233.102.173:43581    10.233.58.95:2181       ESTABLISHED off (0.00/0/0)
Active UNIX domain sockets (servers and established)
Proto RefCnt Flags       Type       State         I-Node   Path
unix  2      [ ACC ]     STREAM     LISTENING     378061   /tmp/kestrel_6e404941c481404b8672b9eb9463ae06
unix  3      [ ]         STREAM     CONNECTED     379101
unix  3      [ ]         STREAM     CONNECTED     378063   /tmp/kestrel_6e404941c481404b8672b9eb9463ae06```
