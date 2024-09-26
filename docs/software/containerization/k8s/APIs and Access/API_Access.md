## API Access

Kubernetes architecture revolves around REST APIs.

The main agent for communication between the cluster agents from outside the cluster is the `kube-apiserver`.

## Checking Access

To check which user has permission to do what:

```bash
kubectl auth can-i create deployments
# yes

kubectl auth can-i create deployments --as bob
# no

kubectl auth can-i create deployments --as bob --namespace developer
# yes

```

There are 3 APIs which can be applied to set who and what can be queried:

- `SelfSubjectAccessReview`: Access review for any user, helpful for delegating to others.
- `LocalSubjectAccessReview`: Review is restricted to a specific namespace.
- `SelfSubjectRulesReview`: A review which shows allowed actions for a user within a particular namespace.

All namespace configuration files are found in:

```
/home/app/.kube/cache/discovery/10.50.42.95_6443/servergroups.json
```
