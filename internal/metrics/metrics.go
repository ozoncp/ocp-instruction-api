package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	opsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ops_count",
			Help: "Count of operations",
		},
		[]string{"operation"},
	)
)

func Register() {
	prometheus.MustRegister(opsCounter)
}

func OpsCounter_Inc(operation string) {
	opsCounter.With(prometheus.Labels{"operation": operation}).Inc()
}
