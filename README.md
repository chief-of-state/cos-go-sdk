# cos-go-binding
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/chief-of-state/cos-go-binding/main)

Chief of State go bindings provides an easy way to create and use the a Chief-of-State client in Golang.

## Features
With the cos-go-binding, one can:
- Create a typed chief-of-state client.
- `ProcessCommand` Processes a generic `proto.Message` command and returns the typed state from the CoS service
- `GetState` Gets a typed state from the CoS service


### Global environment variables
| environment variable | description | default | required |
|--- | --- | --- | --- |
| COS_HOST | The host of the cos server | | Y |
| COS_PORT | The port of the cos server | | Y |

## Example

```
ctx := context.TODO()
entityID := "some-entity-id"

// gets cos config from environment
cfg, err := cos.GetConfigFromEnv()
if err != nil {
	// handles the error
	panic(err)
}

// creates the client
client, err := cos.NewClient[*someProtobufType](ctx, cfg.CosHost, cfg.CosPort)
if err != nil {
	// handles the error
	panic(err)
}

// sends a command to the cos service
state, metadata, err := client.ProcessCommand(ctx, entityID, &someRequest{})
if err != nil {
	// handles the error
	panic(err)
}
// prints the metadata and resulting state
fmt.Println(metadata)
fmt.Println(state) // the state type will be *someProtobufType

// given the entity id gets the state from the cos service
state, metadata, err = client.GetState(ctx, entityID)
if err != nil {
	// handles the error
	panic(err)
}
// prints the metadata and resulting state
fmt.Println(metadata)
fmt.Println(state) // the state type will be *someProtobufType
```
