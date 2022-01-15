// @program:     tiny-stl
// @file:        vector_test.go.go
// @author:      edte
// @create:      2021-12-30 13:22
// @description:
package vector

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	v := New()
	v.PushBack(3)
	v.PushBack(4)
	v.PushBack(5)
	v.PushBack(6)

	for b := v.Begin(); !b.Equal(v.End()); b.Pre() {
		fmt.Println(b.Value())
	}
}

func TestVector_Swap(t *testing.T) {
	v := New(WithData([]interface{}{1, 2, 3, 4, 5}))
	v2 := New(WithData([]interface{}{3, 6, 1, 2, 7, 3, 7}))
	fmt.Println(v)
	fmt.Println(v2)

	v.Swap(v2)

	fmt.Println(v)
	fmt.Println(v2)
}
