## Annotations

Annotations allow for metadata to be included with an object that may be helpful outside of the Kubernetes object interaction. 

Here are some examples of information that could be recorded in annotations:

- Fields managed by a declarative configuration layer. Attaching these fields as annotations distinguishes them from default values set by clients or servers, and from auto-generated fields and fields set by auto-sizing or auto-scaling systems.

- Build, release, or image information like timestamps, release IDs, git branch, PR numbers, image hashes, and registry address.

- Pointers to logging, monitoring, analytics, or audit repositories.

- Client library or tool information that can be used for debugging purposes: for example, name, version, and build information.

- User or tool/system provenance information, such as URLs of related objects from other ecosystem components.

- Lightweight rollout tool metadata: for example, config or checkpoints.

- Phone or pager numbers of persons responsible, or directory entries that specify where that information can be found, such as a team web site.

- Directives from the end-user to the implementations to modify behavior or engage non-standard features.

To annotate `Pod`s within a namespace:

```bash
kubectl annotate pods --all description="Production Pods" -n prod

kubectl annotate --overwrite pod webpod description="Old Production Pods" -n prod

kubectl -n prod annotate pod webpod description-
```

