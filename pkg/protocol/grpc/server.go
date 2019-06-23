package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

    "github.com/timaa2k/go-grpc/pkg/api/v1"
	"google.golang.org/grpc"
)

// RunServer runs a gRPC server with the given configuration
func RunServer(ctx context.Context, v1API v1.CalcServiceServer, addr string) error {

	fmt.Printf("Listening on address [%s]\n", addr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterCalcServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
