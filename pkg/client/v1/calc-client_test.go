package v1

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
    api "github.com/timaa2k/go-grpc/pkg/api/v1"
    "github.com/timaa2k/go-grpc/pkg/service/v1"
	"google.golang.org/grpc"
)

type TestSuite struct {
	suite.Suite
	serverAddr string
	server     *grpc.Server
    client     *CalcClient
	done       chan struct{}
}

func (suite *TestSuite) SetupSuite() {
	listen, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	suite.serverAddr = listen.Addr().String()

	suite.server = grpc.NewServer()
	v1API := v1.NewCalcServiceServer()
	api.RegisterCalcServiceServer(suite.server, v1API)

	errChan := make(chan error)
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("Starting test gRPC server...")
		if err := suite.server.Serve(listen); err != nil {
			errChan <- err
		}
	}()

	go func() {
		select {
		case err := <-errChan:
			log.Printf("Fatal error: %v\n", err)
		case <-stopChan:
			suite.server.GracefulStop()
        }
		log.Fatalf("Test gRPC server has stopped.")
	}()
}

func (suite *TestSuite) SetupTest() {
    client, err := NewCalcClient(suite.serverAddr)
    if err != nil {
        log.Fatal(err)
    }
    suite.client = client
}

func (suite *TestSuite) TestConnectionError() {
    _, err := NewCalcClient("non-existant:1234")
	assert.NotNil(suite.T(), err)
}

func (suite *TestSuite) TestClientAdd() {
	result, err := suite.client.Add(2.5, 3.0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(5.5), result)
}

func (suite *TestSuite) TestClientSub() {
	result, err := suite.client.Sub(2.5, 3.0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-0.5), result)
}

func (suite *TestSuite) TestClientMul() {
	result, err := suite.client.Mul(-2.5, 2.0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-5.0), result)
}

func (suite *TestSuite) TestClientDiv() {
	result, err := suite.client.Div(2, 5.0)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(0.4), result)
}

func (suite *TestSuite) TearDownTest() {
	suite.client.Disconnect()
}

func (suite *TestSuite) TearDownSuite() {
	log.Println("Shutting down test gRPC server...")
	suite.server.GracefulStop()
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
