## Using `DaemonSets`

This controller ensures that a single pod exists on each node in the cluster. Every `Pod` uses the same image. Should a new node be added, the `DaemonSet` controller will deploy a new `Pod` on your behalf. Should a node be removed, the controller will delete the `Pod` also. 

