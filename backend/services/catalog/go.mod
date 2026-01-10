module github.com/kinoshitatakumi/opti/services/catalog

go 1.24.1

require (
	connectrpc.com/connect v1.19.1
	github.com/google/uuid v1.6.0
	github.com/kinoshitatakumi/opti/gen/go v0.0.0
	github.com/kinoshitatakumi/opti/pkg v0.0.0
	golang.org/x/net v0.48.0
)

require (
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/kinoshitatakumi/opti/gen/go => ../../gen/go

replace github.com/kinoshitatakumi/opti/pkg => ../../pkg
