package qdrant

import (
	"github.com/qdrant/go-client/grpc"
)

type PointStructBuilder struct {
	point *grpc.PointStruct
}

func newPointStructBuilder() *PointStructBuilder {
	return &PointStructBuilder{
		point: &grpc.PointStruct{},
	}
}

func WithNumId(id uint64) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		b.point.Id = &grpc.PointId{
			PointIdOptions: &grpc.PointId_Num{
				Num: id,
			},
		}
	}
}

func WithUuid(id string) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		b.point.Id = &grpc.PointId{
			PointIdOptions: &grpc.PointId_Uuid{
				Uuid: id,
			},
		}
	}
}

func WithVector(vector []float32) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		b.point.Vectors = &grpc.Vectors{
			VectorsOptions: &grpc.Vectors_Vector{
				Vector: &grpc.Vector{
					Data: vector,
				},
			},
		}
	}
}

func WithNamedVectors(vectors map[string][]float32) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		namedVectors := make(map[string]*grpc.Vector)
		for name, vector := range vectors {
			namedVectors[name] = &grpc.Vector{
				Data: vector,
			}
		}
		b.point.Vectors = &grpc.Vectors{
			VectorsOptions: &grpc.Vectors_Vectors{
				Vectors: &grpc.NamedVectors{
					Vectors: namedVectors,
				},
			},
		}
	}
}

func WithVectors(vectors map[string]*grpc.Vector) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		b.point.Vectors = &grpc.Vectors{
			VectorsOptions: &grpc.Vectors_Vectors{
				Vectors: &grpc.NamedVectors{
					Vectors: vectors,
				},
			},
		}
	}
}

func WithPayload(payload map[string]interface{}) func(*PointStructBuilder) {
	return func(b *PointStructBuilder) {
		b.point.Payload = NewValueMap(payload)
	}
}

func NewPointStruct(builderFuncs ...func(*PointStructBuilder)) *grpc.PointStruct {
	builder := newPointStructBuilder()
	for _, f := range builderFuncs {
		f(builder)
	}
	return builder.point
}
