package go_sdk

import (
	"crypto/tls"
	"fmt"

	"github.com/gorillamq/go-sdk/internal"
	"github.com/gorillamq/go-sdk/pkg"
)

// Client is our client handler which enables a user
// to communicate with broker server.
type Client interface {
	// Publish messages to broker
	Publish(topic string, data []byte) error

	// Subscribe over broker
	Subscribe(topic string, handler internal.MessageHandler)

	// Unsubscribe from broker
	Unsubscribe(topic string)
}

// NewClient creates a new client to connect to broker server.
func NewClient(uri string) (Client, error) {
	url, err := pkg.UnpackURL(uri)
	if err != nil {
		return nil, fmt.Errorf("invalid uri: %w", err)
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", url.Address, tlsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client, err := internal.NewClient(conn, url.Auth)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	return client, nil
}
