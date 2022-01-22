// @program:     tiny-stl
// @file:        modify_test.go.go
// @author:      edte
// @create:      2022-01-02 22:42
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"stl/functional"
	"stl/iterator"
	"testing"

	"stl/containers/vector"
)

func TestCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	res := vector.New(vector.WithCap(v.Size()))

	Copy(v.Begin(), v.End(), res.Begin())

	assert.Equal(t, res.String(), "9 3 6 3 1 5 2 3 ")

}

func TestCopyN(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	res := vector.New(vector.WithCap(3))
	CopyN(v.Begin(), 3, res.Begin())

	assert.Equal(t, res.String(), "9 3 6 ")

}

func TestCopyIf(t *testing.T) {
	op := func(i interface{}) bool {
		return i.(int) > 0
	}

	v := vector.New(vector.WithData([]interface{}{-3, 25, 15, 5, -5, -15}))
	res := vector.New(vector.WithCap(CountIf(v.Begin(), v.End(), op)))

	CopyIf(v.Begin(), v.End(), res.Begin(), op)

	assert.Equal(t, res.String(), "25 15 5 ")

	ForEach(res.Begin(), res.End())
}

func TestCopyBackward(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{-3, 25, 15, 5, -5, -15}))
	res := vector.New(vector.WithCap(v.Size()))

	CopyBackward(v.Begin(), v.End(), res.End())

	assert.Equal(t, "-3 25 15 5 -5 -15 ", res.String())
}

func TestTransform(t *testing.T) {
	v1 := vector.New(vector.WithData([]interface{}{10, 20, 30, 40, 50}))
	v2 := vector.New(vector.WithCap(v1.Size()))

	Transform(v1.Begin(), v1.End(), v2.Begin(), functional.Increase)

	assert.Equal(t, "11 21 31 41 51 ", v2.String())
}

func TestSwap(t *testing.T) {
	v := vector.New()
	v.PushBack(23)
	v.PushBack(9)
	v.PushBack(22)
	v.PushBack(-1)

	assert.Equal(t, "23 9 22 -1 ", v.String())

	Swap(v.Begin(), v.End().Pre())

	assert.Equal(t, "-1 9 22 23 ", v.String())
}

func TestSwapRanges(t *testing.T) {
	v1 := vector.New(vector.WithCapInit(5, 10))
	v2 := vector.New(vector.WithCapInit(5, 33))

	SwapRanges(Next(v1.Begin()), Pre(v1.End()), v2.Begin())

	assert.Equal(t, "10 33 33 33 10 ", v1.String())
	assert.Equal(t, "10 10 10 33 33 ", v2.String())
}

func TestFill(t *testing.T) {
	v := vector.New(vector.WithCap(10))
	Fill(v.Begin(), v.End(), 10)
	assert.Equal(t, "10 10 10 10 10 10 10 10 10 10 ", v.String())
}

func TestFillN(t *testing.T) {
	v := vector.New(vector.WithCap(6))
	FillN(v.Begin(), 3, 0)
	FillN(NextN(v.Begin(), 3), 3, 1)
	assert.Equal(t, "0 0 0 1 1 1 ", v.String())
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

	assert.Equal(t, "9 3 6 3 1 5 2 3 ", v.String())

	Reverse(v.Begin(), v.End())

	assert.Equal(t, "3 2 5 1 3 6 3 9 ", v.String())
}

func TestReverseCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	out := vector.New(vector.WithCap(v.Size()))
	ReverseCopy(v.Begin(), v.End(), out.Begin())
	assert.Equal(t, "3 2 5 1 3 6 3 9 ", out.String())
}

func TestGenerate(t *testing.T) {
	v := vector.New(vector.WithCap(10))
	Generate(v.Begin(), v.End(), functional.RandomNum)
	fmt.Println(v)
}

func TestGenerateN(t *testing.T) {
	v := vector.New(vector.WithCap(10))
	GenerateN(v.Begin(), 3, functional.RandomNum)
	fmt.Println(v)
}

func TestRotate(t *testing.T) {
	v := vector.New(vector.WithCap(9))
	i := 0
	Generate(v.Begin(), v.End(), func() interface{} {
		i++
		return i
	})

	Rotate(v.Begin(), NextN(v.Begin(), 3), v.End())

	assert.Equal(t, "4 5 6 7 8 9 1 2 3 ", v.String())
}

func TestRotateCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 40, 50, 60, 70}))
	res := vector.New(vector.WithCap(7))
	RotateCopy(v.Begin(), NextN(v.Begin(), 3), v.End(), res.Begin())
	assert.Equal(t, "40 50 60 70 10 20 30 ", res.String())
}

func TestReplace(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	Replace(v.Begin(), v.End(), 20, 99)
	assert.Equal(t, "10 99 30 30 99 10 10 99 ", v.String())
}

func TestReplaceIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6}))
	ReplaceIf(v.Begin(), v.End(), functional.IsOdd, 99)
	assert.Equal(t, "99 2 99 4 99 6 ", v.String())
}

func TestReplaceCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	res := vector.New(vector.WithCap(v.Size()))
	ReplaceCopy(v.Begin(), v.End(), res.Begin(), 20, 99)
	assert.Equal(t, "10 99 30 30 99 10 10 99 ", res.String())
}

func TestReplaceCopyIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6}))
	res := vector.New(vector.WithCap(v.Size()))
	ReplaceCopyIf(v.Begin(), v.End(), res.Begin(), functional.IsOdd, 99)
	assert.Equal(t, "99 2 99 4 99 6 ", res.String())
}

func TestRemoveCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 30, 30, 20, 10, 10, 20}))
	res := vector.New(vector.WithCap(v.Size()))
	RemoveCopy(v.Begin(), v.End(), res.Begin(), 20)
	assert.Equal(t, "10 30 30 10 10 <nil> <nil> <nil> ", res.String())
}

func TestRemoveCopyIf(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6}))
	res := vector.New(vector.WithCap(v.Size()))
	RemoveCopyIf(v.Begin(), v.End(), res.Begin(), functional.IsOdd)
	assert.Equal(t, "2 4 6 <nil> <nil> <nil> ", res.String())
}

func TestUnique(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 20, 20, 30, 30, 20, 20, 10}))

	res := Unique(v.Begin(), v.End())

	v2 := vector.New(vector.WithIter(v.Begin(), res.(iterator.RandomAccessIterator)))

	assert.Equal(t, "10 20 30 20 10 ", v2.String())
}

func TestUniqueCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{10, 20, 20, 20, 30, 30, 20, 20, 10}))
	res := vector.New(vector.WithCap(5))

	UniqueCopy(v.Begin(), v.End(), res.Begin())

	assert.Equal(t, "10 20 30 20 10 ", res.String())
}
