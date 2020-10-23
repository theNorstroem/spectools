package transcoder

import (
	"context"
	"{{.config.module}}/dist/grpc-gateway/autoregister"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/ext"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"io/ioutil"
	"net/http"
	"os"
)

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

// header für client ohne prefixes
func outgoingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

// header vom client im ctx des Servers ohne prefixes senden
func incomingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

// setzt die api-base-url für das backend
func addBaseUrl (ctx context.Context, request *http.Request) metadata.MD {
	return metadata.New(map[string]string{
		"api-base-url": "//" + request.Host,
	})
}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || err == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

func Run(grpcBackendAddr string, transcoderAddr string) error {

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(outgoingMatcher),
		runtime.WithIncomingHeaderMatcher(incomingMatcher),
		runtime.WithMetadata(addBaseUrl),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}))

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := autoregister.RegisterAllEndpoints(grpcBackendAddr, ctx, mux, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(transcoderAddr, tracingWrapper(mux))
}
