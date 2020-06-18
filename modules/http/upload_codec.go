package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func UploadDecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	r.ParseMultipartForm(32 << 20)

	return nil, nil
}

func UploadEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)

}
