package util

import (
	"go-httpframe/internal/errutil"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"runtime"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//CallSite the caller's file & line
func CallSite() interface{} {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return string(file + ":" + strconv.FormatInt(int64(line), 10))
}

func Utf8ToGBK(utf8str string) string {
	result, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), utf8str)
	return result
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		//if r.Method == http.MethodOptions {
		//	json.NewEncoder(w).Encode(protocol.SuccessResponse)
		//	return
		//}

		h.ServeHTTP(w, r)
	})
}

func JsonAPI(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	})
}

//HTTPGet http's get method
func HTTPGet(url string) (string, error) {
	var body []byte
	rspn, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer rspn.Body.Close()
	body, err = ioutil.ReadAll(rspn.Body)

	return string(body), err
}

func CopyFile(dst, src string) error {
	if dst == "" || src == "" {
		return errutil.ErrIllegalParameter
	}

	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	return err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func MakeDirIfNeed(dir string) error {
	dir = strings.TrimRight(dir, "/")

	if FileExists(dir) {
		return nil
	}

	err := os.MkdirAll(dir, os.ModePerm)
	return err
}
