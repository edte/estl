// @program:     tiny-stl
// @file:        foreach_test.go.go
// @author:      edte
// @create:      2021-12-30 15:49
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestForEach(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	ForEach(v.Begin(), v.End(), func(val interface{}) {
		fmt.Println(val)
	})

}
