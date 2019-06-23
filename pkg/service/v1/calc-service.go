package v1

import (
	"context"
    "fmt"

    "github.com/timaa2k/go-grpc/pkg/api/v1"
)

type calcServiceServer struct{}

func NewCalcServiceServer() v1.CalcServiceServer {
	return &calcServiceServer{}
}

func (s *calcServiceServer) Add(ctx context.Context, req *v1.CalcRequest) (*v1.CalcResponse, error) {
	return &v1.CalcResponse{Result: req.Operands.Left + req.Operands.Right}, nil
}

func (s *calcServiceServer) Sub(ctx context.Context, req *v1.CalcRequest) (*v1.CalcResponse, error) {
	return &v1.CalcResponse{Result: req.Operands.Left - req.Operands.Right}, nil
}

func (s *calcServiceServer) Mul(ctx context.Context, req *v1.CalcRequest) (*v1.CalcResponse, error) {
	return &v1.CalcResponse{Result: req.Operands.Left * req.Operands.Right}, nil
}

func (s *calcServiceServer) Div(ctx context.Context, req *v1.CalcRequest) (*v1.CalcResponse, error) {
    if req.Operands.Right == 0.0 {
        return nil, fmt.Errorf("Division by zero")
    }
	return &v1.CalcResponse{Result: req.Operands.Left / req.Operands.Right}, nil
}
