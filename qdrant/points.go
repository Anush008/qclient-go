package qdrant

import (
	"context"

	"github.com/Anush008/qclient-go/grpc"
)

// Returns the number of points in the given collection. Optionally, a filter can be applied.
func (c *Client) Count(ctx context.Context, collectionName string, opts ...CountPointsOption) (uint64, error) {

	request := &grpc.CountPoints{
		CollectionName: collectionName,
	}

	for _, opt := range opts {
		opt(request)
	}

	resp, err := c.Grpc().Points.Count(ctx, request)

	if err != nil {
		return 0, err
	}
	return resp.GetResult().GetCount(), nil
}

// Performs an upsert operation on the given collection
func (c *Client) Upsert(ctx context.Context, collectionName string, opts ...UpsertPointsOption) (*grpc.UpdateResult, error) {
	request := &grpc.UpsertPoints{
		CollectionName: collectionName,
	}

	for _, opt := range opts {
		opt(request)
	}

	resp, err := c.Grpc().Points.Upsert(ctx, request)

	if err != nil {
		return nil, err
	}

	return resp.GetResult(), nil
}

func (c *Client) UpsertPoints(ctx context.Context, collectionName string, points ...*grpc.PointStruct) (*grpc.UpdateResult, error) {
	request := &grpc.UpsertPoints{
		CollectionName: collectionName,
		Points:         points,
	}

	resp, err := c.Grpc().Points.Upsert(ctx, request)

	if err != nil {
		return nil, err
	}

	return resp.GetResult(), nil
}
