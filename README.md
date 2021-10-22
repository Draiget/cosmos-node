## Cosmos Hub Mainnet node
This repo contains various scripts and configuration steps to run cosmos mainnet node (Gaiad).

See more about [Cosmos](https://github.com/cosmos/mainnet).

## Dependencies

To get monitoring metrics from Cosmos node you need to run Prometheus.
See an example of Grafana dashboard in `grafana_dashboard.json`.
#### Prometheus

```shell
docker run \
    --name prometheus \
    -p 9090:9090 \
    -v ~/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus
```