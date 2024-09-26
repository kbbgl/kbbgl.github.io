---
slug: quay-io-tutorial
title: quay.io Tutorial
authors: [kbbgl]
tags: [docker, registry, images]
---

### Log in:

```bash
sudo docker login quay.io
Username: me
Password: T*********0!
```

### Create container:

```bash
sudo docker run busybox echo "fun"

Unable to find image 'busybox:latest' locally
latest: Pulling from library/busybox
9c075fe2c773: Pulling fs layer
9c075fe2c773: Download complete
9c075fe2c773: Pull complete
Digest: sha256:c3dbcbbf6261c620d133312aee9e858b45e1b686efbcead7b34d9aae58a37378
Status: Downloaded newer image for busybox:latest
```

### List container to get container id:

```bash
sudo docker ps -l

CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
5fad7d83f69a        busybox             "echo fun"          7 seconds ago       Exited (0) 6 seconds ago                       nervous_nobel
```

### Create new image

Once a container has terminated in Docker, the next step is to commit the container to an image, and then tag that image with a relevant name so it can be saved to a repository.

Docker lets us do this in one step with the commit command. To do so, we run the docker commit with the container ID from the previous step and tag it to be a repository under quay.io.

```bash
sudo docker commit 5fad7d83f69a quay.io/kobbigal/custom_repos
```

### Push the image to Quay
Now that we've tagged our image with a repository name, we can push the repository to Quay Container Registry :

```bash
sudo docker push quay.io/kobbigal/custom_repos

The push refers to a repository [quay.io/kobbigal/custom_repos] (len: 1)
Sending image list
Pushing repository quay.io/kobbigal/custom_repos (1 tags)
```
