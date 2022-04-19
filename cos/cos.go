package cos

import (
	"context"

	"github.com/chief-of-state/cos-go-binding/logging"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// GetClient instantiates an instance of cos.Client
func GetClient[T proto.Message](ctx context.Context) Client[T] {
	// get the CoS config from env vars
	cosConfig, err := GetConfigFromEnv()
	// log the error and panic
	if err != nil {
		logging.Panic(errors.Wrap(err, "unable to load config from env vars"))
	}
	// return a new instance of cos.Client
	client, err := NewClient[T](ctx, cosConfig.CosHost, cosConfig.CosPort)
	// log the error and panic
	if err != nil {
		logging.Panic(errors.Wrap(err, "unable to create an instance of CoS Client"))
	}
	return client
}
