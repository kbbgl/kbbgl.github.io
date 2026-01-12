---
slug: helm-cheatsheet
title: helm Cheatsheet
description: A cheatsheet for helm.
authors: [kgal-akl]
tags: [cheatsheet, helm, k8s, deploy]
---

## Releases

### List Releases in All Namespaces

```bash
helm list --all-namespaces
```

### Upgrading Release

```bash
helm upgrade --install akeyless-gw akeyless/akeyless-api-gateway -f deploy/akl-gw/values.yaml
```

## Repo

### Adding Repo

```bash
helm repo add akeyless https://akeylesslabs.github.io/helm-charts
helm repo update
```

### Updating Repo

```bash
helm repo akeyless update
```

### Updating Repo with Git Change
If we want to update a repo to use a specific git ref (e.g. branch, commit), we can use [`helm git`](https://github.com/aslafy-z/helm-git). For example, if we pushed a change to a chart in a remote branch (`fix-allowedAccessPermissions-type-api-gw`) and we want to test it out, we can add a new helm repo:

```bash
helm repo add akeyless-dev git+https://github.com/akeylesslabs/helm-charts@charts?ref=fix-allowedAccessPermissions-type-api-gw
```

## Charts

### Checking Chart Versions

```bash
helm search repo $MY_REPO/$MY_CHART --versions
```

### Upgrading Chart to Specific Version

```bash
helm upgrade $MY_RELEASE_NAME $MY_REPO/$MY_CHART -f $MY_CHART_VALUES.yaml --version $NEW_CHART_VERSION
```

### Get Chart Values

```bash
helm show values akeyless/akeyless-api-gateway | yq
```

### Overriding Values

Helm charts use customizable `values.yaml` files that allow you to override default configurations using `--set` flags.

```bash
helm install my-postgres bitnami/postgresql \
  --set global.postgresql.auth.postgresPassword=yourpassword \
  --set global.postgresql.auth.username=myuser \
  --set global.postgresql.auth.password=mypassword \
  --set global.postgresql.auth.database=mydatabase
```


### Generating Post-Processed YAML Manifests

In some situations we will want to generate the post-processed YAML file(s) manifests that include all the Kubernetes resources that `helm` uses
to later `kubectl apply` them to the cluster.

To render the manifest into one file, we can use:

```bash
helm template $RELEASE_NAME $CHART_PATH_OR_NAME --namespace $NAMESPACE > $OUTPUT_FILENAME.yaml

# example for postgres
helm template mypostgres bitnami/postgresql --namespace database > postgresql.yaml
```

To render separate manifests into a directory instead of one file:

```bash
helm template $RELEASE_NAME $CHART_PATH_OR_NAME --namespace $NAMESPACE --output-dir ./generated-manifests
```

```bash
helm template mypostgres bitnami/postgresql --namespace database --output-dir ./generated-manifests
```
