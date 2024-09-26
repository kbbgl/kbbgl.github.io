# Prometheus Architecture

## Prometheus Server

Has three components:

- **Time Series Database**: a database with all resource metrics (CPU/RAM/Disk/Network)
- **Data Retrieval Worker** - pulls metrics data from services and stores them in database.
- **HTTP Server** - Accepts PromQL queries and returns results to clients (such as Grafana/Prometheus Web UI) from database.

## Data Collections from Targets

The Data Retrieval Worker pulls data using HTTP from each target.

Each target needs to expose an endpoint, by default it's `hostname/metrics`

For example, in Kubernetes deployment, we could see the metrics by checking what endpoint the Prometheus operator is listening:

```bash
kubectl get endpoint -n monitoring

app-prom-operator-prometheus-node-exporter  10.50.22.203:9100
```

```bash
curl --silent 10.50.22.203:9100/metrics

# HELP node_disk_io_time_seconds_total Total seconds spent doing I/Os.
# TYPE node_disk_io_time_seconds_total counter
node_disk_io_time_seconds_total{device="dm-0"} 216.868
node_disk_io_time_seconds_total{device="dm-1"} 1174.188
node_disk_io_time_seconds_total{device="dm-2"} 0.048
node_disk_io_time_seconds_total{device="sda"} 1192.516
node_disk_io_time_seconds_total{device="sr0"} 0
```

The target must supply the information in the correct format.

## Target Endpoints and Exporters

Exporters fetch metrics from a service and converts them into the correct format.
