# Create Jaegar Instance

1. Create necessary resources

```bash
kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/crds/jaegertracing.io_jaegers_crd.yaml
kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/service_account.yaml
kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/role.yaml
kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/role_binding.yaml
kubectl create -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/operator.yaml
```

1. Download and modify `cluster_role_binding.yaml` to change the namespace:

```bash
kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role.yaml
kubectl create -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/master/deploy/cluster_role_binding.yaml
```

1. Create example instance:

```yaml
#jaegar_demo.yaml
apiVersion: jaegertracing.io/v1
kind: Jaeger
metadata:
  name: jaeger_demo
```

```bash
kubectl apply -f jaeger_demo.yaml
```

1. Expose app (takes a few minutes to initialize):

```bash
kubectl -n app port-forward jaegar_demo 16686:16686 --address 0.0.0.0
```
