package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"snakesdk/cmd/snaked/protocol"

	"snakepop.com/armory/errutil"
)

func BuildDSN(username, password, host string, port int, dbname, args string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, dbname, args)
}

func EncodeError(_ context.Context, e error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(encodeError(e))
}

func encodeError(e error) interface{} {
	var response protocol.ErrorResponse
	var (
		code = errutil.Code(e)
		err  = e.Error()
	)
	//if raw, ok := e.(kithttp.Error); ok {
	//	code = errutil.Code(raw.Err)
	//	err = raw.Err.Error()
	//}

	if code == errutil.Unknown {
		err = errutil.ErrServerInternal.Error()
	}
	response = protocol.ErrorResponse{Code: code, Error: err}

	return response
}
