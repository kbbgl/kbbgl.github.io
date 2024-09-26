---
slug: gcloud-cp-files-from-to-remote
title: Copy Files from/to Remote
authors: [kbbgl]
tags: [cloud, google, gcloud, file_ops]
---


## Copy file to remote
```bash 
gcloud compute scp /path/in/local $GCP_INSTANCE_NAME:/path/in/remote --zone=$GCP_ZONE --tunnel-through-iap 
```

## Copy folder to remote
```bash
gcloud compute scp /path/in/local $GCP_INSTANCE_NAME:/path/in/remote --zone=$GCP_ZONE --tunnel-through-iap --recurse 
```

## Copy file from remote
```bash
gcloud compute scp $GCP_INSTANCE_NAME:/path/in/remote /path/in/local --zone=$GCP_ZONE --tunnel-through-iap 
```
