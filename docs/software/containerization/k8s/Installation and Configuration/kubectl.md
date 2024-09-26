## `kubectl`

`kubectl` is used to configure and manage the cluster.

It will use the `~/.kube/config` file as the configuration file which includes the cluster IP, the credentials and the context.

A **context** is a combination of a cluster and user credentials.
The context can be changed using:

```bash
kubectl config use-context $SOME_CONTEXT
```