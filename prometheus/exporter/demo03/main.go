package main

import (
	"exporter/demo03/collector"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 自定义 采集器 和 注册表

func main() {
	// 创建一个自定义的注册表
	registry := prometheus.NewRegistry()

	// 可选: 添加 process 和  Go 运行时指标到 自定义的注册表中
	// registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	// registry.MustRegister(prometheus.NewGoCollector())

	// 注册自定义采集器
	registry.MustRegister(collector.NewmemCollector())
	//  暴露指标
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
	http.ListenAndServe(":8050", nil)
}
