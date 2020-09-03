#!/usr/bin/env bash

DOMAINTYPESPATH={{.config.module}}/pkg/domaintypes
TARGETDIR="../../../pkg/domaintypes/"
# enable recursion for /**/*.xxx
shopt -s globstar dotglob

cd {{.config.build.proto.targettypedir}}

FILES=./**/*.proto
for f in $FILES
do
  echo "Processing $f file..."

protoc --proto_path=./ \
-I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
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
Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
Mgoogle/protobuf/empty.proto=github.com/golang/protobuf/ptypes/empty,\
Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
Mgoogle/protobuf/field_mask.proto=google.golang.org/genproto/protobuf/field_mask,\
:$TARGETDIR $f

done