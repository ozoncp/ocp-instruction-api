package metrics

import (
	"github.com/ozoncp/ocp-instruction-api/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
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

func Run() {
	go func() {
		Register()

		http.Handle("/metrics", promhttp.Handler())

		err := http.ListenAndServe(config.Data.Metrics_Listen, nil)
		if err != nil {
			log.Fatal().Msgf("failed to serve metrics: %v", err)
		}
	}()
}
