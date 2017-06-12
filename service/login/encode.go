package login

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeLoginResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
