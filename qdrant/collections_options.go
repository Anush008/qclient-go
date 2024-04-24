package qdrant

import (
	qdrant "github.com/qdrant/go-client/grpc"
)

// ============ BEGIN CREATE COLLECTION OPTIONS ==============
type CreateCollectionOption func(c *qdrant.CreateCollection)

func WithVectorConfig(params *qdrant.VectorParams) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.VectorsConfig = &qdrant.VectorsConfig{
			Config: &qdrant.VectorsConfig_Params{
				Params: params,
			},
		}
	}
}

func WithVectorConfigMap(params map[string]*qdrant.VectorParams) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.VectorsConfig = &qdrant.VectorsConfig{
			Config: &qdrant.VectorsConfig_ParamsMap{
				ParamsMap: &qdrant.VectorParamsMap{
					Map: params,
				},
			},
		}
	}
}

func WithHnswConfig(config *qdrant.HnswConfigDiff) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.HnswConfig = config
	}
}

func WithWalConfig(config *qdrant.WalConfigDiff) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.WalConfig = config
	}
}

func WithOptimizersConfig(config *qdrant.OptimizersConfigDiff) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.OptimizersConfig = config
	}
}

func WithShardNumber(number uint32) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.ShardNumber = &number
	}
}

func WithOnDiskPayload(value bool) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.OnDiskPayload = &value
	}
}

func WithCreateTimeout(timeout uint64) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.Timeout = &timeout
	}
}

func WithReplicationFactor(factor uint32) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.ReplicationFactor = &factor
	}
}

func WithWriteConsistencyFactor(factor uint32) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.WriteConsistencyFactor = &factor
	}
}

func WithInitFromCollection(collection string) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.InitFromCollection = &collection
	}
}

func WithScalarQuantization(config *qdrant.ScalarQuantization) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.QuantizationConfig = &qdrant.QuantizationConfig{
			Quantization: &qdrant.QuantizationConfig_Scalar{
				Scalar: config,
			},
		}
	}
}

func WithBinaryQuantization(config *qdrant.BinaryQuantization) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.QuantizationConfig = &qdrant.QuantizationConfig{
			Quantization: &qdrant.QuantizationConfig_Binary{
				Binary: config,
			},
		}
	}
}

func WithProductQuantization(config *qdrant.ProductQuantization) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.QuantizationConfig = &qdrant.QuantizationConfig{
			Quantization: &qdrant.QuantizationConfig_Product{
				Product: config,
			},
		}
	}
}

func WithShardingMethod(method qdrant.ShardingMethod) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.ShardingMethod = &method
	}
}

func WithSparseVectorsConfig(params map[string]*qdrant.SparseVectorParams) CreateCollectionOption {
	return func(c *qdrant.CreateCollection) {
		c.SparseVectorsConfig = &qdrant.SparseVectorConfig{
			Map: params,
		}
	}
}

// ============ END CREATE COLLECTION OPTIONS ==============

// ============ BEGIN DELETE COLLECTION OPTIONS ==============
type DeleteCollectionOption func(c *qdrant.DeleteCollection)

func WithDeleteTimeout(timeout uint64) DeleteCollectionOption {
	return func(c *qdrant.DeleteCollection) {
		c.Timeout = &timeout
	}
}

// ============ END DELETE COLLECTION OPTIONS ==============
