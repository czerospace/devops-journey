package main

import (
	"fmt"
	"net/http"
)

// 徒手开发 exporter 简单的 http 实现返回 Prometheus 数据格式

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "lexporter_request_count{user=\"admin\"} 1000")
}

func main() {
	http.HandleFunc("/metrics", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
