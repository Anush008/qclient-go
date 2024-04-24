package qdrant

import (
	"context"
)

type Client struct {
	grpcClient *GrpcClient
}

// NewClient creates a new client with the given configuration
func NewClient(ctx context.Context, config *Config) (*Client, error) {
	grpcClient, err := NewGrpcClient(ctx, config)

	if err != nil {
		return nil, err
	}

	return NewClientFromGrpc(grpcClient), nil
}

func NewClientFromGrpc(grpcClient *GrpcClient) *Client {
	return &Client{
		grpcClient,
	}
}

func (c *Client) Grpc() *GrpcClient {
	return c.grpcClient
}

// Closes the gRPC connection if it exists
func (c *Client) Close() error {
	return c.grpcClient.Close()
}
