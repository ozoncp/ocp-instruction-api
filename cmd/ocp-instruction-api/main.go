package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-instruction-api/internal/consumer"
	"github.com/ozoncp/ocp-instruction-api/internal/metrics"
	"github.com/ozoncp/ocp-instruction-api/internal/producer"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	"github.com/ozoncp/ocp-instruction-api/pkg/db"
	"github.com/uber/jaeger-client-go"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-instruction-api/internal/app/ocp-instruction-api"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"

	zerolog "github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
)

const (
	grpcPort           = ":8080"
	grpcServerEndpoint = "localhost:8080"
	jsonGwListen       = ":8081"
	metricsListen      = ":9100"
)

func initDB() (*sql.DB, error) {
	connString := os.Getenv("PG_CONN_STR")
	if connString == "" {
		return nil, errors.New("Env PG_CONN_STR is not set")
	}

	dbConn := db.Connect(connString)

	return dbConn, nil
}

func kafkaAddr() ([]string, error) {
	kafkaAddr := os.Getenv("KAFKA_ADDR")
	if kafkaAddr == "" {
		return nil, errors.New("Env KAFKA_ADDR is not set")
	}

	addrs := make([]string, 1)
	addrs = append(addrs, kafkaAddr)

	return addrs, nil
}

func runMetrics() {
	metrics.Register()

	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(metricsListen, nil)
	if err != nil {
		log.Fatalf("failed to serve metrics: %v", err)
	}
}

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
		log.Fatalf("trace init failed: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	//_ = closer.Close()
}

func run(dbConn *sql.DB, kafkaAddrs []string) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer := opentracing.GlobalTracer()
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(db.NewInterceptorWithDB(dbConn)),
		grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer)),
	)

	ch_s, err := strconv.Atoi(os.Getenv("CREATE_CHUNK_SIZE"))
	if err != nil || ch_s == 0 {
		ch_s = 3
	}
	kProd, err := producer.BuildService(kafkaAddrs, ch_s)
	if err != nil {
		log.Fatalf("failed to connect to kafka: %v", err)
	}

	apiSrv := api.NewOcpInstructionApi(repoService.BuildRequestService(), kProd)
	desc.RegisterOcpInstructionServer(s, apiSrv)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpInstructionHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed json registration haldler: %v", err)
	}

	err = http.ListenAndServe(jsonGwListen, mux)
	if err != nil {
		log.Fatalf("failed to serve json: %v", err)
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
		log.Fatalf("failed to start consuming: %v", err)
	}
}

func main() {
	conn, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	kAddr, err := kafkaAddr()
	if err != nil {
		log.Fatal(err)
	}

	zerolog.Debug().Msg("run json")
	go runJSON()

	zerolog.Debug().Msg("run consumer")
	runConsumerService(conn, kAddr)

	zerolog.Debug().Msg("run metrics")
	go runMetrics()

	zerolog.Debug().Msg("run tracing")
	jaegerAddr := os.Getenv("JAEGER_ADDR")
	if jaegerAddr == "" {
		log.Fatal(errors.New("Env JAEGER_ADDR is not set"))
	}
	runTracing(jaegerAddr)

	zerolog.Debug().Msg("run app")

	if err := run(conn, kAddr); err != nil {
		log.Fatal(err)
	}
}
