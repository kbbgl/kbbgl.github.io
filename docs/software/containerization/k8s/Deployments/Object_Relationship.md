# Object Relationship

The shapes represent controllers (or watch loops) that run threads of `kube-controller-manager`. Each controller queries the `kube-apiserver` for the current state of the object they track. The state of each object on a worker node is sent back from the local `kubelet` after the `kubelet` daemon receives information back from the container engine.
