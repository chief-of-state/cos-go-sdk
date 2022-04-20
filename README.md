# cos-go-sdk
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/chief-of-state/cos-go-sdk/main)

Chief of State go SDK provides an easy way to create and use the a Chief-of-State client in Golang.

## Usage
```bash
$ go get -u github.com/chief-of-state/cos-go-sdk
```

## Features
With the cos-go-sdk, one can:
- Create a typed chief-of-state client.
- `ProcessCommand` Processes a generic `proto.Message` command and returns the `typed` state from the CoS service
- `GetState` Gets a `typed` state from the CoS service
- Test with generated mockers located in [cosmocks](cosmocks/cospb/chief_of_state/v1/)

## Example
#### Using typed functions

```
import (
	"context"
	"fmt"

	"github.com/chief-of-state/cos-go-sdk/cos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// mock a proto.Message state
type fakeState struct{}

func (*fakeState) ProtoReflect() protoreflect.Message { return nil }

// mock a proto.Message command
type fakeCommand struct{}

func (*fakeCommand) ProtoReflect() protoreflect.Message { return nil }

func main() {

	ctx := context.TODO()
	entityID := "some-entity-id"
	target := "localhost:9000

	// create the grpc client
	grpcClient, err := grpc.DialContext(ctx, target)
	if err != nil {
		// handles the error
		panic(err)
	}

	// creates the client
	client, err := cos.NewClient[*fakeState](grpcClient)
	if err != nil {
		// handles the error
		panic(err)
	}

	// sends a command to the cos service
	state, metadata, err := client.ProcessCommand(ctx, entityID, &fakeCommand{})
	if err != nil {
		// handles the error
		panic(err)
	}
	// prints the metadata and resulting state
	fmt.Println(metadata)
	fmt.Println(state) // the state type will be *fakeState

	// given the entity id gets the state from the cos service
	state, metadata, err = client.GetState(ctx, entityID)
	if err != nil {
		// handles the error
		panic(err)
	}
	// prints the metadata and resulting state
	fmt.Println(metadata)
	fmt.Println(state) // the state type will be *fakeState
}
```
