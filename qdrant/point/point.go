package point

import (
	"github.com/Anush008/qclient-go/grpc"
	"github.com/Anush008/qclient-go/qdrant/value"
)

type PointStructBuilder struct {
	point *grpc.PointStruct
}

type PointStructBuilderOption func(p *PointStructBuilder)

func NewPointStruct(opts ...PointStructBuilderOption) *grpc.PointStruct {
	builder := newPointStructBuilder()
	for _, f := range opts {
		f(builder)
	}
	return builder.point
}

func WithNumId(id uint64) PointStructBuilderOption {
	return func(b *PointStructBuilder) {
		b.point.Id = &grpc.PointId{
			PointIdOptions: &grpc.PointId_Num{
				Num: id,
			},
		}
	}
}

func WithUuid(id string) PointStructBuilderOption {
	return func(b *PointStructBuilder) {
		b.point.Id = &grpc.PointId{
			PointIdOptions: &grpc.PointId_Uuid{
				Uuid: id,
			},
		}
	}
}

func WithVector(vector []float32) PointStructBuilderOption {
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

func WithNamedVectors(vectors map[string][]float32) PointStructBuilderOption {
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

func WithVectors(vectors map[string]*grpc.Vector) PointStructBuilderOption {
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

func WithPayload(payload map[string]interface{}) PointStructBuilderOption {
	return func(b *PointStructBuilder) {
		b.point.Payload = value.NewValueMap(payload)
	}
}

func newPointStructBuilder() *PointStructBuilder {
	return &PointStructBuilder{
		point: &grpc.PointStruct{},
	}
}
