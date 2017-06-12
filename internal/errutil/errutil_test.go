package errutil

import (
	"fmt"
	"sort"
	"testing"
)

func TestCode2(t *testing.T) {
	keys := []int{}
	kv := make(map[int]error)
	for e, c := range code {
		keys = append(keys, c)
		kv[c] = e
	}
	sort.Ints(keys)
	for _, key := range keys {
		fmt.Printf("%4d - %s\n", key, kv[key].Error())
	}
}
