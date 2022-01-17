// @program:     tiny-stl
// @file:        iterator_test.go.go
// @author:      edte
// @create:      2022-01-16 22:11
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestDistance(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	fmt.Println(Distance(v.Begin(), v.End()))
}
