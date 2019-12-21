package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	gw "github.com/mmmknt/go-grpc-gateway-example/service"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

func responseHeaderFilter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Del("Grpc-Metadata-Content-Type")
	return nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	runtime.HTTPError = customHTTPError
	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(responseHeaderFilter))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

type errorBody struct {
	Err string `json:"error,omitempty"`
}

func customHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(grpc.Code(err)))
	jErr := json.NewEncoder(w).Encode(errorBody{
		Err: status.Convert(err).Message(),
	})

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
