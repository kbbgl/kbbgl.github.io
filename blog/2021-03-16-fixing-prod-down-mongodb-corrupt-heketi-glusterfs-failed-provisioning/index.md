---
slug: prod-down-mongodb-corrupt-heketi-glusterfs-provisioning
title: Fixing Production Down caused by MongoDB Corruption and Heketi/GlusterFS Failed Provisioning
description: Fill me up!
authors: [kbbgl]
tags: [cluster,gluster,glusterfs,heketi,kubectl,kubernetes,mongo,mongodb,production,troubleshooting]
---
## Introduction

Today I received an escalation from one of our largest and most strategic customers. Over the weekend, the customer had ‘patched’ their 3 Ubuntu 18.04 nodes running Kubernetes 1.17.  They were using [`glusterfs`](https://www.gluster.org/) as their shared storage class.
I was trying to figure out what this ‘patching’ job entailed so we could assess which step of the maintenance to focus on but could not get all the necessary details from the customer support team except for the following order of operations:

1. They ran `kubectl drain` to prepare their 3 nodes for the patching. This ensured that all `Pods` would get evicted, including all persistent storage and services.
1. They ran kernel updates to patch security vulnerabilities.
1. They upgraded Ubuntu packages using `apt update&&apt upgrade`.
1. They restarted all their servers and ran `kubectl uncordon`.

After they finished this maintenance job, their application was no longer loading.
They had not taken any snapshots of the system before running this ‘patch’ job so we had to figure out why the application was not available.

## Salvaging MongoDB

I invited the customer to join me on a conference call so I could take a look and gather preliminary information about why the application is not running.
What I noticed when running `kubectl get pods` was that all the services were stuck on `Init 0/1`, which indicated that the execution of the first `initContainer` was terminating with non-zero return codes. All `initContainer`s were failing their connection to MongoDB as we can see from the spec of one of the `Pods` which was stuck on `Init`:

```yaml
Init Containers:
  init-mongodb:
    Image:         busybox:1.30.1
    Command:
      sh
      -c
      until nc -zv mongodb-service.prod 27017; do echo waiting for mongodb; sleep 2; done;
```

It was pretty evident that the MongoDB service was the cause for the dependent application services inability to initialize because we saw that all 3 replica set member’s `Pod`s were in `Init: CrashLoopBackOff`` state.

We needed to understand which `initContainer` of MongoDB was crashing. We were able to figure out why using `kubectl describe pod mongodb-replicaset-0` and `kubectl logs mongodb-replicaset-0 --all-containers`.
When running `kubectl describe pod mongodb-replicaset-0` we saw that the failing `initContainer` was one called `bootstrap`:

```yaml
Init Containers:
bootstrap:
    Image:         mongo:3.6.8
    Command:
      /work-dir/peer-finder
    Args:
      -on-start=/init/on-start.sh
      -service=mongodb-replicaset
    Environment:
      REPLICA_SET:    rs0
      TIMEOUT:        900
```

This container is a simple peer finder daemon that is useful with `StatefulSet` and related use cases.
All it does is watch DNS for changes in the set of endpoints that are part of the governing service. It periodically looks up the SRV record of the DNS entry that corresponds to a Kubernetes Service which enumerates the set of peers for this the specified service. Not really helpful. So we needed to review the logs and find out what’s causing this container to fail.

When we ran `kubectl logs mongod-replicaset-0 --all-containers` we noticed the following error:

```text
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: unable to read root page from file:WiredTiger.wt: WT_ERROR: non-specific WiredTiger error
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: WiredTiger has failed to open its metadata
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: This may be due to the database files being encrypted, being from an older version or due to corruption on disk
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: You should confirm that you have opened the database with the correct options including all encryption and compression options
E -        WT_ERROR: non-specific WiredTiger error src\mongo\db\storage\wiredtiger\wiredtiger_kv_engine.cpp 397
I STORAGE  WT_ERROR: non-specific WiredTiger error, terminating
```

The log that immediately grabbed our attention was the following:

```text
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: This may be due to the database files being encrypted, being from an older version or due to corruption on disk
```

Although we did not have any incriminating evidence, we were pretty sure that the server patching likely caused some data disk corruption because of an unclean shutdown of the MongoDB service.
At this point we were facing a few of problems:

1. We could not create a mongodump since the MongoDB server was not running and the `initContainers` were in a `CrashLoopBackOff`.
1. We did not have the necessary tools on the production environment since we could not install any additional sidecar containers or utilities.
1. As mentioned earlier, we did not have a snapshot of the server in a working state or a `mongodump` to rely on so we needed to figure out a way to fix this corrupted state somehow.

Luckily, the environment did have a `Deployment` called `system-recovery` which allowed us access to the mounted storage points so we could access the MongoDB flat files. We could then take these flat files, compress them and transfer them to our lab environment to attempt recovery:

```bash
> kubectl scale deployment system-recovery --replicas=1
deployment.apps/system-recovery scaled
> kubectl cp system-recovery-4184bfa40da-sah4131:/mongodb0 /tmp/mongodb0
> tar czvf /tmp/mongodb0.tar.gz /tmp/mongodb0
```

In our lab, we attempted to load MongoDB with the `--repair` flag but saw that it was failing with the same error we saw earlier on the production environment:

```bash
> mongod --version
3.6.8
> mongod --dbpath=/tmp/mongodb0 --repair
E STORAGE  WiredTiger error, file:WiredTiger.wt, connection: This may be due to the database files being encrypted, being from an older version or due to corruption on disk
```

We found out, after further research, that [MongoDB version 4 had a more robust repairing mechanism of corrupted WiredTiger schemas](https://jira.mongodb.org/browse/SERVER-19815). So we installed MongoDB 4 in our lab environment and reran the same repair:

```bash
> mongod --version
4.2.2
> mongod --dbpath=/tmp/mongodb0 --repair
```

We received a message that the operation was successful!
So now that we had a repaired MongoDB, we wanted to create a `mongodump` and then attempt to load the database in the lab environment to check whether the dependent Pods will successfully load which would mean that the application would load as well.

## MongoDB Version Conflict

Earlier, When attempting to repair the MongoDB, we needed to run the MongoDB server with version 4+ because of the new repair mechanism offered. The actual data which we extracted from the production environment was from version 3.6.8. We didn’t think this would be a problem until we actually attempted to load the database in order to generate a `mongodump`:

```bash
> mongod --version
3.6.8
> mongod --dbpath=/tmp/mongodb0
...
Found an invalid featureCompatibilityVersion document (ERROR: BadValue: Invalid value for version, found 3.6, expected '4.2' or '4.0'. Contents of featureCompatibilityVersion document in admin.system.version: { _id: "featureCompatibilityVersion", version: "3.6" }. See http://dochub.mongodb.org/core/4.0-feature-compatibility.). If the current featureCompatibilityVersion is below 4.0, see the documentation on upgrading at 
```

We were now blocked because MongoDB server and the actual data were different versions. We could not run the command:

```javascript
db.setFeatureCompatibility({_id: version"3.6"})
```

as suggested by a few different StackOverflow articles and MongoDB official documentation since we would need to be able to run the database.

This was really bad news since we did not know where this feature flag was actually stored, whether it was stored in clear text and if it was accessible. We decided to run a quick search within all the flat files in the MongoDB directory. It was a shot in the dark but we didn’t really have any other options at this point.

```bash
> grep -rnwi "compatibility" ./mongodb0/*
WiredTiger.turtle:1:Compatibility version
```

We managed to find something interesting within the WiredTiger.turtle file which was clear text. The whole file contents looked like this:

```text
Compatibility version
major=4,minor=2
WiredTiger version string
WiredTiger 4.2.1: (September 15, 2019)
WiredTiger version
major=4,minor=2,patch=1
```

Since we were in dire straits, we decided it was worth to modify the versions in this file and see if we could work around the MongoDB server validation. We modified it to the following:

```text
Compatibility version
major=3,minor=6
WiredTiger version string
WiredTiger 3.6.8: (July 12, 2018)
WiredTiger version
major=3,minor=6,patch=8
```

Miraculously, when running MongoDB daemon again, we were able to load the database and the application! We created a dump file using `mongodump` and transferred it to the production environment.

## Fixing Unhealthy GlusterFS/Heketi

We reconnected to the production environment, transferred the repaired and generated `mongodump` and we needed to somehow restore it onto the shared storage. We saw three possible ways of performing this:

1. Modify the MongoDB `StateFulSet` specification by changing the executed commands to something like an infinite sleep so that the `bootstrap` `initContainer` would not terminate. This would allow us to copy the `mongodump` into the container and run `mongorestore` to load the repaired databased and ensure it’s replicated across all GlusterFS storage locations. This was the preferred option.
1. Replace all MongoDB flat files manually within each one of the GlusterFS shared storage locations. This would not be ideal since it would go against MongoDB replication best practices.
1. Removing all GlusterFS `PersistentVolumeClaims`, recreate them from scratch and then bind them to the MongoDB replica set members.

After reviewing all three options, we decided to go with the 3rd approach. It was decided as the best option after consulting with our DevOps because they mentioned that they have set up dynamic storage allocation upon removal of the GlusterFS `PersistenVolumeClaims` and restart of the MongoDB `StateFulSet` `Pod`s.
The first step would be to delete the `PersistentVolumeClaims`. So we went to work.

```bash
> kubectl delete pvc datadir-mongodb-replicaset-0
> kubectl delete pvc datadir-mongodb-replicaset-1
> kubectl delete pvc datadir-mongodb-replicaset-2
```

But we noticed that these pvcs were stuck on Terminating. We even attempted to force delete them but they were still stuck:

```bash
> kubectl delete pvc datadir-mongodb-replicaset-0 --force --grace-period 0
```

After some research, we found out that these PersistentVolumeClaim specs had [`finalizers`](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#finalizers):

```bash
> kubectl describe pvc datadir-mongodb-replicaset-0

Name:          datadir-mongodb-replicaset-0
Namespace:     prod
StorageClass:  gluster
Status:        Bound
Volume:        pvc-64340467-8109-4856-9a49-2fc36563e9ab
Labels:        app=mongodb-replicaset
Annotations:   pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: kubernetes.io/glusterfs
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      150Gi
Access Modes:  RWO
VolumeMode:    Filesystem
Mounted By:    mongodb-replicaset-0
Events:        <none>
```

We edited the `PVC`s with:

```bash
kubectl edit pvc datadir-mongodb-replicaset-0
```

removed the `finalizer` section and reran the `kubectl delete pvc` commands. This resulted in a successful deletion. We then deleted the MongoDB `StateFulSet` to generate the `PVC`s using `kubectl delete pod mongodb-replicaset-{0,1,2}`. But to our surprise, the `PVC` and the `Pod`s were stuck `Pending`. Something was preventing the dynamic `PVC` allocation.

When we ran `kubectl describe pvc datadir-mongodb-replicaset-0` we saw that there was an `Event` with `ProvisioningFailed` with the following error message:

```bash
> kubectl describe pvc datadir-mongodb-replicaset-0
Failed to provision volume with StorageClass "gluster": failed to create volume.
```

The error message was far from informative so we needed to step back and review the state of `heketi` and `glusterfs`. We noticed that neither the `glusterfs` or the `heketi` `Pod`s had any logs from standard output and all of them, according to `kubectl get pods -n storage`, were in healthy, `Running` state.

We decided it would be best to use the `gluster` CLI and try to troubleshoot why the provisioning is failing. Luckily, the issue was pretty easy to find. We accessed the `glusterfs` Pod and ran the following command to check the status of `heketidbstorage`:

```bash
> kubectl exec glusterfs-k96d2 -it -- bash
> [ec2-user@node1 /] gluster volume status heketidbstorage
Status of volume: heketidbstorage
Gluster process TCP Port RDMA Port Online Pid
------------------------------------------------------------------------------
Brick node1:/var/lib/heketi/mounts/vg
_4a5d18544111232fc76cdc9872d340d6/brick_75c
ec7af846fbf2e770b9312c6bc56fe/brick 49152 0 N 198
Brick node2:/var/lib/heketi/mounts/vg
_3cf3e12449047bba9b1260d301914187/brick_3ea
27507e65df9ec9a4c0841d27962f9/brick 49153 0 N 198
Brick node3:/var/lib/heketi/mounts/vg
_78c5285b77bde0e7ed344df72ef5c630/brick_d54
3d1ea6e83036e7796a89b1594f4cc/brick 49152 0 N 186
Self-heal Daemon on localhost N/A N/A Y 177
Self-heal Daemon on node1 N/A N/A Y 183
Self-heal Daemon on node1 N/A N/A Y 189
```

We noticed that none of the bricks on `node{1,2,3}` were online so we decided to restart it:

```bash
gluster volume stop heketidbstorage 
gluster volume start heketidbstorage
```

This resulted in all bricks on all nodes getting back online, released the GlusterFS `PVC` provisioning and allowed all MongoDB replicaset member `Pod`s to initialize successfully!

## Restoring MongoDB and the Application

At this point, 3 intense and excruciating hours have passed, we had gone through a set a complex problems and we were all hoping that the last phase would result in a recovery of the MongoDB and the application.
We had the GlusterFS `PVC` provisioned, we had all 3 MongoDB replica set `Pod`s up. All that was left was to:

1. Copy the `mongodump` from the host machine to the master MongoDB replica set:

    ```bash
    # Find out who is the master

    > kubectl exec mongodb-replicaset-0 -c mongodb-replicaset -- mongo --eval='db.isMaster().primary' --quiet
    mongodb-replicaset-0.mongodb-replicaset.svc.cluster.local:27017

    # Copy the mongodump into the master Pod

    > kubectl cp /tmp/mongodb0/mongodump mongodb-replicaset-0 -c mongodb-replicaset:/tmp/
    ```

1. Run `mongorestore`:

    ```bash
    # Restore the database

    > kubectl exec mongodb-replicaset-0 -c mongodb-replicaset -- mongorestore -d prod --drop /tmp/mongodump/prod
    ```

1. Clear cache by deleting all `Pod`s in the application namespace:

    ```bash
    > kubectl delete pod --all -n prod
    ```

Voila! All Pods started successfully and the application was loading again!
