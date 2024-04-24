package qdrant

import (
	"context"

	qdrant "github.com/qdrant/go-client/grpc"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn        *grpc.ClientConn
	Qdrant      qdrant.QdrantClient
	Collections qdrant.CollectionsClient
	Points      qdrant.PointsClient
	Snapshots   qdrant.SnapshotsClient
}

// connect connect to Service
func NewGrpcClient(ctx context.Context, config *Config) (*GrpcClient, error) {
	addr := config.GetAddr()

	grpcOptions := config.GrpcOptions

	transportCreds := config.GetTransportCreds()

	grpcOptions = append(grpcOptions, transportCreds)

	conn, err := grpc.DialContext(ctx, addr, grpcOptions...)

	if err != nil {
		return nil, err
	}

	return NewGrpcClientFromConn(conn), nil
}

func NewGrpcClientFromConn(conn *grpc.ClientConn) *GrpcClient {
	c := &GrpcClient{}
	c.conn = conn
	c.Qdrant = qdrant.NewQdrantClient(conn)
	c.Points = qdrant.NewPointsClient(conn)
	c.Collections = qdrant.NewCollectionsClient(conn)
	c.Snapshots = qdrant.NewSnapshotsClient(conn)
	return c
}

// Close close the connection
func (c *GrpcClient) Close() error {
	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}
