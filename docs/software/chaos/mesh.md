---
title: How to Set Up Chaos Mesh on Kubernetes using Helm
slug: how-to-setup-chaos-mesh-k8s-helm
app2or: kgal-akl
tags: [devops, chaos, helm, kubernetes, k8s, testing]
---

## Install 

```bash
helm repo add chaos-mesh https://charts.chaos-mesh.org
kubectl create ns chaos-mesh
```

When using k3s, we deploy it using:

```bash
helm install chaos-mesh chaos-mesh/chaos-mesh \
	--namespace chaos-mesh \
	--set chaosDaemon.runtime=containerd \
	--set chaosDaemon.socketPath=/run/k3s/containerd/containerd.sock \
	--set controllerManager.leaderElection.enabled=false \
	--set dashboard.securityMode=false \
	--version 2.8.0
```

For other container runtimes, check out [this](https://chaos-mesh.org/docs/production-installation-using-helm/#step-4-install-chaos-mesh-in-different-environments). To check what runtime is currently set up:

```bash
kubectl get node kgal-azure-dev \
	--output yaml | \
	yq '.status.nodeInfo.containerRuntimeVersion'
containerd://2.0.5-k3s1.32
```


## Run DNS Fault Experiment

Define an experiment. Make sure to modify the namespaces and the target pods in `.spec.selector.namespaces` and `.spec.selector.pods.$NAMESPACE[0].$TARGET_POD_NAME:

```yaml
kind: DNSChaos
apiVersion: chaos-mesh.org/v1alpha1
metadata:
  namespace: kbbgl-staging
  name: saas-outage
spec:
  action: error
  selector:
	namespaces:
	- kbbgl-staging
	pods:
	  kbbgl-staging:
	  - app-svc-app-6b98bd4645-s77qp
	mode: all
	patterns:
	- app1.kbbgl.svc.local
	- app1-ro.kbbgl.svc.local
	- app2.kbbgl.svc.local
	- app2-ro.kbbgl.svc.local
	- app3.kbbgl.svc.local
	- app3-ro.kbbgl.svc.local
	- app4.kbbgl.svc.local
	- app5.kbbgl.svc.local
	- ms1.kbbgl.svc.local
	- ms1-ro.kbbgl.svc.local
	- ms2.kbbgl.svc.local
	- ms2-ro.kbbgl.svc.local
	- ms3.kbbgl.svc.local
	- ms3-ro.kbbgl.svc.local
	- ms4.kbbgl.svc.local
	- ms4-ro.kbbgl.svc.local
	- api.kbbgl.svc.local
	- rest.kbbgl.svc.local
	- log.kbbgl.svc.local
```

Then apply it:

```bash
kubectl apply -f dns-fault-kbbgl-staging.yaml
```

## Check Experiment Status

```bash
kubectl get dnschaos -n kbbgl-staging

NAME STATUS AGE 
saas-outage AllInjected 15s
```

```bash
kubectl describe dnschaos saas-outage -n kbbgl-staging
```
## Stop Experiment

```bash
kubectl delete -f dns-fault-kbbgl-staging.yaml
```