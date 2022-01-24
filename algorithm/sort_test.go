// @program:     tiny-stl
// @file:        sort_test.go.go
// @author:      edte
// @create:      2021-12-30 17:19
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"stl/containers/stack"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestSelectSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))

	selectSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	selectSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestBubbleSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	// v := vector.New(vector.WithData([]interface{}{-1, 1, 2, 3, 3, 5, 6, 9, 99}))

	bubbleSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	bubbleSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestInsertSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))

	insertSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	insertSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestSellSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))

	shellSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	shellSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestMergeSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	//v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1}))
	mergeSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	mergeSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestSleepSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	sleepSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	sleepSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestBogoSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	bogoSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	bogoSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestCocktailSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	cocktailSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	cocktailSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestQuickSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -234, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 3, -1, 3, 1, 1}))
	quickSort(v.Begin(), v.End(), 9999999)
	assert.True(t, IsSorted(v.Begin(), v.End()))
	quickSort(v.Begin(), v.End(), 9999999, comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestHeapSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -234, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 3, -1, 3, 1, 1}))
	heapSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	heapSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestCountSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 4, 1, 1}))
	countSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	countSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestBucketSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{13, 2, 14, 1, 1}))
	bucketSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	bucketSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestRadioSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{23, 22, 14, 13, 1, 6, 0}))
	radioSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	radioSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestTimSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{8, 10, 15, 6, 18, 12, 7, 5, 4, 3, 20, 8, 11, 22, 9, 9, 17, 18, 2, 1, 19, 16, 14, 25, 24}))
	timSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
	timSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func Test_minRunLength(t *testing.T) {
	assert.Equal(t, 5, minRunLength(5))
	assert.Equal(t, 7, minRunLength(7))
	assert.Equal(t, 17, minRunLength(33))
	assert.Equal(t, 23, minRunLength(45))
	assert.Equal(t, 16, minRunLength(64))
	assert.Equal(t, 16, minRunLength(128))
	assert.Equal(t, 16, minRunLength(1024))
}

func Test_getOrderlyLength(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 3, 5, 2}))
	assert.Equal(t, 3, getOrderlyLength(v.Begin(), v.End()))

	v1 := vector.New(vector.WithData([]interface{}{1, 3, 5, 5, 2}))
	assert.Equal(t, 3, getOrderlyLength(v1.Begin(), v1.End()))

	v1 = vector.New(vector.WithData([]interface{}{1, 3, 5, 5, 2, 1, 0}))
	assert.Equal(t, 3, getOrderlyLength(v1.Begin(), v1.End()))

	v1 = vector.New(vector.WithData([]interface{}{1}))
	assert.Equal(t, 1, getOrderlyLength(v1.Begin(), v1.End()))

	v1 = vector.New(vector.WithData([]interface{}{}))
	assert.Equal(t, 0, getOrderlyLength(v1.Begin(), v1.End()))

	v1 = vector.New(vector.WithData([]interface{}{1, 2}))
	assert.Equal(t, 2, getOrderlyLength(v1.Begin(), v1.End()))

	v1 = vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5}))
	assert.Equal(t, 5, getOrderlyLength(v1.Begin(), v1.End()))

	v = vector.New(vector.WithData([]interface{}{8, 10, 15, 6, 18, 12, 7, 5, 4, 3, 20, 8, 11, 22, 9, 9, 17, 18, 2, 1, 19, 16, 14, 25, 24}))
	assert.Equal(t, 3, getOrderlyLength(v.Begin(), v.End()))
}

func Test_mergeCollapse(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{6, 8, 10, 15, 3, 4, 5, 7, 12, 18}))
	s := stack.New()

	s.Push(MakePair(v.Begin(), NextN(v.Begin(), 4)))
	s.Push(MakePair(NextN(v.Begin(), 4), v.End()))

	mergeCollapse(s)

	assert.Equal(t, "3 4 5 6 7 8 10 12 15 18 ", v.String())
}

func Test_binaryInsertSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))

	binaryInsertSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	binaryInsertSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func TestSort(t *testing.T) {
	v := get()

	Sort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	Sort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func Test_permutationSort(t *testing.T) {
	//v := get()
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))

	permutationSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))

	permutationSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End(), comparator.NewLess()))
}

func Test_insert(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 3, 5, 6, 4}))
	insert(v.Begin(), v.End())
	assert.Equal(t, "1 3 4 5 6 ", v.String())

	v1 := vector.New(vector.WithData([]interface{}{6, 5, 3, 1, 4}))
	insert(v1.Begin(), v1.End(), comparator.NewLess())
	assert.Equal(t, "6 5 4 3 1 ", v1.String())

	v2 := vector.New(vector.WithData([]interface{}{2, 1}))
	insert(v2.Begin(), v2.End())
	assert.Equal(t, "1 2 ", v2.String())

	v3 := vector.New(vector.WithData([]interface{}{2, 1}))
	insert(v3.Begin(), v3.End(), comparator.NewLess())
	assert.Equal(t, "2 1 ", v3.String())
}

func TestName(t *testing.T) {
	v := get()
	// fmt.Println(v)
	// fmt.Println()
	quickSort(v.Begin(), v.End(), 9999999)
	// fmt.Println(v)
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestSortCmp(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	selectSort(v.Begin(), v.End(), comparator.NewGreater())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestSortUser(t *testing.T) {
	type user struct {
		age  int
		name string
	}

	v := vector.New(vector.WithData([]interface{}{
		user{
			age:  17,
			name: "a",
		},
		user{
			age:  20,
			name: "b",
		},
		user{
			age:  20,
			name: "a",
		},
		user{
			age:  16,
			name: "c",
		},
		user{
			age:  17,
			name: "c",
		},
		user{
			age:  16,
			name: "e",
		},
	}))

	StableSort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		return a.(user).age > b.(user).age
	}))

	assert.Equal(t, "{16 c} {16 e} {17 a} {17 c} {20 b} {20 a} ", fmt.Sprint(v))

	Sort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		if a.(user).age != b.(user).age {
			return a.(user).age > b.(user).age
		}
		return a.(user).name < b.(user).name
	}))

	assert.Equal(t, "{16 e} {16 c} {17 c} {17 a} {20 b} {20 a} ", fmt.Sprint(v))
}

func get() *vector.Vector {
	v := vector.New()
	for i := 0; i < 1e5; i++ {
		v.PushBack(rand.Int() % 100)
		// v.PushBack(i)
	}
	return v
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		selectSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		bubbleSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkCocktailSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		cocktailSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		insertSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkSellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		shellSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		mergeSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		quickSort(v.Begin(), v.End(), 9999999)
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		heapSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		countSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		bucketSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkRadioSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		radioSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkBinaryInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		binaryInsertSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkTimSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		timSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		Sort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkStableSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		StableSort(v.Begin(), v.End())
		assert.True(b, IsSorted(v.Begin(), v.End()))
	}
}

//func BenchmarkStdSort(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		b.StopTimer()
//		//v := get()
//		b.StartTimer()
//		//sort.Sort(v)
//		// Sort(v.Begin(), v.End())
//		assert.True(b, IsSorted(v.Begin(), v.End()))
//	}
//}

func TestIsSorted(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	assert.False(t, IsSorted(v.Begin(), v.End()))
	shellSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestIsSortedUntil(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 3, 6, 2, 1, 5, 2, 3}))
	d := IsSortedUntil(v.Begin(), v.End())
	assert.Equal(t, 2, d.Value())
}

func TestGetDepth(t *testing.T) {
	v := vector.New(vector.WithCapInit(1e3, 1))
	depth := getDepth(v.Begin(), v.End())
	assert.Equal(t, 20, depth)
}

func TestPartialSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{69, 23, 80, 42, 17, 15, 26, 51, 19, 12, 45, 72}))
	PartialSort(v.Begin(), v.Begin().NextN(7), v.End())
	assert.Equal(t, "12 15 17 19 23 26 42 80 69 51 45 72 ", v.String())

	v = vector.New(vector.WithData([]interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	PartialSort(v.Begin(), v.Begin().NextN(5), v.End(), comparator.NewLess())
	assert.Equal(t, "9 8 7 6 5 4 3 2 1 ", v.String())
}

func TestPartialSortCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	res := vector.New(vector.WithCap(5))

	PartialSortCopy(v.Begin(), v.End(), res.Begin(), res.End())

	assert.Equal(t, "1 2 3 4 5 ", res.String())
}

func TestNthElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	for i := 0; i < 1e2; i++ {
		Shuffle(v.Begin(), v.End())
		NthElement(v.Begin(), NextN(v.Begin(), 5), v.End())
		assert.Equal(t, 6, v.At(5))
		Shuffle(v.Begin(), v.End())
		NthElement(v.Begin(), NextN(v.Begin(), 5), v.End(), comparator.NewLess())
		assert.Equal(t, 4, v.At(5))
	}
}
