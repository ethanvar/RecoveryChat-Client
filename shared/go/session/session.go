// Package session exposes the chat helpers used by NativeChatModule.cpp.
package session

import (
	"context"
	"fmt"

	"berty.tech/weshnet/v2"
	"berty.tech/weshnet/v2/pkg/protocoltypes"
)

func ServiceCreatePersistent() (*protocoltypes.ServiceGetConfiguration_Reply, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	path := "data"
	client, err := weshnet.NewPersistentServiceClient(path)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	config, err := client.ServiceGetConfiguration(ctx,
		&protocoltypes.ServiceGetConfiguration_Request{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Hello, world! My peer ID is", config.PeerId)

	return config, nil
}
