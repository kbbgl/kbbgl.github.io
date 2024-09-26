## Deployment Configuration

### Details

```yaml
apiVersion: v1         # not the deployment, but k8s API
items:                 # list of items
- apiVersion: apps/v1  
  kind: Deployment
```

### Configuration Metadata

```yaml 
metadata: 
  annotations:
    deployment.kubernetes.io/revision: "1"
  creationTimestamp: 2017-12-21T13:57:07Z
  generation: 1      # how many times object has been edited
  labels:            
    app: dev-web
  name: dev-web
  namespace: default
  resourceVersion: "774003" # tied to etcd to help with concurrent objects
  selfLink: /apis/apps/v1/namespaces/default/deployments/dev-web # how kube-apiserver will ingest information to API
  uid: d52d3a63-e656-11e7-9319-42010a800003 # unique ID for the lifecycle of object
```

### Configuration Spec

```yaml
spec:  
  progressDeadlineSeconds: 600  # time until a progress error is reported during a change (e.g. quotas, image issues, limit ranges)
  replicas: 1  # how many Pods will be created
  revisionHistoryLimit: 10   # how many old ReplicaSet specs to retain for rollback
  selector:     
    matchLabels:      # used to match labels (AND)
      app: dev-web  
  strategy:           # control Pod update
    rollingUpdate:    # how many Pods are deleted at a time with conditions below
      maxSurge: 25%   # max desired Pods to create. Creates a certain number of Pods before deleting old ones.
      maxUnavailable: 25%     # max pods that can be in non-ready state
    type: RollingUpdate # type of strategy chosen
```

### Configuration Pod Template

```yaml
template: #data passed to replicaset to determine how to deploy object
  metadata: 
  creationTimestamp: null
    labels:
      app: dev-web
  spec:
    containers:
    - image: nginx:1.13.7-alpine
      imagePullPolicy: IfNotPresent
      name: dev-web
      resources: {}
      terminationMessagePath: /dev/termination-log
      terminationMessagePolicy: File
    dnsPolicy: ClusterFirst # DNS query should go to coreDNS (dnsPolicy: ClusterFirst) or node dns server (dnsPolicy: Default)
    restartPolicy: Always # should container always restart if killed
    schedulerName: default-scheduler # allows to configure different scheduler
    securityContext: {} # setting security such as SELinux, AppArmor, etc.
    terminationGracePeriodSeconds: 30 # amount of time to wait for SIGTERM to run until a SIGKILL is used to terminate container.
```

### Configuration Status

Generated when information is requested (`kubectl get deployment -o yaml`)

```yaml
status:
  availableReplicas: 2 # how many replicas were configured by ReplicaSet
  conditions:
  - lastTransitionTime: 2017-12-21T13:57:07Z
    lastUpdateTime: 2017-12-21T13:57:07Z
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  observedGeneration: 2 # how often deployment has been updated
  readyReplicas: 2
  replicas: 2
  updatedReplicas: 2
```