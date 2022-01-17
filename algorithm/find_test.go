// @program:     tiny-stl
// @file:        find_test.go.go
// @author:      edte
// @create:      2022-01-17 00:57
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestCount(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{23, 2, 3, 3, 3, 3, 2, 2}))
	count := Count(v.CBegin(), v.CEnd(), 2)
	fmt.Println(count)
}

func TestCountIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{23, 2, 3, 3, 3, 3, 2, 2}))
	res := CountIf(v.CBegin(), v.CEnd(), func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	fmt.Println(res)
}

func TestEqual(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{25, 15, 5, -5, -15}))
	v1 := vector.New(vector.WithData([]interface{}{25, 15, 5, -3, -15}))

	equal := Equal(v.Begin(), v.End(), v1.Begin())
	fmt.Println(equal)
}
