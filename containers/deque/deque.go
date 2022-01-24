// @program:     tiny-stl
// @file:        deque.go
// @author:      edte
// @create:      2022-01-15 17:37
// @description:
package deque

import (
	"fmt"
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
	"stl/locker"
)

type Option func(*Deque)

type Deque struct {
	data      []interface{}
	locker    locker.Locker
	allocator allocator.Allocator
}

func New(opts ...Option) *Deque {
	d := &Deque{
		data:      make([]interface{}, 0, 0),
		locker:    locker.NewMutex(),
		allocator: nil,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithGoroutineSafe(l locker.Locker) Option {
	return func(v *Deque) {
		v.locker = l
	}
}

func WithAllocator(a allocator.Allocator) Option {
	return func(v *Deque) {
		v.allocator = a
	}
}

func WithData(data []interface{}) Option {
	return func(v *Deque) {
		v.data = make([]interface{}, len(data), cap(data))
		for i := range v.data {
			v.data[i] = data[i]
		}
	}
}

func (d *Deque) Swap(t *containers.Container) {
	n := (*t).(*Deque)
	tt := d.data
	d.data = n.data
	n.data = tt
}

func (d *Deque) Clone() containers.Container {
	return New(WithData(d.data))
}

func (d *Deque) Clear() {
	d.data = d.data[:0]
}

func (d *Deque) Size() int {
	return len(d.data)
}

func (d *Deque) MaxSize() int {
	return cap(d.data)
}

func (d *Deque) Capacity() int {
	return cap(d.data)
}

func (d *Deque) Empty() bool {
	return d.Size() == 0
}

func (d *Deque) GetAllocator() allocator.Allocator {
	return d.allocator
}

func (d *Deque) Construct() {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) String() string {
	var res string
	for i := range d.data {
		res += fmt.Sprint(d.data[i], " ")
	}
	return res
}

func (d *Deque) Begin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) End() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) RBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) REnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) CBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) CEnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) CRBegin() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) CREnd() iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) ReSize(size int, value ...int) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) ShrinkToFit() {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) At(i int) interface{} {
	return d.data[i]
}

func (d *Deque) Front() interface{} {
	return d.At(0)
}

func (d *Deque) Back() interface{} {
	return d.At(d.Size() - 1)
}

func (d *Deque) Data() []interface{} {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) SetAt(i int, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Assign(size int, value int) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) AssignIter(a, b iterator.RandomAccessIterator) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) PushBack(data interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Insert(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) InsertSize(pos iterator.RandomAccessIterator, size int, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) InsertIter(pos iterator.RandomAccessIterator, first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Erase(pos iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) EraseIter(first iterator.RandomAccessIterator, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) Emplace(pos iterator.RandomAccessIterator, value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) EmplaceBack(value interface{}) iterator.RandomAccessIterator {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) PopBack() {
	//TODO implement me
	panic("implement me")
}

func (d *Deque) PopFront() {
	//TODO implement me
	panic("implement me")
}
