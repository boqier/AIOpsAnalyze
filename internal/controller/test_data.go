// test_data.go
package controller

import (
	"time"
)

// 模拟的测试数据结构
func GetTestData() TestData {
	// 目标Pod信息
	podName := "my-app-54f7596d5b-8z2xg"
	namespace := "default"

	return TestData{
		// Prometheus CPU指标
		PrometheusCPUMetrics: `
# HELP container_cpu_usage_seconds_total Cumulative cpu time consumed per cpu in seconds.
# TYPE container_cpu_usage_seconds_total counter
container_cpu_usage_seconds_total{container="my-app",namespace="default",pod="my-app-54f7596d5b-8z2xg"} 1234.56 1737358874000
container_cpu_usage_seconds_total{container="my-app",namespace="default",pod="my-app-54f7596d5b-xv3pz"} 567.89 1737358874000
container_cpu_usage_seconds_total{container="redis",namespace="default",pod="redis-67559f8877-v2fr9"} 123.45 1737358874000
`,

		// Prometheus 内存指标
		PrometheusMemoryMetrics: `
# HELP container_memory_usage_bytes Current memory usage in bytes, including all memory regardless of when it was accessed.
# TYPE container_memory_usage_bytes gauge
container_memory_usage_bytes{container="my-app",namespace="default",pod="my-app-54f7596d5b-8z2xg"} 536870912 1737358874000
container_memory_usage_bytes{container="my-app",namespace="default",pod="my-app-54f7596d5b-xv3pz"} 268435456 1737358874000
container_memory_usage_bytes{container="redis",namespace="default",pod="redis-67559f8877-v2fr9"} 134217728 1737358874000
`,

		// Prometheus 重启次数指标
		PrometheusRestartMetrics: `
# HELP kube_pod_container_status_restarts_total The number of container restarts.
# TYPE kube_pod_container_status_restarts_total counter
kube_pod_container_status_restarts_total{container="my-app",namespace="default",pod="my-app-54f7596d5b-8z2xg"} 3 1737358874000
kube_pod_container_status_restarts_total{container="my-app",namespace="default",pod="my-app-54f7596d5b-xv3pz"} 0 1737358874000
`,

		// Prometheus AlertManager 告警
		PrometheusAlerts: `
[{"status":"firing","labels":{"alertname":"HighCPUUsage","container":"my-app","namespace":"default","pod":"my-app-54f7596d5b-8z2xg"},"annotations":{"summary":"High CPU usage detected","description":"Pod my-app-54f7596d5b-8z2xg has been using high CPU for more than 5 minutes"},"startsAt":"2024-12-20T10:00:00Z"},
{"status":"firing","labels":{"alertname":"HighMemoryUsage","container":"my-app","namespace":"default","pod":"my-app-54f7596d5b-8z2xg"},"annotations":{"summary":"High memory usage detected","description":"Pod my-app-54f7596d5b-8z2xg has been using high memory for more than 5 minutes"},"startsAt":"2024-12-20T10:05:00Z"},
{"status":"firing","labels":{"alertname":"ContainerRestart","container":"my-app","namespace":"default","pod":"my-app-54f7596d5b-8z2xg"},"annotations":{"summary":"Container restart detected","description":"Pod my-app-54f7596d5b-8z2xg has restarted 3 times in the last hour"},"startsAt":"2024-12-20T10:10:00Z"},
{"status":"firing","labels":{"alertname":"HighCPUUsage","container":"redis","namespace":"default","pod":"redis-67559f8877-v2fr9"},"annotations":{"summary":"High CPU usage detected","description":"Pod redis-67559f8877-v2fr9 has been using high CPU for more than 5 minutes"},"startsAt":"2024-12-20T10:00:00Z"}]
`,

		// Loki 日志
		LokiLogs: `
{"streams":[{"stream":{"container":"my-app","namespace":"default","pod":"my-app-54f7596d5b-8z2xg"},"values":[
["1737358800000000000","ERROR 2024-12-20T10:00:00Z [main] Failed to connect to database: Connection refused"],
["1737358801000000000","ERROR 2024-12-20T10:00:01Z [main] Retry database connection..."],
["1737358802000000000","ERROR 2024-12-20T10:00:02Z [main] Failed to connect to database: Connection refused"],
["1737358806000000000","WARN 2024-12-20T10:00:06Z [main] Memory usage is high: 90%"],
["1737358810000000000","ERROR 2024-12-20T10:00:10Z [main] CPU usage is high: 95%"],
["1737358815000000000","INFO 2024-12-20T10:00:15Z [main] Application restarting..."]
]},{"stream":{"container":"my-app","namespace":"default","pod":"my-app-54f7596d5b-xv3pz"},"values":[
["1737358800000000000","INFO 2024-12-20T10:00:00Z [main] Application started successfully"]
]}]}
`,

		// 目标Pod信息
		TargetPodName:   podName,
		TargetNamespace: namespace,
	}
}

// 测试数据结构体
type TestData struct {
	PrometheusCPUMetrics     string
	PrometheusMemoryMetrics  string
	PrometheusRestartMetrics string
	PrometheusAlerts         string
	LokiLogs                 string
	TargetPodName            string
	TargetNamespace          string
}

// 筛选Pod相关告警的函数
func FilterPodAlerts(allAlerts string, podName, namespace string) []Alert {
	// 这里应该是解析JSON并筛选的逻辑
	// 为了简化，这里只是模拟返回
	return []Alert{
		{
			Name:        "HighCPUUsage",
			Status:      "firing",
			PodName:     podName,
			Namespace:   namespace,
			Description: "Pod has been using high CPU for more than 5 minutes",
			StartTime:   time.Now().Add(-10 * time.Minute),
		},
		{
			Name:        "HighMemoryUsage",
			Status:      "firing",
			PodName:     podName,
			Namespace:   namespace,
			Description: "Pod has been using high memory for more than 5 minutes",
			StartTime:   time.Now().Add(-5 * time.Minute),
		},
	}
}

// 告警结构体
type Alert struct {
	Name        string
	Status      string
	PodName     string
	Namespace   string
	Description string
	StartTime   time.Time
}
