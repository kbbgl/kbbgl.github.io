# Minikube Cheatsheet

```bash
minikube status
```

```bash
minikube stop
```

```bash
minikube start --driver=virtualbox
```

minikube start returned an error:

```text
machine does not exist
```

then you need to clear minikube's local state:

```bash
minikube delete
```
