// @program:     tiny-stl
// @file:        vector.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package vector

import (
	"fmt"

	"stl/comparator"
	"stl/iterator"
)

// *Vector
// Vectors are sequence containers representing arrays that can change in Size().
//
// Just like arrays, vectors use contiguous storage locations for their elements,
// which means that their elements can also be accessed using offsets on regular pointers to its elements,
// and just as efficiently as in arrays. But unlike arrays, their Size() can change dynamically,
// with their storage being handled automatically by the container.
//
// Internally, vectors use a dynamically allocated array to store their elements.
// This array may need to be reallocated in order to grow in Size() when new elements are inserted,
// which implies allocating a new array and moving all elements to it.
// This is a relatively expensive task in terms of processing time, and thus,
// vectors do not reallocate each time an element is added to the container.
//
// Instead, vector containers may allocate some extra storage to accommodate for possible growth,
// and thus the container may have an actual capacity greater than the storage strictly needed to contain its elements (i.e., its Size()).
// Libraries can implement different strategies for growth to balance between memory usage and reallocations,
// but in any case, reallocations should only happen at logarithmically growing intervals of Size() so that the insertion
// of individual elements at the end of the vector can be provided with amortized constant time complexity (see push_back).
//
// Therefore, compared to arrays, vectors consume more memory in exchange for the ability to manage storage and grow
// dynamically in an efficient way.
//
// Compared to the other dynamic sequence containers (deques, lists and forward_lists),
// vectors are very efficient accessing its elements (just like arrays) and relatively
// efficient adding or removing elements from its end. For operations that involve inserting
// or removing elements at positions other than the end, they perform worse than the others,
// and have less consistent iterators and references than lists and forward_lists.

type Option func(*Vector)

type Vector struct {
	data []interface{}
	cmp  comparator.Comparator
}

// New Construct
func New(opts ...Option) *Vector {
	v := &Vector{
		data: make([]interface{}, 0, 0),
		cmp:  &comparator.Less{},
	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func WithVector(nv *Vector) Option {
	return func(v *Vector) {
		v.data = make([]interface{}, nv.Size(), nv.Capacity())
		for i := range v.data {
			v.data[i] = nv.data[i]
		}
	}
}

func WithData(data []interface{}) Option {
	return func(v *Vector) {
		v.data = make([]interface{}, len(data), cap(data))
		for i := range v.data {
			v.data[i] = data[i]
		}
	}
}

func WithCap(c int) Option {
	return func(v *Vector) {
		v.data = make([]interface{}, 0, c)
	}
}

func WithCapInit(c int, value int) Option {
	return func(v *Vector) {
		v.data = make([]interface{}, 0, c)
		for i := range v.data {
			v.data[i] = value
		}
	}
}

func WithOperator(c comparator.Comparator) Option {
	return func(v *Vector) {
		v.cmp = c
	}
}

func WithCmp(cmp func(a interface{}, b interface{}) bool) Option {
	return func(v *Vector) {
		v.cmp = comparator.New(cmp)
	}
}

func (v *Vector) iterAt(pos int) iterator.RandomAccessIterator {
	return NewIterator(v, pos)
}

// Begin Return iterator to beginning
func (v *Vector) Begin() iterator.RandomAccessIterator {
	return v.iterAt(0)
}

// End Return iterator to end
// (0,size)
// begin=0
// end=size
func (v *Vector) End() iterator.RandomAccessIterator {
	return v.iterAt(v.Size())
}

// CBegin Return reverse iterator to reverse beginning
func (v *Vector) CBegin() iterator.RandomAccessIterator {
	return v.iterAt(v.Size() - 1)
}

// CEnd Return reverse iterator to reverse end
func (v *Vector) CEnd() iterator.RandomAccessIterator {
	return v.iterAt(-1)
}

// CRbegin Return const_iterator to beginning
func (v *Vector) CRbegin() iterator.InputIterator {
	panic("implement me")
}

// CRend Return const_iterator to end
func (v *Vector) CRend() iterator.InputIterator {
	panic("implement me")
}

// Rbegin Return const_reverse_iterator to reverse beginning
func (v *Vector) Rbegin() iterator.RandomAccessIterator {
	panic("implement me")
}

// Rend Return const_reverse_iterator to reverse end
func (v *Vector) Rend() iterator.RandomAccessIterator {
	panic("implement me")
}

// Size Return Size()
func (v *Vector) Size() int {
	return len(v.data)
}

// MaxSize Return maximum Size()
func (v *Vector) MaxSize() int {
	return v.Capacity()
}

// ReSize Change Size()
func (v *Vector) ReSize(size int, value ...int) {
	if v.Size() >= size {
		v.data = v.data[:size]
		return
	}
	if len(value) == 0 {
		return
	}
	for i := size - v.Size(); i != 0; i-- {
		v.PushBack(value[0])
	}
}

// Capacity Return Size() of allocated storage capacity
func (v *Vector) Capacity() int {
	return cap(v.data)
}

// Empty Test whether vector is empty
func (v *Vector) Empty() bool {
	return v.Size() == 0
}

// Reserve Request a change in capacity
func (v *Vector) Reserve(n int) {
	if n <= v.Capacity() {
		return
	}
}

// ShrinkToFit Shrink to fit
func (v *Vector) ShrinkToFit() {
	if v.Size() == v.Capacity() {
		return
	}
	data := make([]interface{}, v.Size(), v.Size())
	for i := range v.data {
		data[i] = v.data[i]
	}
	v.data = data
}

// At Access element
func (v *Vector) At(i int) interface{} {
	return v.data[i]
}

func (v *Vector) SetAt(i int, value interface{}) {
	v.data[i] = value
}

// Front Access first element
func (v *Vector) Front() interface{} {
	return v.At(0)
}

// Back Access last element
func (v *Vector) Back() interface{} {
	return v.At(v.Size() - 1)
}

// Data Access data
func (v *Vector) Data() []interface{} {
	return v.data
}

// Assign vector content
func (v *Vector) Assign(size int, value int) {
	v.data = make([]interface{}, size, size)
	for i := range v.data {
		v.data[i] = value
	}
}

func (v *Vector) AssignIter(a iterator.RandomAccessIterator, b iterator.RandomAccessIterator) {
	i1 := a.Clone()
	i2 := b.Clone()

	for !i1.Equal(i2) {
		v.PushBack(i1.Value())
		i1.Next()
	}
}

// PushBack Add element at the end
func (v *Vector) PushBack(data interface{}) {
	v.data = append(v.data, data)
}

// PopBack Delete last element
func (v *Vector) PopBack() {
	v.data = v.data[:v.Size()-1]
}

// Insert elements
func (v *Vector) Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	p := pos.Position()
	v.insert(p, value)
	return NewIterator(v, p)
}

func (v *Vector) InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator {
	p := pos.Position()
	for i := 0; i < size; i++ {
		v.insert(p+i, value)
	}
	return NewIterator(v, p+size)
}

func (v *Vector) InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	size := first.Distance(last)
	p := pos.Position()
	j := 0
	for i := first; !i.Equal(last); i.Next() {
		v.insert(p+j, i.Value())
		j++
	}
	return NewIterator(v, p+size)
}

func (v *Vector) insert(pos int, value interface{}) {
	if pos < 0 || pos > v.Size() {
		return
	}
	v.data = append(v.data, value)
	for i := len(v.data) - 1; i > pos; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[pos] = value
}

func (v *Vector) erase(first int, last int) {
	if first > last || first < 0 || last > v.Size() {
		return
	}
	l := v.data[:first]
	r := v.data[last:]
	v.data = append(l, r...)
}

// Erase elements
func (v *Vector) Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	p := pos.Position()
	v.erase(p, p+1)
	return NewIterator(v, p)
}

func (v *Vector) EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	l := first.Position()
	r := last.Position()
	v.erase(l, r)
	return NewIterator(v, l)
}

// Swap content
func (v *Vector) Swap(n *Vector) {
	t := v.data
	v.data = n.data
	n.data = t
}

// Clear content
func (v *Vector) Clear() {
	v.data = v.data[:0]
}

// Emplace Construct and insert element
func (v *Vector) Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	return v.Insert(pos, value)
}

// EmplaceBack Construct and insert element at the end
func (v *Vector) EmplaceBack(value interface{}) iterator.RandomAccessIterator {
	return v.Insert(v.End().Pre(), value)
}

func (v *Vector) String() string {
	var res string
	for i := range v.data {
		res += fmt.Sprint(v.data[i], " ")
	}
	return res
}

func (v *Vector) Clone() *Vector {
	return New(WithVector(v))
}