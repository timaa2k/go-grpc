package cmd

import (
	"context"

	"github.com/timaa2k/go-grpc/pkg/service/v1"
	"github.com/timaa2k/go-grpc/pkg/protocol/grpc"
)

// RunServer starts a gRPC server on the given address.
func RunServer(addr string) error {
	ctx := context.Background()
	v1API := v1.NewCalcServiceServer()
	return grpc.RunServer(ctx, v1API, addr)
}
