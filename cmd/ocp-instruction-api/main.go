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

func runTracing(addr string) {
	config := jaegercfg.Configuration{
		ServiceName: "ocp_instruction_api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: addr,
		},
	}

	logger := jaegerlog.StdLogger
	metricsFactory := jaegermetrics.NullFactory

	tracer, _, err := config.NewTracer(
		jaegercfg.Logger(logger),
		jaegercfg.Metrics(metricsFactory),
	)

	if err != nil {
		log.Fatal().Msgf("trace init failed: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	//_ = closer.Close()
}

func run(dbConn *sql.DB, kafkaAddrs []string) error {
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

	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpInstructionHandlerFromEndpoint(ctx, mux, cfg.Data.Grpc_Endpoint, opts)
	if err != nil {
		log.Fatal().Msgf("failed json registration haldler: %v", err)
	}

	err = http.ListenAndServe(cfg.Data.Grpc_Jsongw_Listen, mux)
	if err != nil {
		log.Fatal().Msgf("failed to serve json: %v", err)
	}
}

func runConsumerService(dbConn *sql.DB, kafkaAddrs []string) {
	ctx := context.Background()
	ctx = db.NewContext(ctx, dbConn)

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan("consumer")
	ctx = opentracing.ContextWithSpan(ctx, span)

	serv := consumer.BuildService(kafkaAddrs, "InstructionCUDGroup", "InstructionCUD")
	err := serv.StartConsuming(ctx)
	if err != nil {
		log.Fatal().Msgf("failed to start consuming: %v", err)
	}
}

func main() {
	err := cfg.Load()
	if err != nil {
		log.Fatal().Err(err)
	}

	dbConn := db.Connect(cfg.Data.Pg_conn)

	log.Debug().Msg("run json")
	go runJSON()

	log.Debug().Msg("run consumer")
	runConsumerService(dbConn, cfg.Data.Kafka_addr)

	log.Debug().Msg("run metrics")
	metrics.Run()

	log.Debug().Msg("run tracing")
	runTracing(cfg.Data.Jaeger_addr)

	log.Debug().Msg("run app")
	if err := run(dbConn, cfg.Data.Kafka_addr); err != nil {
		log.Fatal().Err(err)
	}
}
