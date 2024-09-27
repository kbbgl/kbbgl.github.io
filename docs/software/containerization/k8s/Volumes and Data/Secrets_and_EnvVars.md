## `Secret`s and Environmental Variables

Secrets can be specified as environmental variables.
They must exist before being requested.

An example of a `Secret` as an env var:

```yaml
# secret
...
spec:
  container: 
  - image: mysql:5.5
    env:
    - name: MYSQL_ROOT_PASSWORD
      valueFrom:
        name:
          secretKeyRef:
            name: mysql
            key: password
...
```
