package qdrant

import (
	"context"
	"errors"

	"github.com/Anush008/qclient-go/grpc"
)

// Lists all the collections in the Qdrant instance
func (c *Client) ListCollections(ctx context.Context) ([]string, error) {
	resp, err := c.Grpc().Collections.List(ctx, &grpc.ListCollectionsRequest{})
	if err != nil {
		return nil, err
	}

	names := make([]string, len(resp.GetCollections()))
	for i, description := range resp.GetCollections() {
		names[i] = description.GetName()
	}

	return names, nil
}

// Creates a new Qdrant collection with the given name and options
func (c *Client) CreateCollection(ctx context.Context, name string, opts ...CreateCollectionOption) error {

	req := &grpc.CreateCollection{
		CollectionName: name,
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.Grpc().Collections.Create(ctx, req)

	if err != nil {
		return err
	}

	if !resp.Result {
		return errors.New("collection not created")
	}
	return nil
}

func (c *Client) DeleteCollection(ctx context.Context, name string, opts ...DeleteCollectionOption) error {
	req := &grpc.DeleteCollection{
		CollectionName: name,
	}

	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.Grpc().Collections.Delete(ctx, req)

	if err != nil {
		return err
	}

	if !resp.Result {
		return errors.New("collection not deleted")
	}

	return nil
}
