// @program:     tiny-stl
// @file:        iterator_test.go.go
// @author:      edte
// @create:      2022-01-15 17:19
// @description:
package vector

import (
	"fmt"
	"testing"
)

func TestIterator_String(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4}))
	fmt.Println(v.Begin().Next())
}

func TestVector_String(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4}))
	i := NewIterator(v, 0)
	fmt.Println(i.String())
}
