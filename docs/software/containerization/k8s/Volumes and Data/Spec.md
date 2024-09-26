## Volume Spec

One of the may types of storage available is an `emptyDir`. The `kubelet` will create the directory in the container but not mount any storage. Any data created is written to the shared container space so it will not be persistent. When the `Pod` is deleted, the directory is deleted along with the container.

Creating a `Pod` with the following specification would create a container with a volume named `scratch-volume` with a directory `/scratch` inside the container.
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: fordpinto
  namespace: default
spec:
  containers:
  - image: simpleapp
    name: gastank
    command:
      - sleep
      - "3600"
    volumeMounts:
    - mountPath: /scratch
      name: scratch-volume
    volumes:
    - name: scratch-volume
      emptyDir: {}
```

### Shared Volume Example

The following specification creates a `Pod` with 2 containers that share access to a volume:

```yaml
...
containers:
- name: alphacont
  image: busybox
  volumeMounts:
  - mountPath: /alphadir
    name: sharevol
- name: betacont
  image: busybox
  volumeMounts:
  - mountPath: /betadir
    name: sharevol
volumes:
  - name: sharevol
    emptyDir: {}
```

