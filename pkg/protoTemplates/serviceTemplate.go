package protoTemplates

var ServiceTemplate = `// Code generated by furo-proto-gen. DO NOT EDIT.

syntax = "proto3";{{$GenAdditionalBinding := .GenAdditionalBinding}}
package {{.Package}};
{{- range $key, $option := .Options}}
option {{$key}} = {{if  ($option | eq "true")}}{{$option}};{{else}}"{{$option}}";{{end}}{{end}}
{{range $import := .Imports}}
import "{{$import}}";{{end}}
{{range $Service := .Services}}
{{$serviceName := camelcase (snakecase $Service.Name)}}
{{if $Service.Description}}// {{$Service.Description | replace "\n" "\n// "}}{{end}}
service {{$serviceName}} { 
{{- $rpcmap := $Service.Services | rpcmap}}
{{range $rpckey, $method := $rpcmap}}
  {{if $method.Description}}// {{$method.Description | replace "\n" "\n// " | noescape}}{{end}}
  rpc {{if $method.RpcName}}{{$method.RpcName}}{{else}}{{$rpckey}}{{$serviceName}}{{end}} ({{$rpckey}}{{$serviceName}}Request) returns ({{$method.Data.Response}}){
	{{if $method.Deeplink.Description}}//{{$method.Deeplink.Description | replace "\n" "\n// "}}{{else}}// developer was to lazy to describe the rpc, sorry{{end}}
	option (google.api.http) = {
		{{ lower $method.Deeplink.Method}}: "{{$method.Deeplink.Href}}"{{ if $method.Data.Request}}
		{{ if or (eq $method.Deeplink.Method "POST") (eq $method.Deeplink.Method "PATCH") (eq $method.Deeplink.Method "PUT")}}body: "{{$method.Data.BodyField}}"{{end}}{{end}}
		{{- if and (eq $method.Deeplink.Method "PUT") (eq $GenAdditionalBinding true)  }}
		additional_bindings {
            patch: "{{$method.Deeplink.Href}}"
            body: "{{$method.Data.BodyField}}"
        }{{end}}
	};
  }
{{end}}
}
{{end}}

`
