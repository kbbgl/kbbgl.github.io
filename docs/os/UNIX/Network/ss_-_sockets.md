# Sockets

Print all socket information:

```bash
ss
```

Print all listening sockets:

```bash
ss -l
```

Print all TCP connections:

```bash
ss -t
```

Print sockets with PID

```bash
ss -p
```

Filter connections by IP type:

```bash
ss -4
```

summary stats per protocol:

```bash
ss -s

Total: 3471 (kernel 0)
TCP:   1051 (estab 317, closed 708, orphaned 0, synrecv 0, timewait 201/0), ports 0

Transport Total     IP        IPv6
*   0         -         -
RAW   1         0         1
UDP   60        7         53
TCP   343       296       47
INET   404       303       101
FRAG   0         0         0
```
