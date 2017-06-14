//Package encoding encoding the error or response
package encoding

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"net/http"

	"go-httpframe/internal/errutil"

	"bytes"
)

type errorer interface {
	error() error
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(response)

	_, err := w.Write(buf.Bytes())
	return err
	//return json.NewEncoder(w).Encode(response)
}

func EncodeOctetStreamResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="orders.csv"`)

	if tables, ok := response.([][]string); !ok {
		return errutil.ErrWrongType
	} else {
		t := csv.NewWriter(w)
		t.WriteAll(tables)
		return t.Error()
	}

	return nil
}

func EncodePlainResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}

	var err error

	var data []byte
	switch response.(type) {
	case string:
		data = []byte(response.(string))

	case []byte:
		data = response.([]byte)

	default:
		return errutil.ErrIllegalParameter
	}

	_, err = w.Write(data)
	return err
}

func EncodeError(_ context.Context, e error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(encodeError(e))
}

func SimpleEncodeError(err error) interface{} {
	return encodeError(err)
}
