package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ybbus/jsonrpc/v2"
	"log"
	"strconv"
	"time"
)

type Exporter struct {
	gaiadEndpoint string
}

func NewExporter(endpoint string) *Exporter {
	return &Exporter{
		gaiadEndpoint: endpoint,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- statCurrentBlock
	ch <- statTimeDeviation
	ch <- statPeerCount
	ch <- statSyncCatchingUp
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	nodeStatus, netStatus, err := e.LoadStatus()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		log.Println(err)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)

	e.HitGaiadRpcAndUpdateMetrics(nodeStatus, netStatus, ch)
}


func (e *Exporter) LoadStatus() (*MetricStatus, *MetricNetwork, error) {
	rpcClient := jsonrpc.NewClient(e.gaiadEndpoint)

	var status *MetricStatus
	err := rpcClient.CallFor(&status, endpointStatus)
	if err != nil {
		return nil, nil, err
	}

	var network *MetricNetwork
	err = rpcClient.CallFor(&network, endpointNetwork)
	if err != nil {
		return status, nil, err
	}

	return status, network, err
}

func (e *Exporter) HitGaiadRpcAndUpdateMetrics(status *MetricStatus, network *MetricNetwork, ch chan<- prometheus.Metric) {
	timeDeviation := time.Now().Unix() - status.SyncInfo.LatestBlockTime.int64

	currentBlock, _ := strconv.ParseFloat(status.SyncInfo.LatestBlockHeight, 32)
	ch <- prometheus.MustNewConstMetric(
		statCurrentBlock, prometheus.GaugeValue, float64(int(currentBlock)),
	)

	ch <- prometheus.MustNewConstMetric(
		statTimeDeviation, prometheus.GaugeValue, float64(int(timeDeviation)),
	)

	peerCount, _ := strconv.ParseFloat(network.NumberOfPeers, 64)
	ch <- prometheus.MustNewConstMetric(
		statPeerCount, prometheus.GaugeValue, peerCount,
	)

	var catching float64 = 0
	if status.SyncInfo.CatchingUp {
		catching = 1
	}

	ch <- prometheus.MustNewConstMetric(
		statSyncCatchingUp, prometheus.GaugeValue, catching,
	)
}