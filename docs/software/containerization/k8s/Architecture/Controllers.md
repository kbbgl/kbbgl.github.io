## `Controller`

A `Controller` is an agent running in a loop which regulates the system and is responsible for making the current state reach the desired state. It will send messages to the `kube-apiserver` to perform the necessary actions (create/destroy) on the set resource (e.g. create a `Pod`, destroy a `Node`).