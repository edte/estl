// @program:     tiny-stl
// @file:        heap_test.go.go
// @author:      edte
// @create:      2022-01-15 01:23
// @description:
package algorithm

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestMakeHeap(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{7, 2, 1, 4, 5, 3, 6}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))
	MakeHeap(v.Begin(), v.End(), comparator.NewELess())
	fmt.Println(ShowHeap(v.Begin(), v.End()))
}

func TestPushHeap(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))
	v.PushBack(0)
	PushHeap(v.Begin(), v.End())
	fmt.Println(ShowHeap(v.Begin(), v.End()))
}

func TestPopHeap(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))
	PopHeap(v.Begin(), v.End())
	fmt.Println(ShowHeap(v.Begin(), v.End()))
	v.PopBack()
	fmt.Println(ShowHeap(v.Begin(), v.End()))
}

func TestIsHeap(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(IsHeap(v.Begin(), v.End()))
	v1 := vector.New(vector.WithData([]interface{}{1, 8, 3, 3, 5, 6, 7, 8}))
	fmt.Println(IsHeap(v1.Begin(), v1.End()))
}

func TestIsHeap2(t *testing.T) {
	v := vector.New()

	for i := 0; i < 1e5; i++ {
		if i%3 == 0 || i%7 == 0 {
			func() {
				n, _ := rand.Int(rand.Reader, big.NewInt(100))
				v.PushBack(n.Int64())
			}()
		}
	}

	fmt.Println(ShowHeap(v.Begin(), v.End()))

	MakeHeap(v.Begin(), v.End())

	fmt.Println(ShowHeap(v.Begin(), v.End()))

	fmt.Println(IsHeap(v.Begin(), v.End()))
}

func TestIsHeapUntil(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 8, 3, 4, 5, 6, 7, 8}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))
	fmt.Println(IsHeapUntil(v.Begin(), v.End()))
}

func TestGetParent(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{"j", "a", "b", "c", "d", "e", "f", "g", "h", "i"}))
	parent := getParent(v.Begin().Next(), v.Begin().Clone().NextN(4))
	fmt.Println(parent.Value())
}

func TestLeftChild(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{"j", "a", "b", "c", "d", "e", "f", "g", "h", "i"}))
	parent := getLeftChild(v.Begin().Next(), v.Begin().Clone().NextN(4))
	fmt.Println(parent.Value())
}

func TestRightChild(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{"j", "a", "b", "c", "d", "e", "f", "g", "h", "i"}))
	parent := getRightChild(v.Begin().Next(), v.Begin().Clone().NextN(4))
	fmt.Println(parent.Value())
}

func TestShow(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{"asdf", "asdf", "j", "a", "b", "c", "d", "e", "f", "g", "h", "i"}))
	root := v.Begin().NextN(3)

	fmt.Println(ShowHeap(root, v.End()))

	// fmt.Println(root)
	//
	// b := getLeftChild(root, root)
	// fmt.Println(b)
	//
	// c := getRightChild(root, root)
	// fmt.Println(c)
	//
	// fmt.Println(getLeftChild(root, b))
	// fmt.Println(getRightChild(root, b))
	//
	// fmt.Println(getLeftChild(root, c))
	// fmt.Println(getRightChild(root, c))
}

func TestUp(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 0}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))

	siftUp(v.Begin(), v.End().Pre(), comparator.NewGreater())

	fmt.Println()
	fmt.Println()

	fmt.Println(ShowHeap(v.Begin(), v.End()))
}

func TestDown(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{5, 3, 4, 2, 1}))
	fmt.Println(ShowHeap(v.Begin(), v.End()))

	siftDown(v.Begin(), v.End(), v.Begin(), comparator.NewGreater())

	fmt.Println(ShowHeap(v.Begin(), v.End()))
}

func TestGetLevel(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(getLevel(v.Begin(), v.End()))
}

func TestSortHeap(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 8, -1, 88, 1, 3, 4, 5, 6, 7, 8, 8, 2, 6, 1, 3, 1}))
	// fmt.Println(ShowHeap(v.Begin(), v.End()))
	MakeHeap(v.Begin(), v.End())
	// fmt.Println(ShowHeap(v.Begin(), v.End()))
	SortHeap(v.Begin(), v.End())
	// fmt.Println(ShowHeap(v.Begin(), v.End()))

	ForEach(v.Begin(), v.End(), func(val interface{}) {
		fmt.Print(val)
		fmt.Print(" ")
	})
}
