package gateway

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
)

type errorBody struct {
	Error string `json:"error,omitempty"`
	Code  int    `json:"code,omitempty"`
}

func HttpError(ctx context.Context, _ *runtime.ServeMux, marshaller runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", marshaller.ContentType(nil))
	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))
	jErr := json.NewEncoder(w).Encode(
		errorBody{
			Error: status.Convert(err).Message(),
			Code:  runtime.HTTPStatusFromCode(status.Code(err)),
		},
	)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
