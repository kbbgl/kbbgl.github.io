# Debug NodeJS in Chrome

## On server running Kubernetes cluster

1. Set up inspection in Node container:

```bash
kubectl edit deployment api-gateway
```

```yaml
Containers:
 api-gateway:
     Command:
         ...
            node --inspect-brk=0.0.0.0:9229 app.js;
```

1. Forward container debugging port (`9229`) to Kubernetes node:

```bash
kubectl port-forward $POD_NAME --address=0.0.0.0 9229:9229
```

## On local machine

1. Create ssh session to remote machine that will forward connection to port `9221`:

```bash
ssh -L 9221:localhost:9229 user@host
```

1. Open Chrome and visit:

```text
chrome://inspect
```

Click on Configure and add:

```text
localhost:9221
```

To Target discovery settings.

1. Under Remote Target #LOCALHOST you should see app.js.
1. Go to Sources tab, Node and there will be the tree of the project.
1. Set debugger.
