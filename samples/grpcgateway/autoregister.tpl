package autoregister
{{$config := .config}}
import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	{{- range $servicename, $service:= .services}}
	{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	{{$pkgopts._1}} "{{$config.module}}/pkg/grpcservices/{{$pkgopts._0}}" //{{$servicename}}
    {{- end}}
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func RegisterAllEndpoints(grpcBackendAddr string, ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
	var err error
{{- range $servicename, $service:= .services}}
{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	// {{$servicename}}
	err = {{$pkgopts._1}}.Register{{$service.name}}HandlerFromEndpoint(ctx, mux, grpcBackendAddr, opts)
	if err != nil {
		return err
	}
{{end}}
	return err
}
