package metrics

import (
	"context"
	"github.com/ozoncp/ocp-instruction-api/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
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

func Run(ctx context.Context, wg *sync.WaitGroup) {
	Register()

	s := http.Server{Addr: config.Data.Metrics_Listen}

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Error().Msgf("failed to serve metrics: %v", err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			log.Debug().Msg("shutdown metrics")
			if err := s.Shutdown(ctx); err != nil && err != context.Canceled {
				log.Error().Msgf("failed to shutdown metrics: %v", err)
			}
		}
	}()
}
