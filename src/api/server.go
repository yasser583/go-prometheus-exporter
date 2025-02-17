package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	metricsSuffix = "/metrics"
	listenAddress = "localhost:9453"
)

var (
	apiEndpoints = [2]string{"/api/efectivo", "/api/tarjeta"}
)

func main() {
	// register custom exporter
	metricsPrefix := "myexporter"
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	// handle API endpoints
	for _, e := range apiEndpoints {
		http.HandleFunc(e, func(w http.ResponseWriter, r *http.Request) {
			exporter.IncreCounter()
			exporter.IncreCounterWithEndpoint(r.URL.Path)
		})
	}

	// Handle Metrics endpoint
	http.Handle(metricsSuffix, prometheus.Handler())

	log.SetLevel(log.DebugLevel)
	log.Debugf("Metrics exported at http://%s%s", listenAddress, metricsSuffix)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
