# Liveness/Readiness Probes

```yaml
livenessProbe:
  failureThreshold: 3
  httpGet:
    path: /liveness
    port: 14992
    scheme: HTTP
  initialDelaySeconds: 60
  periodSeconds: 20
  successThreshold: 1
  timeoutSeconds: 10
readinessProbe:
  failureThreshold: 3
  httpGet:
    path: /readiness
    port: 14992
    scheme: HTTP
  initialDelaySeconds: 10
  periodSeconds: 10
  successThreshold: 1
  timeoutSeconds: 5
 ```

`initialDelaySeconds`: Number of seconds after the container has started before liveness or readiness probes are initiated. Defaults to 0 seconds. Minimum value is 0.

`periodSeconds`: How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.

`timeoutSeconds`: Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.

`successThreshold`: Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.

`failureThreshold`: When a probe fails, Kubernetes will try `failureThreshold` times before giving up. Giving up in case of liveness probe means restarting the container. In case of readiness probe the Pod will be marked Unready. Defaults to 3. Minimum value is 1.
