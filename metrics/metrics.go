package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	DbQueryCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "evg",
			Name:      "db_query_count",
			Help:      "counts number of calls of db.Query()",
		})
	AllocatorHostRatio = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "evg",
			Name:      "allocator_host_ratio",
			Help:      "counts host allocator hostQueueRatio",
		})

	FifteenSecondError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "evg",
			Name:      "fifteen_second_error",
			Help:      "counts total # of errors in fifteen second units",
		})
)

func RegisterMetrics() {
	prometheus.MustRegister(DbQueryCount)
	prometheus.MustRegister(AllocatorHostRatio)
	prometheus.MustRegister(FifteenSecondError)
}
