// @program:     tiny-stl
// @file:        container.go
// @author:      edte
// @create:      2021-12-30 10:12
// @description:
package containers

import (
	"stl/algorithm/hash"
	"stl/allocator"
	"stl/iterator"
)

type Container interface {
	Swap(t *Container)
	Clone() Container

	Clear()
	Size() int
	MaxSize() int
	Capacity() int
	Empty() bool

	GetAllocator() allocator.Allocator
	Construct()
	Destroy()

	String() string
}

// SequentialContainer
// 顺序容器，顺序表,vector (向量)、list（列表）、deque（队列）、string
type SequentialContainer interface {
	Container

	Begin() iterator.RandomAccessIterator
	End() iterator.RandomAccessIterator
	RBegin() iterator.RandomAccessIterator
	REnd() iterator.RandomAccessIterator
	CBegin() iterator.RandomAccessIterator
	CEnd() iterator.RandomAccessIterator
	CRBegin() iterator.RandomAccessIterator
	CREnd() iterator.RandomAccessIterator

	ReSize(size int, value ...int)

	Reserve(n int)

	ShrinkToFit()

	At(i int) interface{}
	Front() interface{}
	Back() interface{}
	Data() []interface{}

	SetAt(i int, value interface{})
	Assign(size int, value int)
	AssignIter(a, b iterator.RandomAccessIterator)
	PushBack(data interface{})

	Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator
	InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator
	InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator

	Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator
	EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator

	Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator
	EmplaceBack(value interface{}) iterator.RandomAccessIterator
	PopBack()
	PopFront()
}

type Vector interface {
	SequentialContainer
}

type List interface {
	SequentialContainer
	Reserve(n int)
	Slice()
	Remove()
	RemoveIf()
	Unique()
	Merge()
	Sort()
}

type ForwardList interface {
	List
}

type Deque interface {
	SequentialContainer
}

type String interface {
	SequentialContainer
}

// AssociativeContainer
// 关联容器,有序二叉树（红黑树）
// map（集合）、set（映射）、multimap（多重集合）、multiset（多重映射）
type AssociativeContainer interface {
	Container

	Begin() iterator.BidirectionalIterator
	End() iterator.BidirectionalIterator
	RBegin() iterator.BidirectionalIterator
	REnd() iterator.BidirectionalIterator
	CBegin() iterator.BidirectionalIterator
	CEnd() iterator.BidirectionalIterator
	CRBegin() iterator.BidirectionalIterator
	CREnd() iterator.BidirectionalIterator

	Find(key interface{}) iterator.BidirectionalIterator
	Count(key interface{}) int
	LowerBound(key interface{}) iterator.BidirectionalIterator
	UpperBound(key interface{}) iterator.BidirectionalIterator
	EqualRange(key interface{}) (iterator.BidirectionalIterator, iterator.BidirectionalIterator)

	At(key interface{}) interface{}
	SetAt(key interface{}, value interface{})

	Insert(val interface{})
	InsertKey(key interface{}, value interface{})
	InsertPosition(position iterator.InputIterator, value interface{})
	InsertRange(first, last iterator.InputIterator)

	Erase(key interface{})
	ErasePosition(position iterator.InputIterator)
	EraseRange(first, last iterator.InputIterator)

	Emplace(value interface{}) (iterator.BidirectionalIterator, bool)
	EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator
}

type Set interface {
	AssociativeContainer
}

type Map interface {
	AssociativeContainer
}

type MultiMap interface {
	AssociativeContainer
}

type MultiSet interface {
	AssociativeContainer
}

type UnorderedAssociativeContainer interface {
	Container

	Begin() iterator.ForwardIterator
	End() iterator.ForwardIterator
	CBegin() iterator.ForwardIterator
	CEnd() iterator.ForwardIterator

	Find(key interface{}) iterator.ForwardIterator
	Count(key interface{}) int
	EqualRange(key interface{}) (iterator.ForwardIterator, iterator.ForwardIterator)

	At(key interface{}) interface{}
	SetAt(key interface{}, value interface{})

	Insert(val interface{})
	InsertKey(key interface{}, value interface{})
	InsertPosition(position iterator.InputIterator, value interface{})
	InsertRange(first, last iterator.InputIterator)

	Erase(key interface{})
	ErasePosition(position iterator.InputIterator)
	EraseRange(first, last iterator.InputIterator)

	Emplace(value interface{}) (iterator.BidirectionalIterator, bool)
	EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator

	BucketCount() int
	MaxBucketCount() int
	BucketSize(n int) int
	Bucket(k int) int

	LoadFactor() float64
	MaxLoadFactor() float64
	Rehash(n int)
	Reserve(n int)
	HashFunction() hash.Hash
}

type UnorderedSet interface {
	UnorderedAssociativeContainer
}

type UnorderedMap interface {
	UnorderedAssociativeContainer
}

type UnorderedMultiMap interface {
	UnorderedAssociativeContainer
}

type UnorderedMultiSet interface {
	UnorderedAssociativeContainer
}

// ContainerAdapter 三种容器适配器，接口映射，底层使用其它容器 stack,queue,priority_queue
type ContainerAdapter interface {
	Container

	Push(v interface{})
	Emplace(v interface{})
	Pop()
}

type Stack interface {
	ContainerAdapter
	Top() interface{}
}

type Queue interface {
	ContainerAdapter
	Front()
	Back()
}

type PriorityQueue interface {
	ContainerAdapter
	Top() interface{}
}

type Boom interface {
	Empty() bool
	Size() int
	Insert(key []byte)
	Contains(key []byte) bool
	Clear()
}

type BitSet interface {
	// bit access
	Count() uint
	Size() uint
	Test(uint) bool
	Any() bool
	None() bool
	All() bool

	// bit operations
	Set(uint)
	UnSet(uint)
	Clear(uint)
	Flip(uint)

	// bitset operations
	ToUint64() uint64
	ToUint32() uint32
	ToString() string
}
