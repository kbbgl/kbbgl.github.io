## Volume Introduction

A `Pod` specification can declare one or more volumes and where they are made available. Each of them requires:

- A name
- A type
- A mountpoint

The same volume can be made available to all containers in a `Pod`.

A particular access mode is part of a `Pod` request. The 3 access modes are:

- `ReadWriteOnce`: allows `rw` by a single node.
- `ReadOnlyMany`: allows `r` by multiple nodes.
- `ReadWriteMany`: `rw` by multiple nodes.

When a volume is requested. a local `kubelet` uses the `kubelet_pods.go` script to map the raw devices, determine and make the mount point for the container and create a symlink on the host node fs to associate storage to the container. The API server makes a request for the storage to the `StorageClass` plugin.

