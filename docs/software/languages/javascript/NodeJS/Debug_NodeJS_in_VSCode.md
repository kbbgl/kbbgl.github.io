# Debug NodeJS in VSCode

## On remote server where application is running

1. Set up inspection in Node container:

```bash
kubectl edit deployment api-gateway
```

```yaml
Containers:
 api-gateway:
     Command:
         ...
            node --inspect=0.0.0.0 app.js;
```

1. Clone the repository:

```bash
git clone https://gitlab.app.com/appTeam/Product/FE/vnext
```

1. Checkout release:

```bash
cd vnext
git checkout release/l8.2.4
```

1. Open folder in vscode:

```bash
code -a .
```

1. Create configuration in root folder:

```bash
touch launch.json
```

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "address": "10.233.97.254",
            "localRoot": "${workspaceFolder}",
            "name": "Attach to Remote",
            "port": 9229,
            "remoteRoot": "/usr/src/app",
            "request": "attach",
            "skipFiles": [
                "<node_internals>/**"
            ],
            "type": "pwa-node"
        },
    ]
}
```

Where `address` is the IP of the container:

```bash
kubectl get endpoints galaxy -o=jsonpath='{.subsets[0].addresses[0].ip}'
```

1. Run debugger task 'Attach to Remote'

1. Set breakpoint and debug.
