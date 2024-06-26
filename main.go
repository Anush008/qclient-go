package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/Anush008/qclient-go/grpc"
	"github.com/Anush008/qclient-go/qdrant"
	"github.com/Anush008/qclient-go/qdrant/condition"
	"github.com/Anush008/qclient-go/qdrant/point"
)

func main() {
	// Timeouts, cancellations, etc can be passed in the context
	ctx := context.Background()

	host := "localhost"
	port := 6334
	collectionName := fmt.Sprintf("collection-%d", rand.Int63n(9999999))

	fmt.Println("\nInitializing a client with host ", host, " and port ", port)

	client, err := qdrant.NewClient(ctx, &qdrant.Config{
		Host: host,
		Port: port,
	})

	if err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Println("\nCreating a collection with name ", collectionName, " and scalar quantization")

	err = client.CreateCollection(ctx, collectionName,
		qdrant.WithVectorConfig(&grpc.VectorParams{
			Distance: grpc.Distance_Cosine,
			Size:     4,
		}),
		qdrant.WithScalarQuantization(&grpc.ScalarQuantization{
			Type: grpc.QuantizationType_Int8,
		}))

	if err != nil {
		panic(err)
	}

	collections, err := client.ListCollections(ctx)

	if err != nil {
		panic(err)
	}

	println("\nThe available collections are ")
	for _, collection := range collections {
		println("- ", collection)
	}

	points := []*grpc.PointStruct{
		point.NewPointStruct(
			point.WithNumId(1),
			point.WithVector([]float32{0.213, 0.123, 0.456, 0.4124}),
			point.WithPayload(map[string]interface{}{
				"some_number": 1,
				"some_bool":   true,
				"nested": map[string]interface{}{
					"key": false,
				},
			})),

		point.NewPointStruct(
			point.WithUuid("ced3caa0-e1f5-492b-8a7c-be4a274ea2bd"),
			point.WithVector([]float32{0.213, 0.123, 0.456, -0.4124})),

		point.NewPointStruct(
			point.WithNumId(71),
			point.WithVector([]float32{-0.213, -0.123, -0.456, 0.3412}),
			point.WithPayload(map[string]interface{}{
				"some_list": []interface{}{32, 523, false, "something"},
			})),
	}

	println("\n\nUpserting ", len(points), " points to the collection ", collectionName)

	result, err := client.Upsert(ctx, collectionName, qdrant.WithPoints(points), qdrant.WithUpsertWait(true))

	if err != nil {
		panic(err)
	}

	println("\n\nUpsert result: ", result.Status.String())

	println(("\n\nPoints Count without filter"))

	count, err := client.Count(ctx, collectionName)

	if err != nil {
		panic(err)
	}

	println("\n\nCount: ", count)

	println("\n\nPoints Count with filter")

	filter := &grpc.Filter{
		Must: []*grpc.Condition{
			condition.NewMatchBool("some_bool", true),
			condition.NewMatchInt("some_number", 1),
		},
	}

	count, err = client.Count(ctx, collectionName, qdrant.WithCountFilter(filter))

	if err != nil {
		panic(err)
	}

	println("\n\nCount: ", count)

	// println("\n\nDeleting the collection ", collectionName)
	// client.DeleteCollection(ctx, collectionName)
}
