// @program:     tiny-stl
// @file:        multimap.go
// @author:      edte
// @create:      2022-01-02 22:26
// @description:
package multimap

import (
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
)

type MultiMap struct {
}

func (m MultiMap) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Clear() {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Size() int {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Construct() {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) String() string {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Begin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) End() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) RBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) REnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) CBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) CEnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) CRBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) CREnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Find(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Count(key interface{}) int {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) LowerBound(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) UpperBound(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) EqualRange(key interface{}) (iterator.BidirectionalIterator, iterator.BidirectionalIterator) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) At(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) SetAt(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Insert(val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) InsertKey(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) InsertPosition(position iterator.InputIterator, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) InsertRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Erase(key interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) ErasePosition(position iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) EraseRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) Emplace(value interface{}) (iterator.BidirectionalIterator, bool) {
	//TODO implement me
	panic("implement me")
}

func (m MultiMap) EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}
