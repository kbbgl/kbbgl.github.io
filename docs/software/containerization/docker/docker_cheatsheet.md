---
slug: cheatsheet
title: Docker Cheatsheet
authors: [kbbgl]
tags: [docker, cheatsheet]
---


## Create `Dockerfile`

```dockerfile
# the first stage of our build will use a maven 3.6.1 parent image
FROM maven:3.6.1-jdk-8-alpine AS MAVEN_BUILD

# copy the pom and src code to the container
COPY ./ ./

# package our application code
RUN mvn clean package

# the second stage of our build will use open jdk 8 on alpine 3.9
FROM openjdk:8-jre-alpine3.9

# copy only the artifacts we need from the first stage and discard the rest
COPY --from=MAVEN_BUILD /docker-multi-stage-build-demo/target/demo-0.0.1-SNAPSHOT.jar /demo.jar

# set the startup command to execute the jar
CMD ["java", "-jar", "/demo.jar"]
```

## Build container

```bash
docker build -t $TAG .
```

## Run container

```bash
docker run -d -p 8080:8080 anna/docker-multi-stage-build-demo:1.0-SNAPSHOT
```

## Get system information

```bash
sudo docker info

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.12
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 481103c8793316c118d9f795cde18060847c370e
 runc version: 1.0.0~rc10+dfsg2-1
 init version:
 Security Options:
  seccomp
   Profile: default
 Kernel Version: 4.19.118-Re4son-v7l+
 Operating System: Kali GNU/Linux Rolling
 OSType: linux
 Architecture: armv7l
 CPUs: 4
 Total Memory: 7.78GiB
 Name: kali
 ID: 7FOC:BFDK:S6RB:HTJ2:U6UH:DQRU:RWTC:UCKI:YCOA:XW5O:J4P2:DZAN
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: No memory limit support
WARNING: No swap limit support
WARNING: No kernel memory limit support
WARNING: No kernel memory TCP limit support
WARNING: No oom kill disable support
WARNING: No cpu cfs quota support
WARNING: No cpu cfs period support
```

### Optimize the Docker Build cache

Adding `RUN --mount=type=cache` keeps your package cache intact between builds. No more re-downloading the entire internet every time you build your image. It’s especially handy when you’re working with large dependencies. Implement this, and watch your build efficiency go through the roof.

```dockerfile
# Use an official Node base image
FROM node:14
 
# Install dependencies first to leverage Docker cache
COPY package.json package-lock.json ./
 
# Using cache mount for npm install, so unchanged packages aren’t downloaded every time
RUN --mount=type=cache,target=/root/.npm \
    npm install
 
# Copy the rest of your app's source code
COPY . .
 
# Your app's start command
CMD ["npm", "start"]
```
