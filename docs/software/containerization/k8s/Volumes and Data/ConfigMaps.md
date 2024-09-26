## `ConfigMap`s

A similar API resource to `Secret`s is the `ConfigMap`, except the data is not encoded. In keeping with the concept of decoupling in Kubernetes, using a `ConfigMap` decouples a container image from configuration artifacts. 

They store data as sets of key-value pairs or plain configuration files in any format. The data can come from a collection of files or all files in a directory. It can also be populated from a literal value. 

A `ConfigMap` can be used in several different ways. A container can use the data as environmental variables from one or more sources. The values contained inside can be passed to commands inside the `Pod`. A Volume or a file in a Volume can be created, including different names and particular access modes. In addition, cluster components like controllers can use the data.

Let's say you have a file on your local filesystem called `config.js`. You can create a `ConfigMap` that contains this file. The configmap object will have a data section containing the content of the file:

```bash
kubectl get configmap foobar -o yaml
```

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
    name: foobar
data:
    config.js: |
         {
...
```

`ConfigMap`s can be consumed in various ways:

- Pod environmental variables from single or multiple `ConfigMap`s
- Use `ConfigMap` values in Pod commands
- Populate `Volume` from `ConfigMap`
- Add `ConfigMap` data to specific path in Volume
- Set file names and access mode in Volume from `ConfigMap` data
- Can be used by system components and controllers.

### Using `ConfigMap`

`ConfigMap`s can be consumed similar to `Secret`s by using environmental variables or mounting volumes.

To use as environmental variables, the `Pod` manifest will use `valueFrom` key and `configMapKeyRef` value to read the values:

```yaml
env:
- name: SPECIAL_LEVEL_KEY
  valueFrom:
    configMapKeyRef:
      name: special-config
      key: special.how
```

To use with volumes:

```yaml
volumes:
  - name: config-volume
    configMap:
      name: special-config
```

