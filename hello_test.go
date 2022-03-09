package goptr

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HelloSuite struct {
	suite.Suite
}

func (s *HelloSuite) TestSum() {
	s.Equal(3, Sum(1, 2))
}

func TestHelloSuite(t *testing.T) {
	suite.Run(t, new(HelloSuite))
}
