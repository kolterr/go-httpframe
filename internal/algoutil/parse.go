package algoutil

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"wukong/internal/errutil"
)

func RetrieveOrDefault(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}

	return i
}

func Retrieve64OrDefault(s string, d int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return d
	}

	return i
}

func ParseID(r *http.Request, key ...string) (int64, error) {
	var k string
	if len(key) == 0 {
		k = "id"
	} else {
		k = key[0]
	}
	vars := mux.Vars(r)
	id, ok := vars[k]
	if !ok {
		return 0, errutil.ErrParameterMissing
	}

	id = strings.TrimSpace(id)
	if id == "" {
		return 0, errutil.ErrIllegalParameter
	}

	return strconv.ParseInt(id, 10, 0)

}

func ExtractParams(r *http.Request, keys ...string) (map[string]string, error) {
	vars := mux.Vars(r)

	m := make(map[string]string)
	for _, key := range keys {
		val, ok := vars[key]
		if !ok {
			return nil, errutil.ErrParameterMissing
		}

		val = strings.TrimSpace(val)
		if val == "" {
			return nil, errutil.ErrIllegalParameter
		}

		m[key] = val
	}
	return m, nil

}
