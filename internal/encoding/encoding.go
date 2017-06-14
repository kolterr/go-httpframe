// +build !debug

package encoding

import (
	"go-httpframe/internal/errutil"

	"go-httpframe/protocol"
)

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
