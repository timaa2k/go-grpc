package v1

import (
	"context"
	"time"

	"github.com/pkg/errors"
    "github.com/timaa2k/go-grpc/pkg/api/v1"
    "google.golang.org/grpc"
)

type CalcClient struct {
	Conn *grpc.ClientConn
}

func NewCalcClient(serverAddr string) (*CalcClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()
	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrap(err, "Failed gRPC connection")
	}
	return &CalcClient{Conn: conn}, nil
}

func (c *CalcClient) Disconnect() {
	c.Conn.Close()
}

func (c *CalcClient) Add(a, b float32) (float32, error) {
	calcServiceClient := v1.NewCalcServiceClient(c.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  a,
			Right: b,
		},
	}
	resp, err := calcServiceClient.Add(ctx, &req)
	if err != nil {
        return 0, errors.Wrap(err, "Addition failed")
	}
	return resp.Result, nil
}

func (c *CalcClient) Sub(a, b float32) (float32, error) {
	calcServiceClient := v1.NewCalcServiceClient(c.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  a,
			Right: b,
		},
	}
	resp, err := calcServiceClient.Sub(ctx, &req)
	if err != nil {
		return 0, errors.Wrap(err, "Subtraction failed")
	}
	return resp.Result, nil
}

func (c *CalcClient) Mul(a, b float32) (float32, error) {
	calcServiceClient := v1.NewCalcServiceClient(c.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  a,
			Right: b,
		},
	}
	resp, err := calcServiceClient.Mul(ctx, &req)
	if err != nil {
		return 0, errors.Wrap(err, "Multiplication failed")
	}
	return resp.Result, nil
}

func (c *CalcClient) Div(a, b float32) (float32, error) {
	calcServiceClient := v1.NewCalcServiceClient(c.Conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  a,
			Right: b,
		},
	}
	resp, err := calcServiceClient.Div(ctx, &req)
	if err != nil {
		return 0, errors.Wrap(err, "Division failed")
	}
	return resp.Result, nil
}
