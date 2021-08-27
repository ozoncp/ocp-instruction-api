package main

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	cfg "github.com/ozoncp/ocp-instruction-api/internal/config"
	"github.com/ozoncp/ocp-instruction-api/internal/consumer"
	"github.com/ozoncp/ocp-instruction-api/internal/metrics"
	"github.com/ozoncp/ocp-instruction-api/internal/producer"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	"github.com/ozoncp/ocp-instruction-api/pkg/db"
	"github.com/uber/jaeger-client-go"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-instruction-api/internal/app/ocp-instruction-api"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"

	"github.com/rs/zerolog/log"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
)

func runTracing(ctx context.Context) {
	config := jaegercfg.Configuration{
		ServiceName: "ocp_instruction_api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Data.Jaeger_addr,
		},
	}

	logger := jaegerlog.StdLogger
	metricsFactory := jaegermetrics.NullFactory

	tracer, closer, err := config.NewTracer(
		jaegercfg.Logger(logger),
		jaegercfg.Metrics(metricsFactory),
	)

	if err != nil {
		log.Fatal().Msgf("trace init failed: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)

	go func() {
		select {
		case <-ctx.Done():
			log.Debug().Msg("shutdown tracer")
			if err := closer.Close(); err != nil {
				log.Error().Msgf("failed to close tracer: %v", err)
			}
		}
	}()
}

func runGrpc(ctx context.Context, wg *sync.WaitGroup, dbConn *sql.DB, kafkaAddrs []string) {
	listen, err := net.Listen("tcp", cfg.Data.Grpc_Listen)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	tracer := opentracing.GlobalTracer()
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(db.NewInterceptorWithDB(dbConn)),
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
	)

	kProd, err := producer.BuildService(kafkaAddrs, cfg.Data.Inserts_chank_size)
	if err != nil {
		log.Fatal().Msgf("failed to connect to kafka: %v", err)
	}

	apiSrv := api.NewOcpInstructionApi(repoService.BuildRequestService(), kProd)
	desc.RegisterOcpInstructionServer(s, apiSrv)

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := s.Serve(listen); err != nil {
			log.Fatal().Msgf("failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			log.Debug().Msg("shutdown gRPC")
			s.GracefulStop()
		}
	}()
}

func runJSON(ctx context.Context, wg *sync.WaitGroup) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpInstructionHandlerFromEndpoint(ctx, mux, cfg.Data.Grpc_Endpoint, opts)
	if err != nil {
		log.Fatal().Msgf("failed json registration haldler: %v", err)
	}

	s := http.Server{Addr: cfg.Data.Grpc_Jsongw_Listen, Handler: mux}

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("failed to serve json: %v", err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			log.Debug().Msg("shutdown JSON gw")
			err := s.Shutdown(ctx)
			if err != nil && err != context.Canceled {
				log.Error().Msgf("failed to shutdown json: %v", err)
			}
		}
	}()
}

func runConsumerService(ctx context.Context, wg *sync.WaitGroup, dbConn *sql.DB, kafkaAddrs []string) {
	ctx = db.NewContext(ctx, dbConn)

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("consumer")
	ctx = opentracing.ContextWithSpan(ctx, span)

	serv := consumer.BuildService(kafkaAddrs, "InstructionCUDGroup", "InstructionCUD")

	go func() {
		wg.Add(1)
		defer wg.Done()

		if err := serv.Consuming(ctx); err != nil {
			log.Fatal().Msgf("failed to start consuming: %v", err)
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			log.Debug().Msg("shutdown consumer")
			if err := serv.Close(); err != nil {
				log.Error().Msgf("failed to shutdown consuming: %v", err)
			}
		}
	}()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	if err := cfg.Load(); err != nil {
		log.Fatal().Err(err)
	}

	dbConn := db.Connect(cfg.Data.Pg_conn)

	log.Debug().Msg("run consumer")
	runConsumerService(ctx, wg, dbConn, cfg.Data.Kafka_addr)

	log.Debug().Msg("run metrics")
	metrics.Run(ctx, wg)

	log.Debug().Msg("run tracing")
	runTracing(ctx)

	log.Debug().Msg("run json")
	runJSON(ctx, wg)

	log.Debug().Msg("run app")
	runGrpc(ctx, wg, dbConn, cfg.Data.Kafka_addr)

	go func() {
		termChan := make(chan os.Signal)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
		<-termChan
		cancel()
	}()

	wg.Wait()

	if err := dbConn.Close(); err != nil {
		log.Fatal().Err(err)
	}
}
