module github.com/aserto-dev/go-grpc-authz

go 1.16

// replace github.com/aserto-dev/mage-loot => ../mage-loot

require (
	github.com/aserto-dev/go-grpc v0.0.0-20210812160848-636c41fbe0dc
	github.com/aserto-dev/mage-loot v0.4.12
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/magefile/mage v1.11.0
	google.golang.org/genproto v0.0.0-20210811021853-ddbe55d93216
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
