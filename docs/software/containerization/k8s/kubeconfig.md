---
title: Kubeconfig (kubectl configuration)
slug: kubeconfig
tags: [devops, k8s, kubernetes, kubectl, kubeconfig, security]
authors: [kbbgl]
---

# Kubeconfig (kubectl configuration)

A **kubeconfig** is a YAML file that tells `kubectl` (and other Kubernetes clients) **which cluster to talk to** and **how to authenticate**. By default, `kubectl` reads:

- `~/.kube/config`

## What's inside a kubeconfig

A kubeconfig typically contains:

- `clusters`: API server `server:` endpoint + the cluster CA (`certificate-authority-data` or `certificate-authority`)
- `users`: how to authenticate (commonly a `token`, client cert/key, or an `exec` auth plugin)
- `contexts`: a named binding of **(cluster + user + optional namespace)**
- `current-context`: which context `kubectl` uses by default

Minimal example shape:

```yaml
apiVersion: v1
kind: Config
clusters:
- name: my-cluster
  cluster:
    server: https://my-k8s-cluster.example.com
    certificate-authority-data: <base64-ca-cert>
users:
- name: my-user
  user:
    token: <bearer-token>
contexts:
- name: my-context
  context:
    cluster: my-cluster
    user: my-user
current-context: my-context
```

## Common workflows

### List contexts and switch clusters

```bash
kubectl config get-contexts -o name
kubectl config use-context my-context
```

Quick sanity check:

```bash
kubectl get nodes
```

### Use a specific kubeconfig (without changing your default)

Use-case: CI jobs, one-off access to prod, testing a new config you were given.

- **Inline for a single command**:

```bash
kubectl get nodes --kubeconfig "$HOME/.kube/dev_cluster_config"
```

- **Via environment variable**:

```bash
export KUBECONFIG="$HOME/.kube/dev_cluster_config"
kubectl get nodes
```

### Precedence (what wins)

In practice, these commonly override each other in this order:

1. `--kubeconfig ...` on the command line
2. `KUBECONFIG=...` environment variable
3. default `~/.kube/config` (and its `current-context`)

### Merge multiple kubeconfig files into one

Use-case: you have separate `kubeconfig`s for dev/test/prod, but want a single `~/.kube/config` with multiple contexts.

From `~/.kube` (or using full paths), you can merge and flatten:

```bash
KUBECONFIG=config:dev_config:test_config kubectl config view --merge --flatten > config.new
mv "$HOME/.kube/config" "$HOME/.kube/config.old"
mv "$HOME/.kube/config.new" "$HOME/.kube/config"
```

Verify:

```bash
kubectl config get-contexts -o name
```

### Minify (show only the active context)

Use-case: quickly inspect exactly what `kubectl` will use right now.

```bash
kubectl config view --minify
```

### Remove a stale context

Use-case: cluster was deleted, but the context still exists locally.

```bash
kubectl config get-contexts -o name
kubectl config delete-context my-old-context
```

## Creating a least-privilege kubeconfig (concept)

Common pattern: create a dedicated identity (often a `ServiceAccount`) with RBAC permissions restricted to what’s needed (namespace-scoped when possible), then generate a `kubeconfig` that authenticates as that identity.

Typical building blocks:

- `ServiceAccount` in a namespace
- `Role/ClusterRole` granting specific verbs/resources
- `RoleBinding/ClusterRoleBinding` binding the `ServiceAccount` to the role
- A **token** for the `ServiceAccount` (Kubernetes v1.24+ commonly uses an explicit token `Secret` or projected tokens)

Once you have the cluster endpoint, CA data, and token, you can assemble a `kubeconfig` like the minimal example above.

## Security best practices

Kubeconfigs often contain **tokens and/or client keys**. Treat them like credentials.

Lock down permissions:

```bash
chmod 600 "$HOME/.kube/config"
chmod 700 "$HOME/.kube"
```

Avoid accidental commits:

- Keep `kubeconfigs` out of Git repos
- Add patterns to ignore files named like `kubeconfig*` (either per-repo `.gitignore` or your global gitignore)

## Sources

- [Kubeconfig File Explained (Examples, Usage and Configuration) (DevOpsCube)](https://devopscube.com/kubernetes-kubeconfig-file/)
