package collector

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v3/mem"
)

type memCollector struct {
	myDesc      *prometheus.Desc
	labelValues []string
}

// 每个收集器都必须实现descripe函数
func (c *memCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.myDesc
}

// 采集指标
func (c *memCollector) Collect(ch chan<- prometheus.Metric) {
	Mem, _ := mem.VirtualMemory()
	fmt.Println(Mem.Total)
	ch <- prometheus.MustNewConstMetric(c.myDesc, prometheus.GaugeValue, float64(Mem.Total)/1024/1024/1024, c.labelValues...)
}

func NewmemCollector() *memCollector {
	return &memCollector{
		myDesc: prometheus.NewDesc(
			"mem_total",
			"mem total for linux server.",
			// 动态标签的 key 列表
			[]string{"instance_id", "instance_name"},
			// 静态标签
			prometheus.Labels{"module": "linux"},
		),
		// 动态标签的 value 列表，这里必须与声明的动态标签的 key 一一对应
		labelValues: []string{"server", "linux"},
	}
}
