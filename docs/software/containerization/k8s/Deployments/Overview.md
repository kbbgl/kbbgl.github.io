## Overview

The default controller when running `kubectl run` is a `Deplyoment`. When a `Deployment` is added to a cluster, the controller will create a `ReplicaSet` and a `Pod` automatically. We can update the `Deployment` using:

```bash
kubectl edit deployment $deployment_name
#OR
kubectl apply -f modified_deployment.yaml
```

`ReplicationControllers` (RC) ensure that a specified number of `Pod` replicas are running at any time. 