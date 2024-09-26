## `Service`

`Service`s are abstractions which define a set of `Pod`s and a policy by which to access them. They can be seen as microservices handling traffic to particular endpoints which applications expose. The traffic may be relayed using a `NodePort` or `LoadBalancer` to balance traffic between different replicas.

- `Service`s connect `Pods` together. Ensures that `Pod`s relying on other `Pod`s can always communicate on the expected IP/hostname.
- Exposes `Pod` externally.
- Defines `Pod` access policy.

