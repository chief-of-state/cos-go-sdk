package cos

import (
	"context"
	"os"
	"testing"

	helloworldv1 "github.com/chief-of-state/cos-go-binding/gen/helloworld/v1"
	"github.com/stretchr/testify/suite"
)

type cosSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGetClient(t *testing.T) {
	suite.Run(t, new(cosSuite))
}

func (s cosSuite) TestGetClient() {
	s.Run("with valid data", func() {
		// create the context
		ctx := context.TODO()
		// clear env var
		s.Assert().NoError(os.Unsetenv("COS_HOST"))
		s.Assert().NoError(os.Unsetenv("COS_PORT"))
		// set the env vars
		s.Assert().NoError(os.Setenv("COS_HOST", "localhost"))
		s.Assert().NoError(os.Setenv("COS_PORT", "9000"))
		s.Assert().NotPanics(func() {
			// get the instance of CoS Client
			client := GetClient[*helloworldv1.HelloReply](ctx)
			s.Assert().NotNil(client)
		})
	})
	s.Run("when env vars are not set", func() {
		// create the context
		ctx := context.TODO()
		// clear env var
		s.Assert().NoError(os.Unsetenv("COS_HOST"))
		s.Assert().NoError(os.Unsetenv("COS_PORT"))
		s.Assert().Panics(func() {
			// get the instance of CoS Client
			client := GetClient[*helloworldv1.HelloReply](ctx)
			s.Assert().Nil(client)
		})
	})
}
