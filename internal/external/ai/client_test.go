package ai

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(clientSuite))
}

type clientSuite struct {
	suite.Suite
}

//func (s *clientSuite) TestPrompt() {
//	cl := NewClient("")
//
//	res, err := cl.Prompt(context.Background(), models.Create{
//		Image:       "postgres",
//		Replicas:    3,
//		Description: "",
//	})
//	s.Require().NoError(err)
//	s.Require().NotEmpty(res)
//	s.Require().True(strings.Contains(res, "name: postgres"))
//	s.Require().True(strings.Contains(res, "replicas: 3"))
//
//	res, err = cl.Prompt(context.Background(), models.Create{
//		Image:       "cowsay",
//		Replicas:    2,
//		Description: "",
//	})
//	s.Require().NoError(err)
//	s.Require().NotEmpty(res)
//	s.Require().True(strings.Contains(res, "name: cowsay"))
//	s.Require().True(strings.Contains(res, "replicas: 2"))
//}
