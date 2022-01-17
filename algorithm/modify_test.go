// @program:     tiny-stl
// @file:        modify_test.go.go
// @author:      edte
// @create:      2022-01-02 22:42
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestFill(t *testing.T) {
	v := vector.New(vector.WithCap(10))
	Fill(v.Begin(), v.End(), 10)
	assert.Equal(t, "10 10 10 10 10 10 10 10 10 10 ", v.String())
	fmt.Println(v)
}

func TestFillN(t *testing.T) {
	v := vector.New(vector.WithCap(6))
	FillN(v.Begin(), 3, 0)
	FillN(NextN(v.Begin(), 3), 3, 1)
	assert.Equal(t, "0 0 0 1 1 1 ", v.String())
	fmt.Println(v)
}

func TestShuffle(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
	Shuffle(v.Begin(), v.End())
	fmt.Println(v)
}

func TestReverse(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	fmt.Println(v)
	Reverse(v.Begin(), v.End())
	fmt.Println(v)
}
