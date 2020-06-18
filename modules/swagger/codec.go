package swagger

import (
	"context"
	"encoding/json"
	"net/http"
)

// HTTP返回数据编码函数
func HttpEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With,authToken")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	return json.NewEncoder(w).Encode(response)

}
