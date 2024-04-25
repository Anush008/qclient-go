package qdrant

import (
	"github.com/Anush008/qclient-go/grpc"
)

// ============ BEGIN COUNT OPTIONS ==============

type CountPointsOption func(c *grpc.CountPoints)

func WithCountFilter(filter *grpc.Filter) CountPointsOption {
	return func(c *grpc.CountPoints) {
		c.Filter = filter
	}
}

func WithCountExact(exact bool) CountPointsOption {
	return func(c *grpc.CountPoints) {
		c.Exact = &exact
	}
}

func WithCountConsistency(consistency grpc.ReadConsistencyType) CountPointsOption {
	return func(c *grpc.CountPoints) {
		c.ReadConsistency = &grpc.ReadConsistency{
			Value: &grpc.ReadConsistency_Type{
				Type: consistency,
			},
		}
	}
}

func WithCountShardKeySelector(selector *grpc.ShardKeySelector) CountPointsOption {
	return func(c *grpc.CountPoints) {
		c.ShardKeySelector = selector
	}
}

// =========== END COUNT OPTIONS ================

// ============ BEGIN UPSERT OPTIONS ==============

type UpsertPointsOption func(c *grpc.UpsertPoints)

func WithPoints(points []*grpc.PointStruct) UpsertPointsOption {
	return func(c *grpc.UpsertPoints) {
		c.Points = points
	}
}

func WithUpsertWait(wait bool) UpsertPointsOption {
	return func(c *grpc.UpsertPoints) {
		c.Wait = &wait
	}
}

func WithUpsertOrdering(ordering grpc.WriteOrderingType) UpsertPointsOption {
	return func(c *grpc.UpsertPoints) {
		c.Ordering = &grpc.WriteOrdering{
			Type: ordering,
		}
	}
}

func WithUpsertShardKeySelector(shardKeys []*grpc.ShardKey) UpsertPointsOption {
	return func(c *grpc.UpsertPoints) {
		c.ShardKeySelector = &grpc.ShardKeySelector{
			ShardKeys: shardKeys,
		}
	}
}

// =========== END UPSERT OPTIONS ================
