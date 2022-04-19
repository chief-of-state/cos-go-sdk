package cos

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type configSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestConfig(t *testing.T) {
	suite.Run(t, new(configSuite))
}

func (s *configSuite) TestGetConfigFromEnv() {
	s.Run("with valid env vars set", func() {
		err := os.Setenv("COS_HOST", "localhost")
		s.Assert().NoError(err)
		err = os.Setenv("COS_PORT", "447")
		s.Assert().NoError(err)

		// load the config and do some assertions
		cfg, err := GetConfigFromEnv()
		s.Assert().NoError(err)
		s.Assert().NotNil(cfg)
		s.Assert().Equal("localhost", cfg.CosHost)
		s.Assert().Equal(447, cfg.CosPort)

		// clean the env vars
		os.Clearenv()
	})
	s.Run("with missing env vars", func() {
		// set the required env vars
		err := os.Setenv("COS_HOST", "localhost")
		s.Assert().NoError(err)

		// load the config and do some assertions
		cfg, err := GetConfigFromEnv()
		s.Assert().Error(err)
		s.Assert().Nil(cfg)
		// clean the env vars
		os.Clearenv()
	})
}
