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
	"sort"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestSelectSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	selectSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestBubbleSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	// v := vector.New(vector.WithData([]interface{}{-1, 1, 2, 3, 3, 5, 6, 9, 99}))
	bubbleSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestInsertSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	insertSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestSellSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	shellSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestMergeSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	mergeSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestSleepSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	//sleepSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestBogoSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	bogoSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestCocktailSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	cocktailSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestQuickSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -234, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 3, -1, 3, 1, 1}))
	quickSort(v.Begin(), v.End(), 9999999)
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestHeapSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -234, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 3, -1, 3, 1, 1}))
	heapSort(v.Begin(), v.End())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestCountSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{3, 2, 4, 1, 1}))
	countSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestBucketSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{13, 2, 14, 1, 1}))
	bucketSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End()))
}

func TestRadioSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1, 134, 22, -234, -3, 34, 5, 1, 623, -24, 555, 3, 7, -2}))
	// v := vector.New(vector.WithData([]interface{}{5, 2, 99, -1, 134, 22, 1, 3, 34, 5, 1, 623, 6, 555, 3, 7, 9, 3}))
	// v := vector.New(vector.WithData([]interface{}{23, 22, 14, 13, 1, 6, 0}))
	radioSort(v.Begin(), v.End(), comparator.NewLess())
	assert.True(t, IsSorted(v.Begin(), v.End()))
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

	bubbleSort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		return a.(user).age > b.(user).age
	}))

	fmt.Println(v)

	// selectSort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
	// 	if a.(user).age != b.(user).age {
	// 		return a.(user).age > b.(user).age
	// 	}
	// 	return a.(user).name < b.(user).name
	// }))
	//
	// fmt.Println(v)
}

func get() *vector.Vector {
	v := vector.New()
	for i := 0; i < 1e2; i++ {
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
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		bubbleSort(v.Begin(), v.End())
	}
}

func BenchmarkCocktailSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		cocktailSort(v.Begin(), v.End())
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		insertSort(v.Begin(), v.End())
	}
}

func BenchmarkSellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		shellSort(v.Begin(), v.End())
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		mergeSort(v.Begin(), v.End())
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		quickSort(v.Begin(), v.End(), 9999999)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		heapSort(v.Begin(), v.End())
		// fmt.Println(IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		countSort(v.Begin(), v.End())
		// fmt.Println(IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		bucketSort(v.Begin(), v.End())
		// fmt.Println(IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkRadioSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		radioSort(v.Begin(), v.End())
		// fmt.Println(IsSorted(v.Begin(), v.End()))
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		Sort(v.Begin(), v.End())
	}
}

func BenchmarkStdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v := get()
		b.StartTimer()
		sort.Sort(v)
		// Sort(v.Begin(), v.End())
		// fmt.Println(IsSorted(v.Begin(), v.End()))
	}
}

func TestIsSorted(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 3}))
	fmt.Println(IsSorted(v.Begin(), v.End()))
	shellSort(v.Begin(), v.End())
	fmt.Println(v)
	fmt.Println(IsSorted(v.Begin(), v.End()))
}

func TestIsSortedUntil(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 3, 6, 2, 1, 5, 2, 3}))
	d := IsSortedUntil(v.Begin(), v.End())
	fmt.Println(d.Value())
}

func TestGetDepth(t *testing.T) {
	v := vector.New(vector.WithCapInit(1e9, 1))
	depth := getDepth(v.Begin(), v.End())
	fmt.Println(depth)
}

func TestPartialSort(t *testing.T) {
	// v := vector.New(vector.WithData([]interface{}{69, 23, 80, 42, 17, 15, 26, 51, 19, 12, 45, 72}))
	// PartialSort(v.Begin(), v.Begin().NextN(7), v.End())
	// fmt.Println(v)

	v := vector.New(vector.WithData([]interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	PartialSort(v.Begin(), v.Begin().NextN(5), v.End(), comparator.NewLess())
	fmt.Println(v)
}

func TestPartialSortCopy(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1}))
	res := vector.New(vector.WithCap(5))

	PartialSortCopy(v.Begin(), v.End(), res.Begin(), res.End())

	fmt.Println(res)
}
