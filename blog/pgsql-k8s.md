---
slug: pgsql-k8s
title: Deploying PostgresSQL in Kubernetes 
description: Deploying a PostgresSQL database in a Kubernetes cluster
authors: [kgal-akl]
tags: [k8s,kubernetes,database,storage,postgres,web,webapplication,webdevelopment]
---

In this tutorial we deploy a PostgresSQL database in Kubernetes cluster.

<!-- truncate -->

## Set Up Postgres

We first create a new namespace to hold the database-related services.

```bash
kubectl create namespace database
kubectl config set-context --current --namespace database
```

The first step is to create a `PersistenceVolume` and `PersistentVolumeClaim`:

```yaml title
apiVersion: v1
kind: PersistentVolume
metadata:
  name: postgresql-pv
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data"

---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgresql-pv-claim
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

We add them to the namespace:

```bash
kubectl apply -f pgsql-pv.yaml&&kubectl apply -f pgsql-pvc.yaml
```

We then install Postgres using the `bitnami/postgresql` chart:

```bash
helm install psql-test bitnami/postgresql \
    --set persistence.existingClaim=postgresql-pv-claim \
    --set volumePermissions.enabled=true \
    --set primary.service.type=LoadBalancer
```

We can check that the database service is set up correctly by connecting to it using the `psql` client:

```bash
export POSTGRES_PASSWORD=$(kubectl get secret --namespace database psql-test-postgresql -o jsonpath="{.data.postgres-password}" | base64 -d)

echo $POSTGRES_PASSWORD
PGDeiuKIDd

kubectl run psql-test-postgresql-client-2 \
    --rm \
    --tty \
    -i \
    --restart='Never' \
    --namespace database \
    --image docker.io/bitnami/postgresql:16.4.0-debian-12-r0 \
    --env="PGPASSWORD=$POSTGRES_PASSWORD" \
    --command \ 
    -- psql --host psql-test-postgresql -U postgres -d postgres -p 5432
```
