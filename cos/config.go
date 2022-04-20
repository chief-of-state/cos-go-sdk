package cos

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// Config is used to make a call to CoS
type Config struct {
	CosHost string `env:"COS_HOST"` // CosHost is used to connect to ChiefOfState
	CosPort int    `env:"COS_PORT"` // CosPort is used to connect to ChiefOfState
}

// GetConfigFromEnv fetches the LoggingConfig from env vars
func GetConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	// all env vars are required
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.Parse(cfg, opts); err != nil {
		return nil, err
	}

	return cfg, nil
}

// GetTarget returns the cos target
func (c *Config) GetTarget() string {
	return fmt.Sprintf("%s:%v", c.CosHost, c.CosPort)
}
