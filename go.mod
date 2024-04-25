module github.com/Anush008/qclient-go

go 1.22.2

require (
	github.com/Anush008/qclient-go/grpc v0.0.0-00010101000000-000000000000
	github.com/Anush008/qclient-go/qdrant v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/Anush008/qclient-go/qdrant => ./qdrant

replace github.com/Anush008/qclient-go/grpc => ./grpc
