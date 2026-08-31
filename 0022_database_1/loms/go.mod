module github.com/igoroutine-courses/microservices.ecommerce.loms

go 1.25.0

require (
	github.com/caarlos0/env/v10 v10.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.28.0
	github.com/igoroutine-courses/microservices.ecommerce.pkg v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v5 v5.8.0
	github.com/pressly/goose/v3 v3.27.0
	go.uber.org/zap v1.27.1
	google.golang.org/grpc v1.79.2
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.3.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.50.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260209200024-4cfbd4190f57 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260217215200-42d3e9bedb6d // indirect
)

replace github.com/igoroutine-courses/microservices.ecommerce.pkg => ../pkg
