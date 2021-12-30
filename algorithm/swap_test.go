// @program:     tiny-stl
// @file:        swap_test.go.go
// @author:      edte
// @create:      2021-12-25 15:40
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/containers/vector"
)

func TestSwap(t *testing.T) {
	v := vector.New()
	v.PushBack(23)
	v.PushBack(9)
	v.PushBack(22)
	v.PushBack(-1)
	fmt.Println(v)
	Swap(v.Begin(), v.End().Pre())
	fmt.Println(v)
}
