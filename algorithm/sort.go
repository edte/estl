// @program:     tiny-stl
// @file:        sort.go
// @author:      edte
// @create:      2021-12-30 16:07
// @description:
package algorithm

import (
	"fmt"
	"sync"
	"time"

	"stl/comparator"
	"stl/containers/vector"
	"stl/iterator"
)

// 随机访问迭代器的排序实现
// 主要是数组的排序

// selectSort 选择排序
// 时间复杂度 O(N^2)，空间复杂度 O(1)
// 不稳定排序
func selectSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	for i := first.Clone(); !i.Equal(last); i.Next() {
		for j := i.Clone().Next(); !j.Equal(last); j.Next() {
			if c.Operator(i.Value(), j.Value()) {
				Swap(i, j)
			}
		}
	}
}

// bubbleSort 冒泡排序
// 时间复杂度 O(N^2)，空间复杂度 O(1)
// 默认稳定排序
// 重写比较器时，如果相等时不交换，那么就是稳定的，否则不稳定
func bubbleSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	cnt := 1
	isSorted := true

	for i := first.Clone(); !i.Equal(last); i.Next() {
		for j := first.Clone(); !j.Equal(last.Clone().PreN(cnt)); j.Next() {
			if c.Operator(j.Value(), j.ValueAt(j.Position()+1)) {
				isSorted = false
				Swap(j, j.Clone().Next())
			}
		}
		// 优化，如果一趟下来没有交换顺序，说明已经有序了
		if isSorted {
			return
		}
		cnt++
	}
}

// CocktailSort 鸡尾酒排序,冒泡排序的优化
// 就是对冒泡排序增加一个方向的冒法
// 时间复杂度 O(N^2)，空间复杂度 O(1)
func cocktailSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	cnt := 0
	mid := first.Clone().NextN(first.Distance(last) / 2)
	isSorted := true

	for i := first.Clone(); !i.Equal(mid); i.Next() {
		for j := first.Clone().NextN(cnt); !j.Equal(last.Clone().PreN(cnt + 1)); j.Next() {
			if c.Operator(j.Value(), j.ValueAt(j.Position()+1)) {
				Swap(j, j.Clone().Next())
				isSorted = false
			}
		}
		if isSorted {
			break
		}

		for j := last.Clone().PreN(cnt + 2); !j.Equal(first.Clone().NextN(cnt)); j.Pre() {
			if c.Operator(j.ValueAt(j.Position()-1), j.Value()) {
				Swap(j, j.Clone().Pre())
				isSorted = false
			}
		}

		if isSorted {
			break
		}
	}
}

// insertSort 插入排序
// 插入排序看上去似乎可以二分优化，不过数组又需要移动元素，
// 复杂度还是 O (n^2)，链表倒是不用移动元素，但是又不能二分。
// 由于涉及元素移动，故从后往前查找是比较方便的，边查找边交换元素
// 故时间复杂度 O(N^2)，空间复杂度 O(1)
// 默认稳定排序
func insertSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	for i := first.Clone().Next(); !i.Equal(last); i.Next() {
		j := i.Clone().Pre()
		t := i.Value()

		for j.HasPre() && c.Operator(j.Value(), t) {
			j.Clone().Next().SetValue(j.Value())
			j.Pre()
		}

		j.Next().SetValue(t)
	}
}

// 希尔排序
// 希尔排序其实就是增量+插入，就是插入排序外面套一层增量循环
// 增量减少规模+部分排序，充分利用插入排序的性质
// 插入排序的部分就是插入的间隔从 1 变成了增量，其它都没变
// Knuth 增量：<O(n^(3/2)) by Knuth,1973>: 1, 4, 13, 40, 121, ...
// 时间复杂度 O(nlog^n)，空间复杂度 O(1)
// 不稳定排序
func shellSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	l := first.Distance(last)

	// 计算初始增量
	h := 1
	for h < l/3 {
		h = h*3 + 1
	}

	// 增量每次缩小三倍
	for ; h >= 1; h /= 3 {
		// 对每个分组进行排序，需要注意的是，不是一个分组排完才排另一个，而是每个分组轮着来
		// 类似于普通的插入排序，前 h 个不用排，即每个分组第一个不用排
		for i := first.Clone().NextN(h); !i.Equal(last); i.Next() {
			j := i.Clone().PreN(h)
			t := i.Value()

			for j.IsValid() && c.Operator(j.Value(), t) {
				j.Clone().NextN(h).SetValue(j.Value())
				j.PreN(h)
			}

			j.NextN(h).SetValue(t)
		}
	}
}

// 归并排序
// 归并有递归实现，迭代实现，原地合并等多种方式
// 通常最经典和常见的实现是 递归+额外数组实现，而非递归和原地在面试中可能会遇见
// 故时间复杂度 O(N*log^n)，空间复杂度 O(n)
// 默认稳定排序
// 归并排序通常有一些优化点：
// 1. 元素较少时，使用插入排序
// 2. 全局使用一个临时数组
// 3. 本身有序就不用排序了
func mergeSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	if first.Equal(last.Clone().Pre()) {
		return
	}

	if first.Distance(last) <= 7 {
		insertSort(first, last, cmp...)
		return
	}

	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	mid := first.Clone().NextN(first.Distance(last) / 2)
	if first.Distance(last)%2 != 0 {
		mid.Next()
	}
	mergeSort(first, mid.Clone(), cmp...)
	mergeSort(mid.Clone(), last, cmp...)

	i := first.Clone()
	j := mid.Clone()

	if c.Operator(j.Value(), j.Clone().Pre().Value()) {
		return
	}

	tmp := vector.New()

	for !j.Equal(last) && !i.Equal(mid) && !i.Equal(j) {
		if !c.Operator(i.Value(), j.Value()) {
			tmp.PushBack(i.Value())
			i.Next()
		} else {
			tmp.PushBack(j.Value())
			j.Next()
		}
	}
	for !i.Equal(mid) {
		tmp.PushBack(i.Value())
		i.Next()
	}
	for !j.Equal(last) {
		tmp.PushBack(j.Value())
		j.Next()
	}
	Copy(tmp.CBegin(), tmp.CEnd(), first)
}

func timSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
}

// 猴子排序
// 随机打乱数组，直到排序成功
// 复杂度为 O(?)
// 仅供娱乐
func bogoSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	isOrder := func() bool {
		for i := first.Clone(); !i.Equal(last.Clone().Pre()); i.Next() {
			if c.Operator(i.Value(), i.Clone().Next().Value()) {
				return false
			}
		}
		return true
	}

	for !isOrder() {
		Shuffle(first, last)
		ForEach(first, last, func(val interface{}) {
			fmt.Print(val)
		})
		fmt.Println()
	}
}

// 睡眠排序
// 开 n 个协程睡眠对应的时间，按照醒来次序排序
// 复杂度为 O(max(nums))
// 仅供娱乐
func sleepSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	t := vector.New()

	var wg sync.WaitGroup

	for i := first.Clone(); !i.Equal(last); i.Next() {
		wg.Add(1)
		go func(v interface{}) {
			time.Sleep(time.Duration(v.(int)) * time.Second * 2)
			t.PushBack(v)
			wg.Done()
		}(i.Value())
	}

	wg.Wait()
	Copy(t.CBegin(), t.CEnd(), first)
}

// 快速排序
// 时间复杂度 O(N*log^n)，空间复杂度 O(log^n)
// 时间：最好情况每次递归都平分数组，一共需要递归 logn 次，
// 每次需要 n 时间，复杂度为 O (n*logn)，最坏情况每次都把数组分成 1 和 n-1，
// 一共需要递归 n 次，每次需要 n 时间，总体复杂度为 O (n^2)。
// 平均总体时间复杂度为 O (nlogn)。
// 空间：和时间复杂度相关，每次递归需要的空间是固定的，
// 总体空间复杂度即为递归层数，因此平均 / 最好空间复杂度为 O (logn)，
// 最坏空间复杂度为 O (n)
// 非稳定排序
// 快排的优化点：
// 1. 元素较少时，使用插入排序（或者 shell 排序）
// 2. 合理选择锚点，最简单和常见的是选择 nums[l]，优化的有随机选择锚点，或者两边界+中间元素的中位数
// 锚点的选择是为了避免划分数组一方太小而另一方太大，这样容易退化成冒泡排序
// 3. 合理分割集合，这是快排的重点，通常比较常见的是先把锚点放在首元素，其它划分好后再交换回去
// 常见的优化有避免元素重复太多带来的不必要交换，有二路快排或三路快排等方式，就是根据大小划分为几个集合
// 二路快排就是将等于锚点的元素均分到锚点两边，使用对撞指针，然后相等的也交换即可
// 三路快排则是将等于锚点的元素放在中间，分成小于锚点的，等于锚点的，大于锚点的三个集合
// 4. 虽然锚点的合理选择会减少数组划分的过于不均匀，不过通常划分还是不可能完全均匀的
// 所以必然一个元素多，一个少，而常见的优化就是两个集合一个递归展开，一个迭代展开
// 而不是两个都递归展开。通常是元素少的一方递归展开，而多的迭代展开
// 5. 还有就是当元素太多时，递归调用层次会非常深，这样带来的开销非常大
// 故另一个常见的优化是设置固定的递归层数，当递归层数过深时，就转换为堆排序
// 6. 多线程优化
// 7. 预排序检查
func quickSort(first, last iterator.RandomAccessIterator, depth int, cmp ...comparator.Comparator) {
	depth--

	if depth == 0 {
		heapSort(first, last, cmp...)
	}

	// 结束递归
	if first.Distance(last) <= 1 {
		return
	}

	// 如果已经排序，则返回
	if IsSorted(first, last, cmp...) {
		return
	}

	var c comparator.Comparator = comparator.NewEGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	// 元素个数小于 7 时，使用插入排序 (shell)
	if first.Distance(last) < 7 {
		shellSort(first, last, c)
		// insertSort(first, last, cmp...)
	}

	pivot := doPivot(first, last, c)
	i, j := partition(first, last, pivot, c)

	quickSort(first, i, depth, c)
	quickSort(j, last, depth, c)
}

// 划分获取 pivot 锚点
// 有获取首元素，三数取中位数，随机选择几种方式
// 默认取中位数
func doPivot(first, last iterator.RandomAccessIterator, c comparator.Comparator) iterator.RandomAccessIterator {
	// 即首元素，中间元素和最后元素的中位数
	pivot := first
	l := last.Clone().Pre()
	mid := first.Clone().NextN(first.Distance(last) / 2)

	if (c.Operator(first.Value(), mid.Value()) && c.Operator(mid.Value(), l.Value())) || (c.Operator(l.Value(), mid.Value()) && c.Operator(mid.Value(), first.Value())) {
		pivot = mid
	} else if (c.Operator(first.Value(), l.Value()) && c.Operator(l.Value(), mid.Value())) || (c.Operator(mid.Value(), l.Value()) && c.Operator(l.Value(), first.Value())) {
		pivot = l
	}

	return pivot
}

// 划分数组,三路快排
func partition(first, last, pivot iterator.RandomAccessIterator, c comparator.Comparator) (iterator.RandomAccessIterator, iterator.RandomAccessIterator) {
	// 先和首元素交换位置
	Swap(first, pivot)

	// 待处理元素
	i := first.Clone().Next()
	// 小于区的下一个元素
	lt := first.Clone().Next()
	// 大于区的前一个元素
	gt := last.Clone().Pre()
	// 锚点
	v := first.Value()

	for i.IsFrontEqual(gt) {
		if c.Cmp(i.Value(), v) == -1 {
			Swap(lt, i)
			if i.Equal(lt) {
				i.Next()
			}
			lt.Next()
		} else if c.Cmp(i.Value(), v) == 1 {
			Swap(gt, i)
			gt.Pre()
		} else {
			i.Next()
		}
	}

	// 交换回锚点
	Swap(first, lt.Clone().Pre())

	if c.Cmp(v, gt.Value()) == 0 {
		gt.Next()
	}

	return lt.Pre(), gt
}

// 划分数组,二路快排
func partitionTwo(first, last, pivot iterator.RandomAccessIterator, c comparator.Comparator) (iterator.RandomAccessIterator, iterator.RandomAccessIterator) {
	// 先和首元素交换位置
	Swap(first, pivot)

	p := Partition(first.Clone().Next(), last, func(i interface{}) bool {
		return c.Operator(first.Value(), i)
	})

	// 交换回锚点
	Swap(first, p.Clone().Pre())

	return p.Clone().Pre(), p
}

// 获取最大递归深度, returns 2*ceil(lg(n+1)).
func getDepth(first, last iterator.RandomAccessIterator) int {
	var depth int
	for i := first.Distance(last); i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

// 堆排序
func heapSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	fmt.Println(c)

}

// 计数排序
func countSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	fmt.Println(c)
}

// 桶排序
func bucketSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	fmt.Println(c)
}

// 基数排序
func radioSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	fmt.Println(c)
}

// Sort elements in range
//  [first,last)
// greater comparator
// 元素较少时使用 shell 排序，数据量较大时使用快排，递归层数太多使用堆排
// https://feihu.me/blog/2014/sgi-std-sort/
func Sort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	if first.Distance(last) < 47 {
		shellSort(first, last, cmp...)
	}
	quickSort(first, last, getDepth(first, last), cmp...)
}

// StableSort sort elements preserving order of equivalents
func StableSort(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	mergeSort(first, last, cmp...)
}

// PartialSort partially sort elements in range
func PartialSort(first, middle, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	fmt.Println(c)
}

// PartialSortCopy copy and partially sort range
func PartialSortCopy() {

}

// IsSorted check whether range is sorted
func IsSorted(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) bool {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	for i := first.Clone(); !i.Equal(last.Clone().Pre()); i.Next() {
		if c.Operator(i.Value(), i.Clone().Next().Value()) {
			return false
		}
	}

	return true
}

// IsSortedUntil Find first unsorted element in range
func IsSortedUntil(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	var c comparator.Comparator = comparator.NewGreater()
	if len(cmp) != 0 {
		c = cmp[0]
	}

	for i := first.Clone(); !i.Equal(last.Clone().Pre()); i.Next() {
		if c.Operator(i.Value(), i.Clone().Next().Value()) {
			return i.Clone().Next()
		}
	}

	return nil
}

// NthElement sort element in range
// Rearranges the elements in the range [first,last), in such
// a way that the element at the nth position is the element
// that would be in that position in a sorted sequence.
// The other elements are left without any specific order,
// except that none of the elements preceding nth are greater
// than it, and none of the elements following it are less.
// The elements are compared using operator< for the first
// version, and comp for the second.
func NthElement(first, nth, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {

}
