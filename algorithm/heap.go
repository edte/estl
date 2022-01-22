// @program:     tiny-stl
// @file:        heap.go
// @author:      edte
// @create:      2022-01-02 22:40
// @description:
package algorithm

import (
	"fmt"
	"strings"

	"stl/comparator"
	"stl/containers/queue"
	"stl/iterator"
)

// PushHeap push element into heap range
func PushHeap(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	c := comparator.Default(cmp...)
	siftUp(first, last.Clone().(iterator.RandomAccessIterator).Pre().(iterator.RandomAccessIterator), c)
}

// 上浮第 i 个元素
func siftUp(first, i iterator.RandomAccessIterator, c comparator.Comparator) {
	for {
		// 如果已经到达堆顶
		if i.Equal(first) {
			return
		}
		// 获取父节点坐标
		p := getParent(first, i)

		// 如果不满足，则交换
		if !c.Operator(i.Value(), p.Value()) {
			Swap(i, p)
		}

		// 指向父节点
		i = p
	}
}

// 下沉第 i 个元素
func siftDown(first, last, i iterator.RandomAccessIterator, c comparator.Comparator) {
	leaf := getFirstLeaf(first, last)

	for {
		// 如果当前节点变成叶子节点，那么无法继续下沉，故结束
		if i.IsBackEqual(leaf) {
			return
		}

		l := getLeftChild(first, i)
		tmp := l

		// 如果右子节点存在，那么
		r := getRightChild(first, i)
		if r.Position() < last.Position() {
			tmp = BestElement(l, r.Clone().(iterator.RandomAccessIterator).Next().(iterator.RandomAccessIterator), comparator.Reserve(c))
		}

		if c.Operator(i.Value(), tmp.Value()) {
			Swap(i, tmp)
		}
		i = tmp
	}
}

// 首节点坐标为 n
// 当前节点坐标为 j+n
// 则父节点坐标为 (j-1)/2+n
func getParent(first, i iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	n := first.Position()
	j := i.Position() - n
	index := (j-1)/2 + n
	return i.Clone().(iterator.RandomAccessIterator).IteratorAt(index)
}

// 首节点坐标为 n
// 当前节点坐标为 j+n
// 则左子节点坐标为 2*j+1+n
func getLeftChild(first, i iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	n := first.Position()
	j := i.Position() - n
	index := 2*j + 1 + n
	return i.Clone().(iterator.RandomAccessIterator).IteratorAt(index)
}

// 首节点坐标为 n
// 当前节点坐标为 j+n
// 则右子节点坐标为 2*j+2+n
func getRightChild(first, i iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	n := first.Position()
	j := i.Position() - n
	index := 2*j + 2 + n
	return i.Clone().(iterator.RandomAccessIterator).IteratorAt(index)
}

// 获取最后一个非叶子节点
// 总节点个数为 m
// 首节点坐标为 n
// 则最后一个非叶子节点坐标为 n+m/2-1
// 第一叶子节点坐标为 n+m/2
func getLastNotLeaf(first, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	m := first.Distance(last)
	n := first.Position()
	index := n + m/2 - 1
	return first.Clone().(iterator.RandomAccessIterator).IteratorAt(index)
}

// 获取第一个叶子节点
func getFirstLeaf(first, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	m := first.Distance(last)
	n := first.Position()
	index := n + m/2
	return first.Clone().(iterator.RandomAccessIterator).IteratorAt(index)
}

// PopHeap pop element from heap range
func PopHeap(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	c := comparator.Default(cmp...)
	Swap(first, last.Clone().(iterator.RandomAccessIterator).Pre())
	siftDown(first, last.Clone().(iterator.RandomAccessIterator).Pre().(iterator.RandomAccessIterator), first, c)
}

// MakeHeap make heap from range
func MakeHeap(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	c := comparator.Default(cmp...)
	leaf := getLastNotLeaf(first, last)

	// 从下往上遍历所有非叶子节点
	// 对于每个非叶子节点，对其下浮，就生成了堆
	for leaf.IsBackEqual(first) {
		siftDown(first, last, leaf, c)
		leaf.Pre()
	}
}

// SortHeap selectSort elements of heap
// 要使用堆排，必须已经是一个堆才行
func SortHeap(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) {
	c := comparator.Default(cmp...)
	i := Pre(last)

	for !i.Equal(first) {
		Swap(first, i)

		siftDown(first, i, first, c)

		i.Pre()
	}

	Reverse(first, last)
}

// IsHeap test if range is heap
func IsHeap(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) bool {
	c := comparator.Default(cmp...)
	q := queue.New()
	q.Push(first)

	for !q.Empty() {
		i := q.Front().(iterator.RandomAccessIterator)
		q.Pop()

		l := getLeftChild(first, i)
		if l.Position() < last.Position() {
			if !c.Operator(l.Value(), i.Value()) {
				return false
			}
			q.Push(l)
		}
		r := getRightChild(first, i)
		if r.Position() < last.Position() {
			if !c.Operator(r.Value(), i.Value()) {
				return false
			}
			q.Push(r)
		}
	}

	return true
}

// IsHeapUntil find first element not in heap order
func IsHeapUntil(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) iterator.RandomAccessIterator {
	c := comparator.Default(cmp...)
	q := queue.New()
	q.Push(first)

	for !q.Empty() {
		i := q.Front().(iterator.RandomAccessIterator)
		q.Pop()

		l := getLeftChild(first, i)
		if l.Position() < last.Position() {
			if !c.Operator(l.Value(), i.Value()) {
				return i
			}
			q.Push(l)
		}
		r := getRightChild(first, i)
		if r.Position() < last.Position() {
			if !c.Operator(r.Value(), i.Value()) {
				return i
			}
			q.Push(r)
		}
	}

	return nil
}

// 层序打印堆
func ShowHeap(first, last iterator.RandomAccessIterator) string {
	b := strings.Builder{}

	// fmt.Println()
	b.WriteString("\n")

	level := getLevel(first, last)
	nowLevel := 0

	q := queue.New()
	q.Push(first)
	q.Push(0)

	tmp := level
	for tmp != 0 {
		// fmt.Print("  ")
		b.WriteString("   ")
		tmp--
	}

	for {
		if q.Empty() {
			break
		}
		j := q.Front()
		q.Pop()

		_, ok := j.(int)
		if ok {
			if q.Empty() {
				break
			}
			level--
			nowLevel++

			// fmt.Println()
			b.WriteString("\n")

			tmp = level
			for tmp != 0 {
				// fmt.Print("  ")
				b.WriteString("   ")
				tmp--
			}

			q.Push(0)
			continue
		}

		i := j.(iterator.RandomAccessIterator)

		// fmt.Print(i)
		// fmt.Print("  ")

		b.WriteString(fmt.Sprint(i))
		tm := nowLevel
		for tm != 0 {
			b.WriteString("  ")
			tm--
		}

		l := getLeftChild(first, i)
		if l.Position() < last.Position() {
			q.Push(l)
		}
		r := getRightChild(first, i)
		if r.Position() < last.Position() {
			q.Push(r)
		}
	}

	// fmt.Println()
	// fmt.Println()
	b.WriteString("\n")
	b.WriteString("\n")

	return b.String()
}

// 获取堆的层次
func getLevel(first, last iterator.RandomAccessIterator) int {
	f := getFirstLeaf(first, last)
	level := 1
	i := Clone(first)

	for {
		if !i.IsFront(f) {
			break
		}
		level++
		i = getLeftChild(first, i)
	}

	return level
}
