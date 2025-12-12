---
slug: how-to-setup-gitlab-self-hosted-runner
title: How to Set Up GitLab Self-Hosted Runner
description: Steps to set up GitLab Self-Hosted Runner
authors: [kgal-akl]
tags: [cicd, ci, gitlab, self-hosted]
---


## Register Runner

```bash
gitlab-runner register \
	--url https://gitlab.com \
	--token glrt-$TOKEN \
	--tag-list macos
```

Token can be by creating one in [Personal access tokens](https://docs.gitlab.com/user/profile/personal_access_tokens). When we register a runner, we will be asked to add tags to the runner. Tags control what runner will actually run the pipeline. 

## Running GitLab Pipeline in Docker

We need to create a GitLab volume to persist the configuration:

```bash
docker create volume gitlab-runner-config
```

We then run `gitlab-runner` daemon in Docker:
```bash
docker run -d --name gitlab-runner \
--restart always \
-v /var/run/docker.sock:/var/run/docker.sock \
-v gitlab-runner-config:/etc/gitlab-runner \
-v /Users/kgal/.gitlab-runner/config.toml:/etc/gitlab-runner/config.toml \
gitlab/gitlab-runner:latest;
```

We then trigger a pipeline in GitLab. In the pipeline defined in `gitlab-ci.yaml`, if we register the `macos` tag to the runner and then later add the same tag to `get-resource-job.tags[0]` as below:
```yaml
variables:
  ACCOUNT_ID: 123456
  FQDN: http://some.domain.dev
  OUT: /tmp/res

stages:
  - get_resource

get-resource-job:
  stage: get_resource
  image: alpine/ci:0.0.1
  tags:
    - macos
  before_script:
    - echo "getting resource..."
    - curl "$FQDN/get/resource" -H "x-acc-id: $ACCOUNT_ID" > $OUT
  script:
    - echo "done getting resource, saved in '$OUT'"
```

We should see the following output where we can see that the job was trigger in our local Docker:

```bash
docker logs -f gitlab-runner
Runtime platform                                    arch=arm64 os=linux pid=7 revision=4d7093e1 version=18.0.2
Starting multi-runner from /etc/gitlab-runner/config.toml...  builds=0 max_builds=0
Running in system-mode.

Usage logger disabled                               builds=0 max_builds=1
Configuration loaded                                builds=0 max_builds=1
listen_address not defined, metrics & debug endpoints disabled  builds=0 max_builds=1
[session_server].listen_address not defined, session endpoints disabled  builds=0 max_builds=1
Initializing executor providers                     builds=0 max_builds=1
Checking for jobs... received                       job=10227995527 repo_url=https://gitlab.com/kbbgl/demo.git runner=9mUX02bq0
Added job to processing list                        builds=1 job=10227995527 max_builds=1 project=70251116 repo_url=https://gitlab.com/kbbgl/demo.git time_in_queue_seconds=0
Appending trace to coordinator...ok                 code=202 job=10227995527 job-log=0-2859 job-status=running runner=9mUX02bq0 sent-log=0-2858 status=202 Accepted update-interval=3s
Job succeeded                                       duration_s=5.574595877 job=10227995527 project=70251116 runner=9mUX02bq0
Appending trace to coordinator...ok                 code=202 job=10227995527 job-log=0-3559 job-status=running runner=9mUX02bq0 sent-log=2859-3558 status=202 Accepted update-interval=3s
Updating job...                                     bytesize=3559 checksum=crc32:1c8e0aee job=10227995527 runner=9mUX02bq0
Submitting job to coordinator...accepted, but not yet completed  bytesize=3559 checksum=crc32:1c8e0aee code=202 job=10227995527 job-status=running runner=9mUX02bq0 update-interval=1s
Updating job...                                     bytesize=3559 checksum=crc32:1c8e0aee job=10227995527 runner=9mUX02bq0
Submitting job to coordinator...ok                  bytesize=3559 checksum=crc32:1c8e0aee code=200 job=10227995527 job-status=success runner=9mUX02bq0 update-interval=0s
Removed job from processing list                    builds=0 job=10227995527 max_builds=1 project=70251116 repo_url=https://gitlab.com/kbbgl/demo.git time_in_queue_seconds=0
```

## Managing Runners

We can see what runners are registered:

```bash
gitlab-runner list
```

We can start/stop the service and check the status:
```bash
gitlab-runner stop
gitlab-runner start
gitlab-runner status
```
## Configuration

```
cat ~/.gitlab-runner/config.toml
```
