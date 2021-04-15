package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	dbQueryCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "evg",
			Name:      "db_query_count",
			Help:      "counts number of calls of db.Query()",
		})
	allocatorHostRatio = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "evg",
			Name:      "allocator_host_ratio",
			Help:      "counts host allocator hostQueueRatio",
		})

	fifteenSecondCrons = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "evg",
			Name:      "fifteen_second_crons",
			Help:      "counts total # of crons in fifteen second units",
		})
)

func IncDbQueryCount() {
	dbQueryCount.Inc()
}

func SetAllocatorHostRatio(val float64) {
	allocatorHostRatio.Set(val)
}

func IncFifteenSecondCrons() {
	fifteenSecondCrons.Inc()
}
