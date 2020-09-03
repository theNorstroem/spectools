package transcoder

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"{{.config.module}}/pkg/grpc-gateway/autoregister"
	"golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"net/http"
)

// header für client ohne prefixes
func outgoingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

// header vom client im ctx des Servers ohne prefixes senden
func incomingMatcher(headerName string) (mdName string, ok bool) {
	return headerName, true
}

func Run(grpcBackendAddr string, transcoderAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithOutgoingHeaderMatcher(outgoingMatcher), runtime.WithIncomingHeaderMatcher(incomingMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := autoregister.RegisterAllEndpoints(grpcBackendAddr, ctx, mux, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(transcoderAddr, mux)
}
