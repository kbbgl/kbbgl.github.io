## Working with `Job/CronJob`s

### `Job`
1) Create a `Job` spec:

```yaml
# job.yaml

apiVersion: batch/v1
kind: Job
metadata:
  name: sleepy
spec:
  completions: 5 # number of completions
  parallelism: 2 # number of jobs to run in parallel
  activeDeadlineSeconds: 15 # the job/pods will end when reaching 15 seconds
  template:
    spec:
      containers:
      - name: resting
        image: busybox
        command: ["/bin/sleep"]
        args: ["3"]
      restartPolicy: Never
```

2) Create `Job`:

```bash
kubectl create -f job.yaml
```

### `CronJob`

