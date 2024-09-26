# Analyzing Coredump with GDB

## Create Coredump

### Analyze Coredump

1) Copy coredump file into EC Pod:

```bash
kubectl cp core.12510.160763620 ec-test-qry-...:/tmp/
```

2) From within the container, Initialize debugger

```bash
kubectl exec ec-test-qry... -it -- bash

> gdb /opt/app/monetdb/bin/mserver5 core.12510.160763620

# With arguments
gdb --args /opt/app/monetdb/bin/mserver5 --zk_system_name=S1 --zk_address=bi-engine-s /opt/app/monetdb/bin/mserver5 core.12510.160763620

```
