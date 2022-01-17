// @program:     tiny-stl
// @file:        modify_test.go.go
// @author:      edte
// @create:      2022-01-02 22:42
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

func TestCopy(t *testing.T) {

}

func TestFill(t *testing.T) {
	//v := vector.New(vector.WithCap(10))
	//Fill(v.Begin(), v.End(), 10)
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
