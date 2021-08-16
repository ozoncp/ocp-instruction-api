package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api "github.com/ozoncp/ocp-instruction-api/internal/app/ocp-instruction-api"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"

	zerolog "github.com/rs/zerolog/log"
)

const (
	grpcPort           = ":80"
	grpcServerEndpoint = "localhost:80"
	jsonGwListen       = ":8081"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpInstructionServer(s, api.NewOcpInstructionApi())

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
		panic(err)
	}

	err = http.ListenAndServe(jsonGwListen, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	zerolog.Debug().Msg("run")

	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
