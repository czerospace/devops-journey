package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 使用 SDK 开发 exporter

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8050", nil)
}
