# Requests and Limits

```yaml
containers:
  - name: app
    image: super.mycompany.com/app:v4
    env:
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
  - name: log-aggregator
    image: super.mycompany.com/log-aggregator:v6
    resources:
      requests:
        memory: "64Mi"
        cpu: "250m"
      limits:
        memory: "128Mi"
        cpu: "500m"
```

The following Pod has two Containers. Each Container has a request of `0.25` cpu and `64MiB` (226 bytes) of memory.

Each Container has a limit of `0.5` cpu and `128MiB` of memory.

You can say the Pod has a request of `0.5` cpu and `128 MiB` of memory, and a limit of `1` cpu and `256MiB` of memory.
