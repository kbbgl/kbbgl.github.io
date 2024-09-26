## Objects

`Node`: represents a machine (physical or virtual) that is part of a cluster.

`Service Account`: Provides an identifier for processes running in a `Pod` to access the API server and performs actions that it is authorized to.

`Resource Quota`: Allows defining quotas per namespace for limiting resources, e.g. number of `Pod`s scheduled.

`Endpoint`: Represents a set of IPs for `Pod`s that match a particular service.

`Deployment`: Controller which manages the state of the `ReplicaSets` and the `Pod`s within. 

`ReplicaSet`: Orchestrates individual `Pod` lifecycle and updates. 

`Pod`: Lowest manageable unit, runs application containers.

`DaemonSet`: Controller that runs on every node. When a node is added/removed, the controller ensures that the `Pod` is removed from the node as well. 

`StatefulSet`: `Pod`s deployed using a `StatefulSet` use the same `Pod` specification. How this is different than a `Deployment` is that a `StatefulSet` considers each Pod as unique and provides ordering to `Pod` deployment.

### `batch` API Group
- `Job`: Used to run `Pod`s to completion. If it fails, it will restart until number of completions is reached.
- `CronJob`: Similar to Linux jobs with the same time syntax.

### `autoscaling` API Group, 

- Horizontal Pod Autoscalers (**HPA**) resources. They automatically scale `Replication Controllers`, `ReplicaSets` or `Deployments`. 
- Cluster Autoscaler (**CA**) adds/removes nodes to the cluster based on inability to deploy `Pod`s or having nodes with low utilization for at least 10 minutes. When using this type of autoscaler, we use `cluster-autoscaler` commands.

### `RBAC` API Group

Used for access control to API.

- `ClusterRole`
- `Role`
- `ClusterRoleBinding`
- `RoleBinding`