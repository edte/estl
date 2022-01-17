// @program:     tiny-stl
// @file:        copy_test.go.go
// @author:      edte
// @create:      2022-01-17 01:14
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestCopyIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{25, 15, 5, -5, -15}))
	res := vector.New(vector.WithCap(v.Size()))

	CopyIf(v.CBegin(), v.CEnd(), res.OBegin(), func(i interface{}) bool {
		return i.(int) > 0
	})

	ForEach(res.Begin(), res.End(), func(val interface{}) {
		if val == nil {
			return
		}
		fmt.Print(val)
		fmt.Print(" ")
	})
}
