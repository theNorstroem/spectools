#!/usr/bin/env bash

DOMAINTYPESPATH={{.config.module}}/pkg/domaintypes
TARGETDIR="../../../pkg/grpcservices/"
# enable recursion for /**/*.xxx
shopt -s globstar dotglob

cd {{.config.build.proto.targetservicedir}}

FILES=./**/*.proto
for f in $FILES
do
  echo "Processing $f file..."

# the following lines are included, because the furo types are not complete :-(
# Mgoogle/type/date.proto=google.golang.org/genproto/genproto/googleapis/type/date,\
# Mgoogle/type/money.proto=google.golang.org/genproto/genproto/googleapis/type/money,\


protoc --proto_path=./ \
-I. \
-I$GOPATH/src/github.com/googleapis/googleapis \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I../Messages \
--go_out=\
{{$myDict := dict}}
{{- range $typename, $type := .types}}
{{- $k := print "M" $type.path "/" $type.typespec.__proto.targetfile}}
{{- $v := print "$DOMAINTYPESPATH/" $type.path ",\\"}}
{{- $_ := set $myDict $k $v}}
{{- end}}
{{- range $k,$v := $myDict}}
{{- $k}}={{$v}}
{{end -}}
Mgoogle/type/date.proto=google.golang.org/genproto/googleapis/type/date,\
Mgoogle/type/money.proto=google.golang.org/genproto/googleapis/type/money,\
Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
Mgoogle/protobuf/empty.proto=github.com/golang/protobuf/ptypes/empty,\
Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
Mgoogle/protobuf/field_mask.proto=google.golang.org/genproto/protobuf/field_mask,\
plugins=grpc:$TARGETDIR \
--grpc-gateway_out=logtostderr=true:$TARGETDIR/  $f

done