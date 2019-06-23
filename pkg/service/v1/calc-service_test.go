package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
    "github.com/timaa2k/go-grpc/pkg/api/v1"
)

type TestSuite struct {
	suite.Suite
	server v1.CalcServiceServer
}

func (suite *TestSuite) SetupTest() {
	suite.server = NewCalcServiceServer()
}

func (suite *TestSuite) TestAddition() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  2.5,
			Right: 3.0,
		},
	}
	resp, err := suite.server.Add(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(5.5), resp.Result)
}

func (suite *TestSuite) TestNegativeAddition() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  -3.5,
			Right: 3.0,
		},
	}
	resp, err := suite.server.Add(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-0.5), resp.Result)
}

func (suite *TestSuite) TestSubtraction() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  1.5,
			Right: 2.0,
		},
	}
	resp, err := suite.server.Sub(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-0.5), resp.Result)
}

func (suite *TestSuite) TestNegativeSubtraction() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  -1.5,
			Right: -2.0,
		},
	}
	resp, err := suite.server.Sub(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(0.5), resp.Result)
}

func (suite *TestSuite) TestMultiplication() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  3.0,
			Right: 2.5,
		},
	}
	resp, err := suite.server.Mul(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(7.5), resp.Result)
}

func (suite *TestSuite) TestNegativeMultiplication() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  3.0,
			Right: -2.5,
		},
	}
	resp, err := suite.server.Mul(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-7.5), resp.Result)
}

func (suite *TestSuite) TestDivision() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  11,
			Right: 3,
		},
	}
	resp, err := suite.server.Div(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(3.66666666667), resp.Result)
}

func (suite *TestSuite) TestNegativeDivision() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  3,
			Right: -11,
		},
	}
	resp, err := suite.server.Div(context.TODO(), &req)
    assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(-0.27272727272), resp.Result)
}

func (suite *TestSuite) TestDivisionByZero() {
	req := v1.CalcRequest{
		Operands: &v1.Operands{
			Left:  3,
			Right: 0,
		},
	}
	_, err := suite.server.Div(context.TODO(), &req)
    assert.NotNil(suite.T(), err)
    assert.Equal(suite.T(), "Division by zero", err.Error())
}

func TestRunTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
