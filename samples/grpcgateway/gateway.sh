#!/bin/bash

TARGETDIR=pkg/grpc-gateway
mkdir -p $TARGETDIR/cmd
mkdir -p $TARGETDIR/autoregister

# the gateway
spectools exportAsYaml | simple-generator -t scripts/grpcgateway/command.tpl > $TARGETDIR/cmd/cmd.go
spectools exportAsYaml | simple-generator -t scripts/grpcgateway/gateway.tpl > $TARGETDIR/gateway.go

# the registration of the services..
spectools exportAsYaml | simple-generator -t scripts/grpcgateway/autoregister.tpl > $TARGETDIR/autoregister/autoregister.go

echo "gateway sources generated
"

#go run pkg/grpc-gateway/cmd/cmd.go
echo "run your grpc gateway with
go run pkg/grpc-gateway/cmd/cmd.go
"

#go build pkg/grpc-gateway/cmd/cmd.go
echo "build your grpc gateway with
go build pkg/grpc-gateway/cmd/cmd.go
"