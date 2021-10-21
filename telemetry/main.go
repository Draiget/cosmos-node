package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const namespace = "gaiad"
const endpointStatus = "status"
const endpointNetwork = "net_info"

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}

	listenAddress = flag.String("metrics.listen-address", ":9560",
		"Address to listen on for telemetry")
	metricsPath = flag.String("metrics.path", "/metrics",
		"Path under which to expose metrics")

	// Metrics
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last Gaiad node query successful.",
		nil, nil,
	)

	statCurrentBlock = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "current_block"),
		"Current Gaiad node block.",
		nil, nil,
	)
	statTimeDeviation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "time_deviation"),
		"Current Gaiad time deviation (current time - block time) in seconds.",
		nil, nil,
	)
	statPeerCount = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "peer_count"),
		"Current amount of peers this Gaiad node connected to.",
		nil, nil,
	)
	statSyncCatchingUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "sync_catching_up"),
		"Is synchronization in progress.",
		nil, nil,
	)
)

func main(){
	flag.Parse()
	gaiadEndpoint := os.Getenv("GAIAD_ENDPOINT")

	r := prometheus.NewRegistry()
	exporter := NewExporter(gaiadEndpoint)
	r.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}