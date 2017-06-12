package util

import (
	"fmt"
)

func BuildDSN(username, password, host string, port int, dbname, args string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, dbname, args)
}
