# Generate the protoc command

[Back to overview](../readme.md)

This script will generate the proto2DomainTypes.sh and proto2DomainServices.sh scripts which will go 
through all proto files and produce the  the *.pb.go, *.pb.gw.go and the *.pb.go for the messages.

The [grpc gateway generator](grpcgateway/readme.md) relys on these files.
