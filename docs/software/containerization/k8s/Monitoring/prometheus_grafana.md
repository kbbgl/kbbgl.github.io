---
title: How to Deploy Prometheus and Grafana in Kubernetes
slug: how-to-deploy-prometheus-grafana-k8s
app2or: kgal-akl
tags: [devops, monitoring, kubernetes, grafana, prometheus, metrics, k8s, helm]
---

```bash
kubectl create namespace monitoring
```

```bash
helm repo add prometheus-community \ 
	https://prometheus-community.github.io/helm-charts
helm repo update
helm install prometheus prometheus-community/prometheus --namespace monitoring
```

```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm install grafana grafana/grafana \
	--namespace monitoring \
	--set adminPassword="admin" \
	--set service.type=LoadBalancer \
	--set service.port=30080
```


## Add Prometheus Data Source to Grafana

```bash
kubectl port-forward -n monitoring grafana 3000:3000

After logging into Grafana, go to `http:///localhost:3000/connections/datasources` > Add New Data Source > Prometheus. 

In the Prometheus server URL enter `http://prometheus-server.default.svc.cluster.local`.